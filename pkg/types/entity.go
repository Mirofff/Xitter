package types

// Entity - ddd entity interface.
type Entity[T any] interface {
	isSame(other T) bool
}
