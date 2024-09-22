package slotify

import "sort"

// T is input type. this func sorts inputs internally
type SortFunc[T any] func(int, int, []T) bool

// T is input type. this func convert input to *Block internally
type MapInFunc[T any] func(T) (*Block, error)

// T is output type. this func convert *Slot to output internally
type MapOutFunc[T any] func(*Slot) (T, error)

// It returns a list of available time slots with converting your struct.
func FindWithMapper[S, T any](inputs []S, span *Span, sorter SortFunc[S], mapin MapInFunc[S], mapout MapOutFunc[T]) ([]T, error) {
	slots := []T{}
	if span == nil {
		return slots, nil
	}

	target := span.Clone()
	if len(inputs) == 0 {
		slot, err := mapout(target.ToSlot())
		if err != nil {
			return nil, err
		}
		slots = append(slots, slot)
		return slots, nil
	}

	sort.Slice(inputs, func(i, j int) bool {
		return sorter(i, j, inputs)
	})

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
			slots = append(slots, slot)
			target.Shorten(block)
			continue
		}

		if block.OverlapAtEnd(target) {
			slot, err := mapout(createSlotFrom(target, block))
			if err != nil {
				return nil, err
			}
			slots = append(slots, slot)
			target.Drop()
			break
		}
	}

	if target.Remain() {
		slot, err := mapout(target.ToSlot())
		if err != nil {
			return nil, err
		}
		slots = append(slots, slot)
	}
	return slots, nil
}

// It returns a list of available time slots.
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
