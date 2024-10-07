package middlewares

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"strings"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		token = strings.Replace(token, "Bearer ", "", 1)
		claims := &jwt.MapClaims{}
		tkn, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte("your_secret_key"), nil
		})

		if err != nil || !tkn.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("userID", (*claims)["id"])
		c.Next()
	}
}
