package slotify_test

import (
	"slotify"
	"time"
)

type TestingHelper struct {
	now time.Time
}

func NewTestingHelper(now time.Time) *TestingHelper {
	return &TestingHelper{now}
}

func (t *TestingHelper) Span(start, end int) *slotify.Span {
	span, _ := slotify.NewSpan(t.now.Add(time.Duration(start)*time.Hour), t.now.Add(time.Duration(end)*time.Hour))
	return span
}

func (t *TestingHelper) Block(start, end int) *slotify.Block {
	return slotify.NewBlockWithoutValidating(t.now.Add(time.Duration(start)*time.Hour), t.now.Add(time.Duration(end)*time.Hour))
}

func (t *TestingHelper) Slot(start, end int) *slotify.Slot {
	slot, _ := slotify.NewSlot(t.now.Add(time.Duration(start)*time.Hour), t.now.Add(time.Duration(end)*time.Hour))
	return slot
}

func (t *TestingHelper) HugeBlocks(start, end int) []*slotify.Block {
	r := []*slotify.Block{}
	for i := start; i < end; i = i + 2 {
		r = append(r, t.Block(i, i+1))
	}
	return r
}

func (t *TestingHelper) HugeSlots(start, end int) []*slotify.Slot {
	r := []*slotify.Slot{}
	for i := start; i < end; i = i + 2 {
		r = append(r, t.Slot(i, i+1))
	}
	return r
}
