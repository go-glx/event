package event

type consumer[T any] struct {
	id uint
	fn func(T)
}

func newConsumer[T any](id uint, fn func(T)) *consumer[T] {
	return &consumer[T]{
		id: id,
		fn: fn,
	}
}

func (c *consumer[T]) consume(data T) {
	c.fn(data)
}
