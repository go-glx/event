# glx-event

Tiny library for event handling, designed for game engines

version: 0.2
__NOT READY FOR REAL USE__

## Usage

```go
import "github.com/fe3dback/glx-event/event"

disp := event.NewDispatcher()
topic := event.TopicFor[SomeEvent](disp)

// subscribe
topic.On(func(e SomeEvent) { .. })
topic.On(func(e SomeEvent) { .. })
topic.On(func(e SomeEvent) { .. })

// emit (will queue two events)
topic.Emit(SomeEvent{ .. })
topic.Emit(SomeEvent{ .. })

// HandleEvents usually should be called on frame start
// before World.Update and other logic.
// 
// this will send first event to all subscribers
// than send second event to all subscribers
disp.HandleEvents()

// next run do nothing, because all
// queued events already handled
disp.HandleEvents()
```

## Example

see [example in unit test](./event/dispatcher_test.go)
