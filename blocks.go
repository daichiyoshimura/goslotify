package slotify

import (
	"fmt"
	"time"
)

type Block struct {
	start time.Time
	end   time.Time
	Period
}

// These are the events that are already scheduled without validating.
func NewBlockWithoutValidating(start, end time.Time) *Block {
	return &Block{
		start: start,
		end:   end,
	}
}

// These are the events that are already scheduled.
func NewBlock(start, end time.Time) (*Block, error) {
	if start.After(end) {
		return nil, fmt.Errorf("invalid time arguments")
	}
	return NewBlockWithoutValidating(start, end), nil
}

// Please define a function for the second argument that maps the values passed to `slotify.NewBlock` to the fields of your struct.
func NewBlocks[T any](inputs []T, mapper func(T) (*Block, error)) ([]*Block, error) {
	blocks := []*Block{}
	for _, in := range inputs {
		block, err := mapper(in)
		if err != nil {
			return nil, err
		}
		blocks = append(blocks, block)
	}
	return blocks, nil
}

func (b *Block) Start() time.Time {
	return b.start
}

func (b *Block) End() time.Time {
	return b.end
}

func (b *Block) String() string {
	return format(b)
}

func beforeEq(s, t time.Time) bool {
	return s.Before(t) || s.Equal(t)
}

func (b *Block) Contains(other Period) bool {
	return beforeEq(b.start, other.Start()) && beforeEq(other.End(), b.end)
}

func (b *Block) IsContainedIn(other Period) bool {
	return beforeEq(other.Start(), b.start) && beforeEq(b.end, other.End())
}

func (b *Block) OverlapAtEnd(other Period) bool {
	return beforeEq(other.Start(), b.start) && beforeEq(other.End(), b.end) && b.start.Before(other.End())
}

func (b *Block) OverlapAtStart(other Period) bool {
	return beforeEq(b.start, other.Start()) && beforeEq(b.end, other.End()) && other.Start().Before(b.end)
}
