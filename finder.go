package goslotify

import (
	"sort"
)

// Sort your struct in chronological order.
type SortFunc[T any] func(int, int, []T) bool

// Map your struct to a Block.
type MapInFunc[T any] func(T) (*Block, error)

// Map the Slot to your struct.
type MapOutFunc[T any] func(*Slot) (T, error)

// Calculate available time slots (Slot). Provide the scheduled block (Block) and the target period (Span).
// Use this when passing and returning your struct.
func FindWithMapper[S, T any](inputs []S, span *Span, sorter SortFunc[S], mapin MapInFunc[S], mapout MapOutFunc[T]) ([]T, error) {
	if span == nil || !span.Remain() {
		return []T{}, nil
	}

	target := span.Clone()
	if len(inputs) == 0 {
		slot, err := mapout(target.ToSlot())
		if err != nil {
			return nil, err
		}
		return []T{slot}, nil
	}

	sort.Slice(inputs, func(i, j int) bool {
		return sorter(i, j, inputs)
	})

	j := 0
	slots := make([]T, len(inputs)+1)
	for _, input := range inputs {
		block, err := mapin(input)
		if err != nil {
			return nil, err
		}

		if block.Contains(target) {
			target.Drop()
			break
		}

		if block.OverlapAtStart(target) {
			target.Shorten(block)
			continue
		}

		if block.IsContainedIn(target) {
			slot, err := mapout(createSlotFrom(target, block))
			if err != nil {
				return nil, err
			}
			slots[j] = slot
			j++
			target.Shorten(block)
			continue
		}

		if block.OverlapAtEnd(target) {
			slot, err := mapout(createSlotFrom(target, block))
			if err != nil {
				return nil, err
			}
			slots[j] = slot
			j++
			target.Drop()
			break
		}
	}

	if target.Remain() {
		slot, err := mapout(target.ToSlot())
		if err != nil {
			return nil, err
		}
		slots[j] = slot
		j++
	}
	return slots[:j], nil
}

// It returns a list of available time slots.
// Use this when passing and returning the pre-defined struct.
func Find(blocks []*Block, span *Span) []*Slot {
	sorter := func(i, j int, blocks []*Block) bool {
		return blocks[i].Start().Before(blocks[j].Start())
	}
	mapin := func(b *Block) (*Block, error) {
		return b, nil
	}
	mapout := func(s *Slot) (*Slot, error) {
		return s, nil
	}
	r, _ := FindWithMapper(blocks, span, sorter, mapin, mapout)
	return r
}
