package usecases

import (
	"github.com/nitoba/poll-voting/internal/domain/poll/application/cryptography"
	"github.com/nitoba/poll-voting/internal/domain/poll/application/repositories"
	"github.com/nitoba/poll-voting/internal/domain/poll/application/usecases/errors"
	"github.com/nitoba/poll-voting/internal/domain/poll/enterprise/entities"
	"github.com/nitoba/poll-voting/internal/domain/poll/enterprise/value_objects"
)

type RegisterVoterRequest struct {
	Name     string
	Email    string
	Password string
}

type RegisterVoterUseCase struct {
	voterRepository repositories.VotersRepository
	hasher          cryptography.HashGenerator
}

func (u *RegisterVoterUseCase) Execute(req *RegisterVoterRequest) error {
	existsVoter, err := u.voterRepository.FindByEmail(req.Email)

	if existsVoter != nil && err == nil {
		return errors.ErrVoterAlreadyExists
	}

	email, err := value_objects.NewEmail(req.Email)

	if err != nil {
		return err
	}

	passwordHashed, err := u.hasher.Hash(req.Password)

	if err != nil {
		return err
	}

	voter, err := entities.NewVoter(req.Name, email, passwordHashed)

	if err != nil {
		return err
	}

	if err = u.voterRepository.Create(voter); err != nil {
		return err
	}

	return nil
}

func NewRegisterVoterUseCase(voterRepository repositories.VotersRepository, hasher cryptography.HashGenerator) *RegisterVoterUseCase {
	return &RegisterVoterUseCase{
		voterRepository: voterRepository,
		hasher:          hasher,
	}
}
