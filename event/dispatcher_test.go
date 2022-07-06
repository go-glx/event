package event

import (
	"fmt"
)

type eventMouseMoved struct {
	X int
	Y int
}

func ExampleNewDispatcher() {
	d := NewDispatcher()

	topic := TopicOf[eventMouseMoved](d)

	// C1
	cancelC1 := topic.On(func(evt eventMouseMoved) {
		fmt.Printf("C1 X:%d, Y:%d\n", evt.X, evt.Y)
	})

	// Some events, will be queued
	topic.Emit(eventMouseMoved{X: 10, Y: 15})
	topic.Emit(eventMouseMoved{X: 5, Y: 50})
	topic.Emit(eventMouseMoved{X: 100, Y: 20})

	// C2
	topic.On(func(evt eventMouseMoved) {
		fmt.Printf("C2 X:%d, Y:%d\n", evt.X, evt.Y)
	})

	// Cancel C1
	cancelC1()

	// process step 1
	d.HandleEvents()
	fmt.Println("Step 1 handled")

	topic.Emit(eventMouseMoved{X: 99, Y: 0})

	// C3
	topic.On(func(evt eventMouseMoved) {
		fmt.Printf("C3 X:%d, Y:%d\n", evt.X, evt.Y)
	})

	// process step 2
	d.HandleEvents()
	fmt.Println("Step 2 handled")

	// Output:
	// C2 X:10, Y:15
	// C2 X:5, Y:50
	// C2 X:100, Y:20
	// Step 1 handled
	// C2 X:99, Y:0
	// C3 X:99, Y:0
	// Step 2 handled
}
