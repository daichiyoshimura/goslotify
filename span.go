package timeslots

import (
	"fmt"
	"time"
)

// This is the period for searching for free time. The term ‘Span’ will be standardized here. Note that the Span is mutable.
type Span struct {
	start time.Time
	end   time.Time
	Period
}

// Specify the period you want to search.
func NewSpan(start, end time.Time) (*Span, error) {
	if start.After(end) {
		return nil, fmt.Errorf("invalid time arguments")
	}
	return newSpan(start, end), nil
}

func newSpan(start, end time.Time) *Span {
	return &Span{
		start: start,
		end:   end,
	}
}

// Start time of the period.
func (s *Span) Start() time.Time {
	return s.start
}

// End time of the period.
func (s *Span) End() time.Time {
	return s.end
}

// Represents the start time and end time as strings.
func (s *Span) String() string {
	return format(s)
}

// Copy the values into a new instance to avoid mutating the original.
func (s *Span) Clone() *Span {
	return newSpan(s.start, s.end)
}

// Convert the Span into a Slot.
func (s *Span) ToSlot() *Slot {
	return newSlot(s.start, s.end)
}

// Whether there is remaining time in the period.
func (s *Span) Remain() bool {
	return s.start.Before(s.end)
}

// Shorten the period. (This assumes sorting, so it shortens from the start time)
func (s *Span) Shorten(block *Block) {
	s.start = block.end
}

// Eliminate the period
func (s *Span) Drop() {
	s.start = s.end
}
