package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nitoba/poll-voting/internal/domain/poll/application/usecases"
	usecases_errors "github.com/nitoba/poll-voting/internal/domain/poll/application/usecases/errors"
)

type AuthenticateVoterController struct {
	authenticateVoterUseCase *usecases.AuthenticateUseCase
}

type AuthenticateVoterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AuthenticateVoterResponse struct {
	AccessToken string `json:"access_token"`
}

// GetJWT godoc
// @Summary      Authenticate Voters
// @Description  Authenticate voters in the API
// @Tags         voters
// @Accept       json
// @Param        request   body     AuthenticateVoterRequest  true  "voter credentials"
// @Produce      json
// @Success      201  {object} AuthenticateVoterResponse
// @Failure      400  {object} Error
// @Failure      500  {object} Error
// @Router       /auth/authenticate [post]
func (ct *AuthenticateVoterController) Handle(c *gin.Context) {
	var reqBody AuthenticateVoterRequest
	err := c.ShouldBindJSON(&reqBody)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	res, err := ct.authenticateVoterUseCase.Execute(usecases.AuthenticateRequest{
		Email:    reqBody.Email,
		Password: reqBody.Password,
	})

	if err != nil && err == usecases_errors.ErrWrongCredentials {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err != nil && err != usecases_errors.ErrWrongCredentials {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.SetCookie("auth", res.AccessToken, 3600*24*7, "/", "localhost", false, true)

	c.JSON(http.StatusOK, AuthenticateVoterResponse{
		AccessToken: res.AccessToken,
	})
}

func NewAuthenticateVoterController(authenticateVoterUseCase *usecases.AuthenticateUseCase) *AuthenticateVoterController {
	return &AuthenticateVoterController{
		authenticateVoterUseCase: authenticateVoterUseCase,
	}
}
