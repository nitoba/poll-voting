package polls_presenter

import "github.com/nitoba/poll-voting/internal/domain/poll/enterprise/entities"

type FetchPollsResponse struct {
	Title   string   `json:"title"`
	Options []string `json:"options"`
}

func PollsToHttp(polls []*entities.Poll) []FetchPollsResponse {
	var pollsToResponse []FetchPollsResponse = []FetchPollsResponse{}
	for _, p := range polls {
		options := []string{}
		for _, o := range p.Options {
			options = append(options, o.Title)
		}

		pollsToResponse = append(pollsToResponse, FetchPollsResponse{
			Title:   p.Title,
			Options: options,
		})
	}

	return pollsToResponse
}
