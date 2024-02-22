package services

type BaseService[T any] struct {
	idx  int
	max  int
	data *[]T
}

func NewBaseService[T any]() *BaseService[T] {
	return &BaseService[T]{
		idx:  1,
		max:  5,
		data: nil,
	}
}
