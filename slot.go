package slotify

import (
	"fmt"
	"time"
)

// Available time slot returned as the result.
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

func NewSlot(start, end time.Time) (*Slot, error) {
	if start.After(end) {
		return nil, fmt.Errorf("invalid time arguments")
	}
	return newSlot(start, end), nil
}

func createSlotFrom(span *Span, block *Block) *Slot {
	return newSlot(span.start, block.start)
}

func (s *Slot) Start() time.Time {
	return s.start
}

func (s *Slot) End() time.Time {
	return s.end
}

func (s *Slot) String() string {
	return format(s)
}

func (s *Slot) Equal(other *Slot) bool {
	return equal(s, other)
}

func (s *Slot) SmallerThan(other *Slot) bool {
	return s.start.Before(other.start)
}
