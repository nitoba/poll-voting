package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nitoba/poll-voting/internal/domain/poll/application/usecases"
)

type RegisterVoterController struct {
	registerVoterUseCase *usecases.RegisterVoterUseCase
}

type RegisterVoterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (ct *RegisterVoterController) Handle(c *gin.Context) {
	var reqBody RegisterVoterRequest
	c.Bind(&reqBody)

	err := ct.registerVoterUseCase.Execute(&usecases.RegisterVoterRequest{
		Name:     reqBody.Name,
		Email:    reqBody.Email,
		Password: reqBody.Password,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}

func NewRegisterVoterController(registerVoterUseCase *usecases.RegisterVoterUseCase) *RegisterVoterController {
	return &RegisterVoterController{
		registerVoterUseCase: registerVoterUseCase,
	}
}
