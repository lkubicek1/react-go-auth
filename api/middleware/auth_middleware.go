package middleware

import (
	"fmt"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type AuthMiddleware struct{}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (a *AuthMiddleware) ValidateToken(c *gin.Context) {
	auth0Domain := os.Getenv("AUTH0_DOMAIN")
	auth0Audience := os.Getenv("AUTH0_AUDIENCE")

	issuerURL, err := url.Parse("https://" + auth0Domain + "/")
	if err != nil {
		log.Fatalf("Failed to parse the issuer URL: %v", err)
	}

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{auth0Audience},
		validator.WithAllowedClockSkew(time.Minute), // Allow a 1-minute skew
	)
	if err != nil {
		log.Fatalf("Failed to set up the JWT validator: %v", err)
	}

	// Get the token from the request header
	authHeader := c.GetHeader("Authorization")
	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid authorization header",
		})
		c.Abort()
		return
	}

	// Validate the token
	_, err = jwtValidator.ValidateToken(c.Request.Context(), authHeaderParts[1])
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid token",
		})
		c.Abort()
		return
	}

	// Parse the token to access its claims
	token, _, err := new(jwt.Parser).ParseUnverified(authHeaderParts[1], jwt.MapClaims{})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Error parsing token"})
		c.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token claims"})
		c.Abort()
		return
	}

	userID, ok := claims["sub"].(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User ID not found in token"})
		c.Abort()
		return
	}

	// Set userID in Gin context
	c.Set("userID", userID)

	// Continue to the next middleware
	c.Next()
}
