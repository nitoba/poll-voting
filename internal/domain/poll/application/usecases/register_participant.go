package usecases

import (
	"github.com/nitoba/poll-voting/internal/domain/poll/application/cryptography"
	"github.com/nitoba/poll-voting/internal/domain/poll/application/repositories"
	"github.com/nitoba/poll-voting/internal/domain/poll/application/usecases/errors"
	"github.com/nitoba/poll-voting/internal/domain/poll/enterprise/entities"
	"github.com/nitoba/poll-voting/internal/domain/poll/enterprise/value_objects"
)

type RegisterParticipantRequest struct {
	Name     string
	Email    string
	Password string
}

type RegisterParticipantUseCase struct {
	participantRepository repositories.ParticipantsRepository
	hasher                cryptography.HashGenerator
}

func (u *RegisterParticipantUseCase) Execute(req *RegisterParticipantRequest) error {
	existsParticipant, err := u.participantRepository.FindByEmail(req.Email)

	if err != nil || existsParticipant != nil {
		return errors.ErrParticipantAlreadyExists
	}

	email, err := value_objects.NewEmail(req.Email)

	if err != nil {
		return err
	}

	passwordHashed, err := u.hasher.Hash(req.Password)

	if err != nil {
		return err
	}

	participant, err := entities.NewParticipant(req.Name, email, passwordHashed)

	if err != nil {
		return err
	}

	if err = u.participantRepository.Create(participant); err != nil {
		return err
	}

	return nil
}

func NewRegisterParticipantUseCase(participantRepository repositories.ParticipantsRepository, hasher cryptography.HashGenerator) *RegisterParticipantUseCase {
	return &RegisterParticipantUseCase{
		participantRepository: participantRepository,
		hasher:                hasher,
	}
}
