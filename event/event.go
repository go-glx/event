package event

type (
	eventID string

	event struct {
		id   eventID
		data any
	}
)

func newEvent(data any) *event {
	return &event{
		id:   eventID(resolveTypeID(data)),
		data: data,
	}
}
