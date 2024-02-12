package core

type Entity struct {
	Id UniqueEntityId
	// Equals(other BaseEntity) bool
}

func NewEntity(id ...UniqueEntityId) *Entity {
	if len(id) == 0 {
		id = append(id, NewUniqueEntityId())
	}

	println("Tamanho de id: ", len(id))

	return &Entity{
		Id: id[0],
	}
}
