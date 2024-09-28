// In this example, you will create a struct provided by the library and pass it as an argument. 
// The returned value will also be something provided by this library.
package main

import (
	"fmt"
	"timeslots"
	"time"
)

const TimeFormat = "2006-01-02 15:04:05"

func main() {

	now := time.Now()

	// This variable will probably be retrieved from something like a request. Since this is an example, we’ll create it artificially.
	span, err := timeslots.NewSpan(now, now.Add(8*time.Hour))
	if err != nil {
		panic(err)
	}
	fmt.Println("Search Span:\n" + span.String() + "\n")

	// This variable will probably be retrieved from something like a database record. Since this is an example, we’ll create it artificially.
	type ScheduledEvent struct {
		Start time.Time
		End   time.Time
	}
	events := []*ScheduledEvent{
		{
			Start: now.Add(-1 * time.Hour),
			End:   now.Add(1 * time.Hour),
		},
		{
			Start: now.Add(2 * time.Hour),
			End:   now.Add(4 * time.Hour),
		},
		{
			Start: now.Add(3 * time.Hour),
			End:   now.Add(5 * time.Hour),
		},
	}

	// Please define a function for the second argument that maps the values passed to `timeslots.NewBlock` to the fields of your struct.
	blocks, err := timeslots.NewBlocks(events, func(s *ScheduledEvent) (*timeslots.Block, error) {
		return timeslots.NewBlock(s.Start, s.End)
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Scheduled Events (blocks):\n" + timeslots.ToString(blocks))

	// Find available time slots!
	slots := timeslots.Find(blocks, span)
	fmt.Println("Available Time Slots(slots):\n" + timeslots.ToString(slots))
}
