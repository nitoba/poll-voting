package polls_presenter

import "github.com/nitoba/poll-voting/internal/domain/poll/enterprise/entities"

type GetPollByIdResponse struct {
	Id      string   `json:"id"`
	Title   string   `json:"title"`
	Options []string `json:"options"`
}

func PollToHttp(polls *entities.Poll) GetPollByIdResponse {
	options := []string{}

	for _, o := range polls.Options {
		options = append(options, o.Title)
	}

	return GetPollByIdResponse{
		Id:      polls.Id.String(),
		Title:   polls.Title,
		Options: options,
	}
}
