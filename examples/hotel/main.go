package main

import (
	"fmt"
	"goslotify"
	"strings"
	"time"
)

type Booking struct {
	RoomId uint
	StartAt  time.Time
	EndAt    time.Time
}


// Please implement this method in your struct to satisfy the Period interface.
func (b *Booking) Start() time.Time {
	return b.StartAt
}

// Please implement this method in your struct to satisfy the Period interface.
func (b *Booking) End() time.Time {
	return b.EndAt
}

func main() {

	now := time.Now()
	oneday := 24 * time.Hour

	// This variable will probably be retrieved from something like a request. Since this is an example, we’ll create it artificially.
	searchPeriod, err := goslotify.NewSpan(now, now.Add(8*oneday))
	if err != nil {
		panic(err)
	}
	fmt.Println("Search Period:\n" + searchPeriod.String() + "\n")

	// This variable will probably be retrieved from something like a database record. Since this is an example, we’ll create it artificially.
	events := []*Booking{
		{
			RoomId: 1,
			StartAt:  now.Add(0 * oneday),
			EndAt:    now.Add(1 * oneday),
		},
		{
			RoomId: 1,
			StartAt:  now.Add(2 * oneday),
			EndAt:    now.Add(3 * oneday),
		},
		{
			RoomId: 1,
			StartAt:  now.Add(6 * time.Hour),
			EndAt:    now.Add(7 * time.Hour),
		},
	}
	fmt.Println("Scheduled Events (blocks: []*Booking):\n" + toString(events, func(builder *strings.Builder, event *Booking) {
		builder.WriteString(fmt.Sprintf("%s, %s", event.StartAt.String(), event.EndAt.String()))
		builder.WriteString("\n")
	}))

	// To convert internally, define the map function for your input
	mapin := func(s *Booking) *goslotify.Block {
		b, _ := goslotify.NewBlock(s.StartAt, s.EndAt)
		return b
	}

	// To convert internally, define the map function for your output
	type RoomSlot struct {
		RoomId uint
		Start  time.Time
		End    time.Time
	}
	mapout := func(s *goslotify.Slot) *RoomSlot {
		return &RoomSlot{
			RoomId: 1,
			Start:  s.Start(),
			End:    s.End(),
		}
	}

	// To filter internally, define the filter function for your output
	filter := func(t *RoomSlot) bool {
		return t.End.Sub(t.Start) < (2 * oneday) // If the free time is within 48 hours, it will not be considered as free time.
	}

	// Find available time slots!
	slots := goslotify.FindWithMapper(events, searchPeriod, mapin, mapout, goslotify.WithFilter(filter))
	fmt.Println("Available Times (slots: []*TimeSlot):\n" + toString(slots, func(builder *strings.Builder, slot *RoomSlot) {
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
