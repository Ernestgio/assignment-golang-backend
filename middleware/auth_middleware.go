package middleware

import (
	"assignment-golang-backend/appconstants"
	"assignment-golang-backend/sentinelerrors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (m *Middleware) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader != "" {
			bearerToken := strings.Split(authHeader, " ")
			if len(bearerToken) == 2 {
				authToken := bearerToken[1]
				claim, err := m.hashUtils.ValidateToken(authToken)
				if err != nil {
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
					return
				}
				c.Set(appconstants.UserContextKey, claim.Id)
				c.Set(appconstants.WalletContextKey, claim.WalletId)
				c.Next()
				return
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": sentinelerrors.ErrInvalidToken})
			return
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": sentinelerrors.ErrInvalidToken.Error()})
	}
}
