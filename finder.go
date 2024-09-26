package goslotify

import (
	"sort"
)

// Sort your struct in chronological order.
type SortFunc[I any] func(int, int, []I) bool

// Map your struct to a Block.
type MapInFunc[I any] func(I) (*Block, error)

// Map the Slot to your struct.
type MapOutFunc[O any] func(*Slot) (O, error)

// Filter your struct in your condition.
type FilterFunc[O any] func(O) bool

// Calculate available time slots (Slot). Provide the scheduled block (Block) and the target period (Span).
// Use this when passing and returning your struct.
func FindWithMapper[I, O any](inputs []I, span *Span, sorter SortFunc[I], mapin MapInFunc[I], mapout MapOutFunc[O], filter FilterFunc[O]) ([]O, error) {
	if span == nil || !span.Remain() {
		return []O{}, nil
	}

	target := span.Clone()
	if len(inputs) == 0 {
		slot, err := mapout(target.ToSlot())
		if err != nil {
			return nil, err
		}
		return []O{slot}, nil
	}

	sort.Slice(inputs, func(i, j int) bool {
		return sorter(i, j, inputs)
	})

	j := 0
	slots := make([]O, len(inputs)+1)
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
			target.Shorten(block)
			if filter(slot) {
				continue
			}
			slots[j] = slot
			j++
			continue
		}

		if block.OverlapAtEnd(target) {
			slot, err := mapout(createSlotFrom(target, block))
			if err != nil {
				return nil, err
			}
			target.Drop()
			if filter(slot) {
				break
			}
			slots[j] = slot
			j++
			break
		}
	}

	if target.Remain() {
		slot, err := mapout(target.ToSlot())
		if err != nil {
			return nil, err
		}
		if filter(slot) {
			return slots[:j], nil
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
	filter := func(s *Slot) bool {
		return false
	}
	r, _ := FindWithMapper(blocks, span, sorter, mapin, mapout, filter)
	return r
}
