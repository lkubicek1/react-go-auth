package main

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

type Session struct {
	SessionID string    `bson:"session_id"`
	Username  string    `bson:"username"`
	ExpiresAt time.Time `bson:"expires_at"`
}

// Encrypts text with the given key
func encrypt(text string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	b := []byte(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], b)
	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

// Decrypts text with the given key
func decrypt(text string, key []byte) (string, error) {
	ciphertext, err := base64.URLEncoding.DecodeString(text)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	if len(ciphertext) < aes.BlockSize {
		return "", err // ciphertext too short
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)
	return string(ciphertext), nil
}

var client *mongo.Client
var sessionCollection *mongo.Collection

func initMongoDB(dbUser string, dbPass string, dbName string, dbHost string, dbPort string) {
	var err error
	connectionString := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	sessionCollection = client.Database(dbName).Collection("sessions")

	log.Println("Connected to MongoDB!")
}

func main() {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	initMongoDB(dbUser, dbPass, dbName, dbHost, dbPort)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Specify allowed origin(s) here
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	/*router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))*/

	router.POST("/register", registerHandler)
	router.POST("/login", loginHandler)
	router.GET("/logout", logoutHandler)
	router.GET("/session", sessionHandler)

	router.Run(":8080")
}

func registerHandler(c *gin.Context) {
	var newUser User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}
	newUser.Password = string(hashedPassword)

	dbName := os.Getenv("DB_NAME")
	userCollection := client.Database(dbName).Collection("users")
	_, err = userCollection.InsertOne(context.Background(), newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while registering user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "registered"})
}

func loginHandler(c *gin.Context) {
	var loginUser User
	if err := c.BindJSON(&loginUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dbName := os.Getenv("DB_NAME")
	userCollection := client.Database(dbName).Collection("users")
	var result User
	err := userCollection.FindOne(context.Background(), bson.M{"username": loginUser.Username}).Decode(&result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(loginUser.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	// Create a session
	sessionID := uuid.NewString()
	expiration := time.Now().Add(24 * time.Hour)

	session := Session{
		SessionID: sessionID,
		Username:  loginUser.Username,
		ExpiresAt: expiration,
	}

	_, err = sessionCollection.InsertOne(context.Background(), session)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating session"})
		return
	}

	encryptionKey := os.Getenv("ENCRYPTION_KEY")

	// Encryption key should be 16, 24, or 32 bytes long to select AES-128, AES-192, or AES-256.
	key := []byte(encryptionKey)

	encryptedSessionID, err := encrypt(sessionID, key)
	if err != nil {
		log.Printf("Error encrypting session ID: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to secure session"})
		return
	}

	// Set encrypted session ID as cookie
	c.SetCookie("session_id", encryptedSessionID, 86400, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"status": "logged in"})
}

func logoutHandler(c *gin.Context) {
	sessionID, err := c.Cookie("session_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Session not found"})
		return
	}

	_, err = sessionCollection.DeleteOne(context.Background(), bson.M{"session_id": sessionID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error terminating session"})
		return
	}

	// Clear cookie
	c.SetCookie("session_id", "", -1, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"status": "logged out"})
}

func sessionHandler(c *gin.Context) {
	encryptedSessionID, err := c.Cookie("session_id")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "Session not found", "loggedIn": false})
		return
	}

	encryptionKey := os.Getenv("ENCRYPTION_KEY")

	// Decrypt session ID
	key := []byte(encryptionKey) // Use the same key as used for encryption

	sessionID, err := decrypt(encryptedSessionID, key)
	if err != nil {
		log.Printf("Error decrypting session ID: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decrypt session"})
		return
	}

	var session Session
	err = sessionCollection.FindOne(context.Background(), bson.M{"session_id": sessionID}).Decode(&session)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "Session not valid", "loggedIn": false})
		return
	}

	// Check if session is expired
	if time.Now().After(session.ExpiresAt) {
		_, err = sessionCollection.DeleteOne(context.Background(), bson.M{"session_id": sessionID})
		if err != nil {
			log.Printf("Error deleting expired session: %v", err)
			// Not returning here to allow informing the client about the expired session
		}
		c.JSON(http.StatusOK, gin.H{"error": "Session expired", "loggedIn": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Session valid", "loggedIn": true, "username": session.Username})
}
