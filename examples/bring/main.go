// In this example, you will pass a struct that already exists in your implementation as an argument. 
// The returned value will also be a struct that exists in your implementation. 
// To achieve this mechanism, you will need to implement the Period interface and the mapping function between the struct you provided and the struct provided by the library.
package main

import (
	"fmt"
	"goslotify"
	"strings"
	"time"
)

type ScheduledEvent struct {
	TableId uint
	StartAt   time.Time
	EndAt     time.Time
}

// Please implement this method in your struct to satisfy the Period interface.
func (r *ScheduledEvent) Start() time.Time {
	return r.StartAt
}

// Please implement this method in your struct to satisfy the Period interface.
func (r *ScheduledEvent) End() time.Time {
	return r.EndAt
}

func main() {

	now := time.Now()

	// This variable will probably be retrieved from something like a request. Since this is an example, we’ll create it artificially.
	span, err := goslotify.NewSpan(now, now.Add(8*time.Hour))
	if err != nil {
		panic(err)
	}
	fmt.Println("Search Span:\n" + span.String() + "\n")

	// This variable will probably be retrieved from something like a database record. Since this is an example, we’ll create it artificially.
	events := []*ScheduledEvent{
		{
			TableId: 1,
			StartAt:   now.Add(0 * time.Hour),
			EndAt:     now.Add(1 * time.Hour),
		},
		{
			TableId: 1,
			StartAt:   now.Add(2 * time.Hour),
			EndAt:     now.Add(3 * time.Hour),
		},
		{
			TableId: 1,
			StartAt:   now.Add(6 * time.Hour),
			EndAt:     now.Add(7 * time.Hour),
		},
	}
	fmt.Println("Scheduled Events (blocks: []*ScheduledEvent):\n" + toString(events, func(builder *strings.Builder, event *ScheduledEvent) {
		builder.WriteString(fmt.Sprintf("%s, %s", event.StartAt.String(), event.EndAt.String()))
		builder.WriteString("\n")
	}))

	// To convert internally, define the map function for your input
	mapin := func(s *ScheduledEvent) *goslotify.Block {
		b, _ := goslotify.NewBlock(s.StartAt, s.EndAt)
		return b
	}

	// To convert internally, define the map function for your output
	type DiningTableSlot struct {
		TableId uint
		Start   time.Time
		End     time.Time
	}
	mapout := func(s *goslotify.Slot) *DiningTableSlot {
		return &DiningTableSlot{
			TableId: 1,
			Start:   s.Start(),
			End:     s.End(),
		}
	}

	// To filter internally, define the filter function for your output
	filter := func(t *DiningTableSlot) bool {
		return t.End.Sub(t.Start) < (2 * time.Hour) // If the free time is within 2 hours, it will not be considered as free time.
	}

	// Find available time slots!
	slots := goslotify.FindWithMapper(events, span, mapin, mapout, goslotify.WithFilter(filter))
	fmt.Println("Available Time Slots(slots: []*TimeSlot):\n" + toString(slots, func(builder *strings.Builder, slot *DiningTableSlot) {
		builder.WriteString(fmt.Sprintf("%s, %s", slot.Start.String(), slot.End.String()))
		builder.WriteString("\n")
	}))
}

func toString[T any](args []T, stringer func(*strings.Builder, T)) string {
	var builder strings.Builder
	for _, v := range args {
		stringer(&builder, v)
	}
	return builder.String()
}
