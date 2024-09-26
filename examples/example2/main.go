package main

import (
	"fmt"
	"goslotify"
	"strings"
	"time"
)

func main() {

	now := time.Now()

	// This variable will probably be retrieved from something like a request. Since this is an example, we’ll create it artificially.
	searchPeriod, err := goslotify.NewSpan(now, now.Add(8*time.Hour))
	if err != nil {
		panic(err)
	}
	fmt.Println("Search Period:\n" + searchPeriod.String() + "\n")

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
	fmt.Println("Scheduled Events (blocks: []*ScheduledEvent):\n" + toString(events, func(builder *strings.Builder, event *ScheduledEvent) {
		builder.WriteString(fmt.Sprintf("%s, %s", event.Start.String(), event.End.String()))
		builder.WriteString("\n")
	}))

	// To sort internally, define the sort function for your input
	sorter := func(i, j int, events []*ScheduledEvent) bool {
		return events[i].Start.Before(events[j].Start)
	}

	// To convert internally, define the map function for your input
	mapin := func(s *ScheduledEvent) (*goslotify.Block, error) {
		return goslotify.NewBlock(s.Start, s.End)
	}

	// To convert internally, define the map function for your output
	type TimeSlot struct {
		Start time.Time
		End   time.Time
	}
	mapout := func(s *goslotify.Slot) (*TimeSlot, error) {
		return &TimeSlot{
			Start: s.Start(),
			End:   s.End(),
		}, nil
	}
	filter := func(t *TimeSlot) bool {
		return false
	}

	// Find available time slots!
	slots, err := goslotify.FindWithMapper(events, searchPeriod, sorter, mapin, mapout, goslotify.WithFilter(filter))
	if err != nil {
		panic(err)
	}

	fmt.Println("Available Times (slots: []*TimeSlot):\n" + toString(slots, func(builder *strings.Builder, slot *TimeSlot) {
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
