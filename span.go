package slotify

import (
	"fmt"
	"time"
)

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

func (s *Span) Start() time.Time {
	return s.start
}

func (s *Span) End() time.Time {
	return s.end
}

func (s *Span) String() string {
	return format(s)
}

func (s *Span) Clone() *Span {
	return newSpan(s.start, s.end)
}

func (s *Span) ToSlot() *Slot {
	return newSlot(s.start, s.end)
}

func (s *Span) Remain() bool {
	return s.start.Before(s.end)
}

func (s *Span) Shorten(block *Block) {
	s.start = block.end
}

func (s *Span) Drop() {
	s.start = s.end
}
