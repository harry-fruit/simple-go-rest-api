package interfaces

type Repository[T any] interface {
	FindById(int) *T
}
