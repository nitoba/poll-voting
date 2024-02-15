package core

type Entity struct {
	Id UniqueEntityId
	// Equals(other BaseEntity) bool
}

func NewEntity(id ...UniqueEntityId) *Entity {
	if len(id) == 0 {
		id = append(id, NewUniqueEntityId())
	}

	return &Entity{
		Id: id[0],
	}
}
