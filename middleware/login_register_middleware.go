package middleware

import (
	"assignment-golang-backend/dto"
	"assignment-golang-backend/sentinelerrors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func (m *middleware) LoginRegisterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request dto.UserDto
		if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": sentinelerrors.ErrInvalidRequestBody.Error()})
			return
		}
		c.Next()
	}
}
