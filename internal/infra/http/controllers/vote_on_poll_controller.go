package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nitoba/poll-voting/internal/domain/poll/application/usecases"
)

type VoteOnPollController struct {
	voteOnPollUseCase *usecases.VoteOnPollUseCase
}

type VoteOnPollRequest struct {
	Title   string   `json:"title" binding:"required,min=5,max=100"`
	Options []string `json:"options" binding:"required,min=2"`
}

// GetJWT godoc
// @Summary      Vote On Poll
// @Description  Vote On poll in the API
// @Tags         polls
// @Accept       json
// @Param        request   body     VoteOnPollRequest  true  "poll data"
// @Produce      json
// @Success      201
// @Failure      400  {object} Error
// @Failure      500  {object} Error
// @Router       /polls [post]
// @Security ApiKeyAuth
func (ct *VoteOnPollController) Handle(c *gin.Context) {
	var reqBody VoteOnPollRequest
	err := c.ShouldBindJSON(&reqBody)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	userId := c.GetString("user_id")

	if userId == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	c.Status(http.StatusCreated)
}

func NewVoteOnPollController(voteOnPollUseCase *usecases.VoteOnPollUseCase) *VoteOnPollController {
	return &VoteOnPollController{
		voteOnPollUseCase: voteOnPollUseCase,
	}
}
