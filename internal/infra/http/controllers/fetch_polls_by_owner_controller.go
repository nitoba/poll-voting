package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nitoba/poll-voting/internal/domain/poll/application/usecases"
	polls_presenter "github.com/nitoba/poll-voting/internal/infra/http/presenters/polls"
)

type FetchPollsByOwnerController struct {
	fetchPollsByOwnerUseCase *usecases.FetchPollsByOwnerUseCase
}

// GetJWT godoc
// @Summary      Fetch Polls from owner
// @Description  Fetch polls from owner in the API
// @Tags         polls
// @Accept       json
// @Produce      json
// @Success      200  {array} polls_presenter.FetchPollsResponse
// @Failure      400  {object} Error
// @Failure      500  {object} Error
// @Router       /polls [get]
// @Security ApiKeyAuth
func (ct *FetchPollsByOwnerController) Handle(c *gin.Context) {
	userId := c.GetString("user_id")

	if userId == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	res, err := ct.fetchPollsByOwnerUseCase.Execute(usecases.FetchPollsByOwnerRequest{
		OwnerId: userId,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, polls_presenter.PollsToHttp(res.Polls))
}

func NewFetchPollsByOwnerController(fetchPollsByOwnerUseCase *usecases.FetchPollsByOwnerUseCase) *FetchPollsByOwnerController {
	return &FetchPollsByOwnerController{
		fetchPollsByOwnerUseCase: fetchPollsByOwnerUseCase,
	}
}
