package core

type BaseEntity interface {
	Equals(other *BaseEntity) bool
}
