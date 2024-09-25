// Find available time slot on schedule calender.
package slotify

import (
	"fmt"
	"time"
)

// This refers to already scheduled events. The term ‘Block’ will be standardized here.”
type Block struct {
	start time.Time
	end   time.Time
	Period
}

// Creates a new Block without validation. Use this when the order of start and end is guaranteed.
func NewBlockWithoutValidating(start, end time.Time) *Block {
	return &Block{
		start: start,
		end:   end,
	}
}

// Creates a new Block with validation. It verifies the order of start and end.
func NewBlock(start, end time.Time) (*Block, error) {
	if start.After(end) {
		return nil, fmt.Errorf("invalid time arguments")
	}
	return NewBlockWithoutValidating(start, end), nil
}

// Generates a slice of Blocks. Specify your struct with time-related fields as input, and define a mapping function between input and Block in the mapper.
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

// Start time of the period.
func (b *Block) Start() time.Time {
	return b.start
}

// End time of the period.
func (b *Block) End() time.Time {
	return b.end
}

// Represents the start time and end time as strings.
func (b *Block) String() string {
	return format(b)
}

func beforeEq(s, t time.Time) bool {
	return s.Before(t) || s.Equal(t)
}

// Whether the Block contains the given Period.
func (b *Block) Contains(other Period) bool {
	return beforeEq(b.start, other.Start()) && beforeEq(other.End(), b.end)
}

// Whether the Block is contained within the given Period.
func (b *Block) IsContainedIn(other Period) bool {
	return beforeEq(other.Start(), b.start) && beforeEq(b.end, other.End())
}

// Whether a period overlaps across the Block’s end time.
func (b *Block) OverlapAtEnd(other Period) bool {
	return beforeEq(other.Start(), b.start) && beforeEq(other.End(), b.end) && b.start.Before(other.End())
}

// Whether a period overlaps across the Block’s start time.
func (b *Block) OverlapAtStart(other Period) bool {
	return beforeEq(b.start, other.Start()) && beforeEq(b.end, other.End()) && other.Start().Before(b.end)
}
