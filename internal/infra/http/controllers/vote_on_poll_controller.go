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
	OptionId string `json:"option_id" binding:"required"`
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
// @Router       /polls/{pollId}/vote [post]
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

	pollId := c.Param("pollId")

	if pollId == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Missing pollId",
		})
		return
	}

	err = ct.voteOnPollUseCase.Execute(&usecases.VoteOnPollUseCaseRequest{
		PollId:       pollId,
		VoterId:      userId,
		PollOptionId: reqBody.OptionId,
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
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
