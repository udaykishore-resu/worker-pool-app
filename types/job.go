package types

type Job[T any] struct {
	ID   int
	Data T
}
