package interfaces

type Entity[T any] interface {
	Get(int) *T
}
