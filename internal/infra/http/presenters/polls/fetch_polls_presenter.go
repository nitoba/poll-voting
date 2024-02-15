package polls_presenter

import "github.com/nitoba/poll-voting/internal/domain/poll/enterprise/entities"

type FetchPollsResponse struct {
	Id      string   `json:"id"`
	Title   string   `json:"title"`
	Options []option `json:"options"`
}

func PollsToHttp(polls []*entities.Poll) []FetchPollsResponse {
	var pollsToResponse []FetchPollsResponse = []FetchPollsResponse{}
	for _, p := range polls {
		options := []option{}
		for _, o := range p.Options {
			options = append(options, option{
				Id:    o.Id.String(),
				Title: o.Title,
			})
		}

		pollsToResponse = append(pollsToResponse, FetchPollsResponse{
			Id:      p.Id.String(),
			Title:   p.Title,
			Options: options,
		})
	}

	return pollsToResponse
}
