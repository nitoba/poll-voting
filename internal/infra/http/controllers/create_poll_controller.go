package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nitoba/poll-voting/internal/domain/poll/application/usecases"
)

type CreatePollController struct {
	createPollUseCase *usecases.CreatePollUseCase
}

type CreatePollRequest struct {
	Title   string   `json:"title" binding:"required,min=5,max=100"`
	Options []string `json:"options" binding:"required,min=2"`
}

// GetJWT godoc
// @Summary      Create Polls
// @Description  Create polls in the API
// @Tags         polls
// @Accept       json
// @Param        request   body     CreatePollRequest  true  "poll data"
// @Produce      json
// @Success      201
// @Failure      404  {object} Error
// @Failure      500  {object} Error
// @Router       /polls [post]
// @Security ApiKeyAuth
func (ct *CreatePollController) Handle(c *gin.Context) {
	var reqBody CreatePollRequest
	err := c.ShouldBindJSON(&reqBody)

	userId := c.GetString("user_id")

	if userId == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = ct.createPollUseCase.Execute(usecases.CreatePollRequest{
		Title:   reqBody.Title,
		Options: reqBody.Options,
		OwnerId: userId,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}

func NewCreatePollController(createPollUseCase *usecases.CreatePollUseCase) *CreatePollController {
	return &CreatePollController{
		createPollUseCase: createPollUseCase,
	}
}
