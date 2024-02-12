package value_objects

type NewVote struct {
	PollOptionId string
	Count        int
}

func CreateNewVote(pollOptionId string, count int) *NewVote {
	return &NewVote{
		PollOptionId: pollOptionId,
		Count:        count,
	}
}
