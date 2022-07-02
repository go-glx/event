package event

import "reflect"

type Dispatcher struct {
	topics map[string]any
}

type drainable interface {
	drain()
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		topics: make(map[string]any),
	}
}

func TopicFor[T any](d *Dispatcher) *Topic[T] {
	var TType T
	topicID := resolveTypeID(TType)

	if topic, exist := d.topics[topicID]; exist {
		return reflect.ValueOf(topic).Interface().(*Topic[T])
	}

	topic := NewTopic[T]()
	d.topics[topicID] = topic

	return topic
}

func (d *Dispatcher) HandleEvents() {
	for _, topic := range d.topics {
		topic.(drainable).drain()
	}
}
