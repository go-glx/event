package event

type Topic[T any] struct {
	queue     []*event
	consumers []func(T)
}

func NewTopic[T any]() *Topic[T] {
	return &Topic[T]{
		queue:     make([]*event, 0),
		consumers: make([]func(T), 0),
	}
}

func (d *Topic[T]) Produce(data T) {
	d.queue = append(d.queue, newEvent(data))
}

func (d *Topic[T]) Consume(fn func(T)) {
	d.consumers = append(d.consumers, fn)
}

func (d *Topic[T]) drain() {
	for _, event := range d.queue {
		for _, consume := range d.consumers {
			consume(event.data.(T))
		}
	}

	d.queue = nil
}
