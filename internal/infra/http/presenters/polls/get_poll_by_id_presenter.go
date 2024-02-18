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
	Total   int      `json:"total"`
}

func PollToHttp(poll *entities.Poll) GetPollByIdResponse {
	options := []option{}

	for _, o := range poll.Options {
		options = append(options, option{
			Id:    o.Id.String(),
			Title: o.Title,
		})
	}

	return GetPollByIdResponse{
		Id:      poll.Id.String(),
		Title:   poll.Title,
		Options: options,
		Total:   poll.Votes,
	}
}
