package usecases

import (
	"errors"
	"slices"

	"github.com/nitoba/poll-voting/internal/domain/core"
	"github.com/nitoba/poll-voting/internal/domain/poll/application/repositories"
	"github.com/nitoba/poll-voting/internal/domain/poll/enterprise/entities"
)

type VoteOnPollUseCase struct {
	voteRepo  repositories.VotesRepository
	pollRepo  repositories.PollsRepository
	voterRepo repositories.VotersRepository
}

type VoteOnPollUseCaseRequest struct {
	PollId       string
	VoterId      string
	PollOptionId string
}

func (u *VoteOnPollUseCase) Execute(req *VoteOnPollUseCaseRequest) error {
	poll, err := u.pollRepo.FindById(req.PollId)

	if err != nil && poll == nil {
		return err
	}

	// check if poll option exists
	if !slices.ContainsFunc(poll.Options, func(option *entities.PollOption) bool {
		return option.Id.String() == req.PollOptionId
	}) {
		return errors.New("invalid poll option")
	}

	voter, err := u.voterRepo.FindById(req.VoterId)

	if err != nil && voter == nil {
		return err
	}

	// check if voter has voted on this poll option already
	previousVote, err := u.voteRepo.FindByOptionId(req.PollOptionId)

	if err == nil && previousVote != nil {
		// check if vote is the same as previous vote
		if previousVote.IsTheSameVoteOption(req.PollOptionId) {
			return errors.New("voter has already voted on this poll option")
		}

		// if previous vote exists, delete it and create a new one
		if err := u.voteRepo.Delete(previousVote); err != nil {
			return err
		}

		vote, err := entities.NewVote(poll.Id, core.NewUniqueEntityId(req.PollOptionId), voter.Id)

		if err != nil {
			return err
		}

		if err := u.voteRepo.Create(vote); err != nil {
			return err
		}
	} else {
		// if previous vote does not exist, create a new one
		vote, err := entities.NewVote(poll.Id, core.NewUniqueEntityId(req.PollOptionId), voter.Id)

		if err != nil {
			return err
		}

		if err := u.voteRepo.Create(vote); err != nil {
			return err
		}
	}

	return nil
}

func NewVoteOnPollUseCase(voteRepo repositories.VotesRepository, pollRepo repositories.PollsRepository, voterRepo repositories.VotersRepository) *VoteOnPollUseCase {
	return &VoteOnPollUseCase{
		voteRepo:  voteRepo,
		pollRepo:  pollRepo,
		voterRepo: voterRepo,
	}
}
