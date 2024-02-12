package repositories_test

type InMemoryCountingVotesRepository struct {
	Votes map[string]map[string]int
}

func (r *InMemoryCountingVotesRepository) IncrementCountVotesByOptionId(pollId string, optionId string) (int, error) {

	if _, ok := r.Votes[pollId]; !ok {
		r.Votes[pollId] = make(map[string]int)
	}
	r.Votes[pollId][optionId]++

	return r.Votes[pollId][optionId], nil
}

func (r *InMemoryCountingVotesRepository) DecrementCountVotesByOptionId(pollId string, optionId string) (int, error) {
	if _, ok := r.Votes[pollId]; !ok {
		r.Votes[pollId] = make(map[string]int)
	}
	r.Votes[pollId][optionId]--

	return r.Votes[pollId][optionId], nil
}

func NewInMemoryCountingVotesRepository() *InMemoryCountingVotesRepository {
	return &InMemoryCountingVotesRepository{
		Votes: make(map[string]map[string]int),
	}
}
