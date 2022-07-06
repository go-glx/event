package event

import (
	"github.com/fe3dback/glx-event/event/internal/datas"
)

const warmQueueCount = 128

type Topic[T any] struct {
	queue          []T
	consumerLastId uint
	consumers      *datas.LinkedList[*consumer[T]]
}

func newTopic[T any]() *Topic[T] {
	return &Topic[T]{
		queue:          make([]T, 0, warmQueueCount),
		consumerLastId: 0,
		consumers:      datas.NewLinkedList[*consumer[T]](),
	}
}

// Emit will queue data to next handle call
// after handle call, Topic will send T to
// all consumers
func (d *Topic[T]) Emit(data T) {
	d.queue = append(d.queue, data)
}

// On will add new subscription
func (d *Topic[T]) On(fn func(T)) (cancel func()) {
	currentID := d.consumerLastId
	d.consumerLastId++

	currentConsumer := newConsumer(currentID, fn)
	d.consumers.Append(currentConsumer)

	return func() {
		d.consumers.Remove(currentConsumer)
	}
}

func (d *Topic[T]) handle() {
	for _, event := range d.queue {
		d.consumers.Iterate(func(consumer *consumer[T]) {
			consumer.consume(event)
		})
	}

	// set size to 0
	// actually we not reallocate all this memory
	// underline slice
	d.queue = d.queue[:0]
}
