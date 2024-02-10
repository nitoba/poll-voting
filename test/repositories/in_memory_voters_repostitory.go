package repositories_test

import "github.com/nitoba/poll-voting/internal/domain/poll/enterprise/entities"

type InMemoryVotersRepository struct {
	Voters []*entities.Voter
}

func (repo *InMemoryVotersRepository) Create(voter *entities.Voter) error {
	repo.Voters = append(repo.Voters, voter)
	return nil
}

func (repo *InMemoryVotersRepository) FindById(id string) (*entities.Voter, error) {
	for _, p := range repo.Voters {
		if p.Id.String() == id {
			return p, nil
		}
	}
	return nil, nil
}

func (repo *InMemoryVotersRepository) FindByEmail(email string) (*entities.Voter, error) {
	for _, p := range repo.Voters {
		if p.Email.Value() == email {
			return p, nil
		}
	}
	return nil, nil
}
