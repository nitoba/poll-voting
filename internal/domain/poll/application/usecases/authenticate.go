package usecases

import (
	"fmt"

	"github.com/nitoba/poll-voting/internal/domain/poll/application/cryptography"
	"github.com/nitoba/poll-voting/internal/domain/poll/application/repositories"
	"github.com/nitoba/poll-voting/internal/domain/poll/application/usecases/errors"
)

type AuthenticateUseCase struct {
	voterRepository repositories.VotersRepository
	hasher          cryptography.HashComparer
	encrypter       cryptography.Encrypter
}

type AuthenticateRequest struct {
	Email    string
	Password string
}

type AuthenticateResponse struct {
	AccessToken string
}

func (u *AuthenticateUseCase) Execute(req AuthenticateRequest) (*AuthenticateResponse, error) {
	voter, err := u.voterRepository.FindByEmail(req.Email)

	if err != nil || voter == nil {
		return nil, errors.ErrWrongCredentials
	}

	if !u.hasher.Compare(req.Password, voter.Password) {
		fmt.Println(voter.Email.Value())
		return nil, errors.ErrWrongCredentials
	}

	token := u.encrypter.Encrypt(map[string]interface{}{
		"sub":   voter.Id.String(),
		"email": voter.Email.Value(),
		"name":  voter.Name,
	})

	return &AuthenticateResponse{
		AccessToken: token,
	}, nil
}

func NewAuthenticateUseCase(voterRepository repositories.VotersRepository, hasher cryptography.HashComparer, encrypter cryptography.Encrypter) *AuthenticateUseCase {
	return &AuthenticateUseCase{
		voterRepository: voterRepository,
		hasher:          hasher,
		encrypter:       encrypter,
	}
}
