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
	searchSpan, err := goslotify.NewSpan(now, now.Add(8*time.Hour))
	if err != nil {
		panic(err)
	}
	fmt.Println("Search Period:\n" + searchSpan.String() + "\n")

	// This variable will probably be retrieved from something like a database record. Since this is an example, we’ll create it artificially.
	type Reservation struct {
		TableId uint
		Start   time.Time
		End     time.Time
	}
	reservations := []*Reservation{
		{
			TableId: 1,
			Start:   now.Add(0 * time.Hour),
			End:     now.Add(1 * time.Hour),
		},
		{
			TableId: 1,
			Start:   now.Add(2 * time.Hour),
			End:     now.Add(3 * time.Hour),
		},
		{
			TableId: 1,
			Start:   now.Add(6 * time.Hour),
			End:     now.Add(7 * time.Hour),
		},
	}
	fmt.Println("Scheduled Events (blocks: []*Reservation):\n" + toString(reservations, func(builder *strings.Builder, event *Reservation) {
		builder.WriteString(fmt.Sprintf("%s, %s", event.Start.String(), event.End.String()))
		builder.WriteString("\n")
	}))

	// To sort internally, define the sort function for your input
	sorter := func(i, j int, reservations []*Reservation) bool {
		return reservations[i].Start.Before(reservations[j].Start)
	}

	// To convert internally, define the map function for your input
	mapin := func(s *Reservation) *goslotify.Block {
		b, _ := goslotify.NewBlock(s.Start, s.End)
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
	slots := goslotify.FindWithMapper(reservations, searchSpan, sorter, mapin, mapout, goslotify.WithFilter(filter))
	fmt.Println("Available Times (slots: []*TimeSlot):\n" + toString(slots, func(builder *strings.Builder, slot *DiningTableSlot) {
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
