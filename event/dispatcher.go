package event

import "reflect"

type (
	abstractTopic interface {
		handle()
	}
	topicTypeID = string

	Dispatcher struct {
		topics map[topicTypeID]abstractTopic
	}
)

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		topics: make(map[string]abstractTopic),
	}
}

func TopicOf[T any](d *Dispatcher) *Topic[T] {
	var TType T
	topicType := reflect.TypeOf(TType).String()

	if topic, exist := d.topics[topicType]; exist {
		return topic.(*Topic[T])
	}

	topic := newTopic[T]()
	d.topics[topicType] = topic

	return topic
}

func (d *Dispatcher) HandleEvents() {
	for _, topic := range d.topics {
		topic.handle()
	}
}
