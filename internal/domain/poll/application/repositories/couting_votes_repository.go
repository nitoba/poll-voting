package repositories

type CountingVotesRepository interface {
	IncrementCountVotesByOptionId(pollId string, optionId string) (int, error)
	DecrementCountVotesByOptionId(pollId string, optionId string) (int, error)
	CountVotesByOptionId(pollId string, optionId string) (int, error)
}
