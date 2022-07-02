package event

import (
	"fmt"
)

type testEventMovedInto struct {
	X int
	Y int
}

func ExampleNewDispatcher() {
	d := NewDispatcher()

	movedIntoTopic := TopicFor[testEventMovedInto](d)

	// C1
	movedIntoTopic.Consume(func(evt testEventMovedInto) {
		fmt.Printf("C1 X:%d, Y:%d\n", evt.X, evt.Y)
	})

	// Some events, will be queued
	movedIntoTopic.Produce(testEventMovedInto{X: 10, Y: 15})
	movedIntoTopic.Produce(testEventMovedInto{X: 5, Y: 50})
	movedIntoTopic.Produce(testEventMovedInto{X: 100, Y: 20})

	// C2
	movedIntoTopic.Consume(func(evt testEventMovedInto) {
		fmt.Printf("C2 X:%d, Y:%d\n", evt.X, evt.Y)
	})

	// process step 1
	d.HandleEvents()
	fmt.Println("Step 1 handled")

	movedIntoTopic.Produce(testEventMovedInto{X: 99, Y: 0})

	// C3
	movedIntoTopic.Consume(func(evt testEventMovedInto) {
		fmt.Printf("C3 X:%d, Y:%d\n", evt.X, evt.Y)
	})

	// process step 2
	d.HandleEvents()
	fmt.Println("Step 2 handled")

	// Output:
	// C1 X:10, Y:15
	// C2 X:10, Y:15
	// C1 X:5, Y:50
	// C2 X:5, Y:50
	// C1 X:100, Y:20
	// C2 X:100, Y:20
	// Step 1 handled
	// C1 X:99, Y:0
	// C2 X:99, Y:0
	// C3 X:99, Y:0
	// Step 2 handled
}
