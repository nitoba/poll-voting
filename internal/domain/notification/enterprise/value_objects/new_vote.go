package value_objects

type NewVote struct {
	VoterId      string
	PollId       string
	PollOptionId string
}

func (n NewVote) NewVote(voterId, pollId, pollOptionId string) *NewVote {
	return &NewVote{
		VoterId:      voterId,
		PollId:       pollId,
		PollOptionId: pollOptionId,
	}
}
