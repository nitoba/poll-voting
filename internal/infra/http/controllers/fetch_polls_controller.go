package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nitoba/poll-voting/internal/domain/poll/application/usecases"
	polls_presenter "github.com/nitoba/poll-voting/internal/infra/http/presenters/polls"
)

type FetchPollsController struct {
	fetchPollsUseCase *usecases.FetchPollsUseCase
}

// GetJWT godoc
// @Summary      Fetch Polls
// @Description  Fetch polls in the API
// @Tags         polls
// @Accept       json
// @Produce      json
// @Success      200  {array} polls_presenter.FetchPollsResponse
// @Failure      400  {object} Error
// @Failure      500  {object} Error
// @Router       /polls/all [get]
// @Security ApiKeyAuth
func (ct *FetchPollsController) Handle(c *gin.Context) {
	userId := c.GetString("user_id")

	if userId == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	polls, err := ct.fetchPollsUseCase.Execute()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, polls_presenter.PollsToHttp(polls))
}

func NewFetchPollsController(fetchPollsUseCase *usecases.FetchPollsUseCase) *FetchPollsController {
	return &FetchPollsController{
		fetchPollsUseCase: fetchPollsUseCase,
	}
}
