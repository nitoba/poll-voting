package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nitoba/poll-voting/internal/domain/poll/application/usecases"
	usecases_errors "github.com/nitoba/poll-voting/internal/domain/poll/application/usecases/errors"
)

type Error struct {
	Message string `json:"message"`
}

type RegisterVoterController struct {
	registerVoterUseCase *usecases.RegisterVoterUseCase
}

type RegisterVoterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// GetJWT godoc
// @Summary      Register Voters
// @Description  Register new voters in the API
// @Tags         voters
// @Accept       json
// @Param        request   body     RegisterVoterRequest  true  "voter credentials"
// @Produce      json
// @Success      201
// @Failure      400  {object} Error
// @Failure      500  {object} Error
// @Router       /auth/register [post]
func (ct *RegisterVoterController) Handle(c *gin.Context) {
	var reqBody RegisterVoterRequest
	err := c.ShouldBindJSON(&reqBody)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = ct.registerVoterUseCase.Execute(&usecases.RegisterVoterRequest{
		Name:     reqBody.Name,
		Email:    reqBody.Email,
		Password: reqBody.Password,
	})

	if err != nil && err == usecases_errors.ErrVoterAlreadyExists {
		c.JSON(http.StatusConflict, gin.H{
			"message": err.Error(),
		})

		return
	}

	if err != nil && err != usecases_errors.ErrVoterAlreadyExists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}

func NewRegisterVoterController(registerVoterUseCase *usecases.RegisterVoterUseCase) *RegisterVoterController {
	return &RegisterVoterController{
		registerVoterUseCase: registerVoterUseCase,
	}
}
