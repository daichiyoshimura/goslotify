package timeslots

import (
	"fmt"
	"time"
)

// This refers to available free time. The term ‘Slot’ will be standardized here.
type Slot struct {
	start time.Time
	end   time.Time
	Period
}

func newSlot(start, end time.Time) *Slot {
	return &Slot{
		start: start,
		end:   end,
	}
}

// Creates a new Slot.
func NewSlot(start, end time.Time) (*Slot, error) {
	if start.After(end) {
		return nil, fmt.Errorf("invalid time arguments")
	}
	return newSlot(start, end), nil
}

func createSlotFrom(span *Span, block *Block) *Slot {
	return newSlot(span.start, block.start)
}

// Start time of the period.
func (s *Slot) Start() time.Time {
	return s.start
}

// End time of the period.
func (s *Slot) End() time.Time {
	return s.end
}

// Represents the start time and end time as strings.
func (s *Slot) String() string {
	return format(s)
}

// Whether two Slots represent the same period
func (s *Slot) Equal(other *Slot) bool {
	return equal(s, other)
}

// Whether it has a start time earlier than the given Slot.
func (s *Slot) SmallerThan(other *Slot) bool {
	return s.start.Before(other.start)
}
