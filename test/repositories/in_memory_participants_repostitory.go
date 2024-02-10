package repositories_test

import "github.com/nitoba/poll-voting/internal/domain/poll/enterprise/entities"

type InMemoryParticipantsRepository struct {
	Participants []*entities.Participant
}

func (repo *InMemoryParticipantsRepository) Create(participant *entities.Participant) error {
	repo.Participants = append(repo.Participants, participant)
	return nil
}

func (repo *InMemoryParticipantsRepository) FindById(id string) (*entities.Participant, error) {
	for _, p := range repo.Participants {
		if p.Id.String() == id {
			return p, nil
		}
	}
	return nil, nil
}

func (repo *InMemoryParticipantsRepository) FindByEmail(email string) (*entities.Participant, error) {
	for _, p := range repo.Participants {
		if p.Email.Value() == email {
			return p, nil
		}
	}
	return nil, nil
}
