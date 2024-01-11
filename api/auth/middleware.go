package auth

import (
	"api/config"
	"fmt"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Middleware struct {
	config config.EnvVars
}

func NewMiddleware(config config.EnvVars) *Middleware {
	return &Middleware{
		config: config,
	}
}

func (a *Middleware) ValidateToken(c *gin.Context) {
	issuerURL, err := url.Parse("https://" + a.config.AUTH0_DOMAIN + "/")
	if err != nil {
		log.Fatalf("Failed to parse the issuer url: %v", err)
	}

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{a.config.AUTH0_AUDIENCE},
	)
	if err != nil {
		log.Fatalf("Failed to set up the jwt validator")
	}

	// get the token from the request header
	authHeader := c.GetHeader("Authorization")
	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid authorization header",
		})
		return
	}

	// Validate the token
	tokenInfo, err := jwtValidator.ValidateToken(c, authHeaderParts[1])
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid token",
		})
		return
	}

	fmt.Println(tokenInfo)

	// If everything is ok, continue to the next middleware
	c.Next()
}
