package timeslots_test

import (
	"timeslots"
	"time"
)

type TestingHelper struct {
	now time.Time
}

func NewTestingHelper(now time.Time) *TestingHelper {
	return &TestingHelper{now}
}

func (t *TestingHelper) Span(start, end int) *timeslots.Span {
	span, _ := timeslots.NewSpan(t.now.Add(time.Duration(start)*time.Hour), t.now.Add(time.Duration(end)*time.Hour))
	return span
}

func (t *TestingHelper) Block(start, end int) *timeslots.Block {
	return timeslots.NewBlockWithoutValidating(t.now.Add(time.Duration(start)*time.Hour), t.now.Add(time.Duration(end)*time.Hour))
}

func (t *TestingHelper) Slot(start, end int) *timeslots.Slot {
	slot, _ := timeslots.NewSlot(t.now.Add(time.Duration(start)*time.Hour), t.now.Add(time.Duration(end)*time.Hour))
	return slot
}

func (t *TestingHelper) HugeBlocks(start, end int) []*timeslots.Block {
	r := []*timeslots.Block{}
	for i := start; i < end; i = i + 2 {
		r = append(r, t.Block(i, i+1))
	}
	return r
}

func (t *TestingHelper) HugeSlots(start, end int) []*timeslots.Slot {
	r := []*timeslots.Slot{}
	for i := start; i < end; i = i + 2 {
		r = append(r, t.Slot(i, i+1))
	}
	return r
}
