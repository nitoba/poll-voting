package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/nitoba/poll-voting/internal/domain/poll/application/cryptography"
	"github.com/nitoba/poll-voting/internal/domain/poll/application/repositories"
)

func AuthRequiredCookie(encrypter cryptography.Encrypter, voterRepository repositories.VotersRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		access_token, err := c.Cookie("auth")

		if err != nil {
			sendUnauthorizedResponse(c, "Token Required")
			return
		}

		payload, err := encrypter.Verify(access_token)

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
