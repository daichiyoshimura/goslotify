package goslotify_test

import (
	"goslotify"
	"time"
)

type TestingHelper struct {
	now time.Time
}

func NewTestingHelper(now time.Time) *TestingHelper {
	return &TestingHelper{now}
}

func (t *TestingHelper) Span(start, end int) *goslotify.Span {
	span, _ := goslotify.NewSpan(t.now.Add(time.Duration(start)*time.Hour), t.now.Add(time.Duration(end)*time.Hour))
	return span
}

func (t *TestingHelper) Block(start, end int) *goslotify.Block {
	return goslotify.NewBlockWithoutValidating(t.now.Add(time.Duration(start)*time.Hour), t.now.Add(time.Duration(end)*time.Hour))
}

func (t *TestingHelper) Slot(start, end int) *goslotify.Slot {
	slot, _ := goslotify.NewSlot(t.now.Add(time.Duration(start)*time.Hour), t.now.Add(time.Duration(end)*time.Hour))
	return slot
}

func (t *TestingHelper) HugeBlocks(start, end int) []*goslotify.Block {
	r := []*goslotify.Block{}
	for i := start; i < end; i = i + 2 {
		r = append(r, t.Block(i, i+1))
	}
	return r
}

func (t *TestingHelper) HugeSlots(start, end int) []*goslotify.Slot {
	r := []*goslotify.Slot{}
	for i := start; i < end; i = i + 2 {
		r = append(r, t.Slot(i, i+1))
	}
	return r
}
