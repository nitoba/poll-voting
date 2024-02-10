package repositories_test

import "github.com/nitoba/poll-voting/internal/domain/poll/enterprise/entities"

type InMemoryPollsRepository struct {
	Polls []*entities.Poll
}

func (r *InMemoryPollsRepository) Create(poll *entities.Poll) error {
	r.Polls = append(r.Polls, poll)
	return nil
}

func (r *InMemoryPollsRepository) FindById(id string) (*entities.Poll, error) {
	for _, poll := range r.Polls {
		if poll.Id.String() == id {
			return poll, nil
		}
	}
	return nil, nil
}
