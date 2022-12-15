package middleware

import (
	"assignment-golang-backend/dto"
	"assignment-golang-backend/sentinelerrors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func (m *Middleware) LoginRegisterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.ShouldBindBodyWith(&dto.UserDto{}, binding.JSON); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": sentinelerrors.ErrInvalidRequestBody.Error()})
			return
		}
		c.Next()
	}
}

func (m *Middleware) TopupMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.ShouldBindBodyWith(&dto.TopupRequestDto{}, binding.JSON); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": sentinelerrors.ErrInvalidRequestBody.Error()})
			return
		}
		c.Next()
	}
}

func (m *Middleware) TransferMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.ShouldBindBodyWith(&dto.TransferDto{}, binding.JSON); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": sentinelerrors.ErrInvalidRequestBody.Error()})
			return
		}
		c.Next()
	}
}
