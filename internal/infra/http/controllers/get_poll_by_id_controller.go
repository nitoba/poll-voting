package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nitoba/poll-voting/internal/domain/poll/application/usecases"
	polls_presenter "github.com/nitoba/poll-voting/internal/infra/http/presenters/polls"
)

type GetPollByIdController struct {
	getPollByIdUseCase *usecases.GetPollByIdUseCase
}

// GetJWT godoc
// @Summary      Get Poll by ID
// @Description  Get poll by ID in the API
// @Tags         polls
// @Accept       json
// @Produce      json
// @Success      200  {object} polls_presenter.GetPollByIdResponse
// @Failure      400  {object} Error
// @Failure      500  {object} Error
// @Param pollId path string true "poll id"
// @Router       /polls/{pollId} [get]
// @Security ApiKeyAuth
func (ct *GetPollByIdController) Handle(c *gin.Context) {
	userId := c.GetString("user_id")

	if userId == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	pollId := c.Param("id")

	if pollId == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Missing poll id",
		})
		return
	}

	poll, err := ct.getPollByIdUseCase.Execute(pollId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, polls_presenter.PollToHttp(poll))
}

func NewGetPollByIdController(getPollByIdUseCase *usecases.GetPollByIdUseCase) *GetPollByIdController {
	return &GetPollByIdController{
		getPollByIdUseCase: getPollByIdUseCase,
	}
}
