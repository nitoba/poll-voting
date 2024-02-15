package polls_presenter

import "github.com/nitoba/poll-voting/internal/domain/poll/enterprise/entities"

type option struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

type GetPollByIdResponse struct {
	Id      string   `json:"id"`
	Title   string   `json:"title"`
	Options []option `json:"options"`
}

func PollToHttp(polls *entities.Poll) GetPollByIdResponse {
	options := []option{}

	for _, o := range polls.Options {
		options = append(options, option{
			Id:    o.Id.String(),
			Title: o.Title,
		})
	}

	return GetPollByIdResponse{
		Id:      polls.Id.String(),
		Title:   polls.Title,
		Options: options,
	}
}
