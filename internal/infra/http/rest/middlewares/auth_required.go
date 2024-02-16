package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nitoba/poll-voting/internal/domain/poll/application/cryptography"
	"github.com/nitoba/poll-voting/internal/domain/poll/application/repositories"
)

func sendUnauthorizedResponse(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"message": msg,
	})
}

func AuthRequired(encrypter cryptography.Encrypter, voterRepository repositories.VotersRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")

		if authorization == "" || !strings.Contains(authorization, "Bearer") {
			sendUnauthorizedResponse(c, "Token Required")
			return
		}

		token := strings.TrimPrefix(authorization, "Bearer ")

		payload, err := encrypter.Verify(token)

		if err != nil {
			sendUnauthorizedResponse(c, "Invalid token")
			return
		}

		userID := payload["sub"].(string)

		user, err := voterRepository.FindById(userID)

		if err != nil || user == nil {
			sendUnauthorizedResponse(c, "Unauthorized")
			return
		}
		c.Set("user_id", userID)
		c.Next()
	}
}
