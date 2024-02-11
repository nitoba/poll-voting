package core

import "github.com/google/uuid"

type UniqueEntityId struct {
	value string
}

func (id *UniqueEntityId) String() string {
	return id.value
}

func (id UniqueEntityId) Equals(other UniqueEntityId) bool {
	if id == other {
		return true
	}
	return id.String() == other.String()
}

func NewUniqueEntityId(id ...string) UniqueEntityId {
	if len(id) > 0 {
		return UniqueEntityId{
			value: id[0],
		}
	}

	return UniqueEntityId{
		value: uuid.New().String(),
	}
}
