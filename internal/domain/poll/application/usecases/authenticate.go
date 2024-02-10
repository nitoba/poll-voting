package usecases

import (
	"fmt"

	"github.com/nitoba/poll-voting/internal/domain/poll/application/cryptography"
	"github.com/nitoba/poll-voting/internal/domain/poll/application/repositories"
	"github.com/nitoba/poll-voting/internal/domain/poll/application/usecases/errors"
)

type AuthenticateUseCase struct {
	participantRepository repositories.ParticipantsRepository
	hasher                cryptography.HashComparer
	encrypter             cryptography.Encrypter
}

type AuthenticateRequest struct {
	Email    string
	Password string
}

type AuthenticateResponse struct {
	AccessToken string
}

func (u *AuthenticateUseCase) Execute(req AuthenticateRequest) (*AuthenticateResponse, error) {
	participant, err := u.participantRepository.FindByEmail(req.Email)

	if err != nil || participant == nil {
		return nil, errors.ErrWrongCredentials
	}

	if !u.hasher.Compare(req.Password, participant.Password) {
		fmt.Println(participant.Email.Value())
		return nil, errors.ErrWrongCredentials
	}

	token := u.encrypter.Encrypt(map[string]interface{}{
		"sub": participant.Id.String(),
	})

	return &AuthenticateResponse{
		AccessToken: token,
	}, nil
}

func NewAuthenticateUseCase(participantRepository repositories.ParticipantsRepository, hasher cryptography.HashComparer, encrypter cryptography.Encrypter) *AuthenticateUseCase {
	return &AuthenticateUseCase{
		participantRepository: participantRepository,
		hasher:                hasher,
		encrypter:             encrypter,
	}
}
