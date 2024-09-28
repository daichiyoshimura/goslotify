package timeslots_test

import (
	"timeslots"
	"timeslots/internal/slice"
	"testing"
	"time"
)

type testCase struct {
	name   string
	blocks []*timeslots.Block
	search *timeslots.Span
	filter timeslots.FilterFunc[*timeslots.Slot]
	want   []*timeslots.Slot
}

func testCases(h *TestingHelper) []testCase {
	return []testCase{
		{
			name:   "No blocks",
			blocks: []*timeslots.Block{},
			search: h.Span(0, 1),
			filter: func(s *timeslots.Slot) bool {
				return false
			},
			want: []*timeslots.Slot{h.Slot(0, 1)},
		},
		{
			name:   "Nil span",
			blocks: []*timeslots.Block{},
			search: nil,
			filter: func(s *timeslots.Slot) bool {
				return false
			},
			want: []*timeslots.Slot{},
		},
		{
			name:   "Empty span",
			blocks: []*timeslots.Block{h.Block(0, 1)},
			search: h.Span(0, 0),
			filter: func(s *timeslots.Slot) bool {
				return false
			},
			want: []*timeslots.Slot{},
		},
		{
			name:   "One block before slot",
			blocks: []*timeslots.Block{h.Block(-2, -1)},
			search: h.Span(0, 8),
			filter: func(s *timeslots.Slot) bool {
				return false
			},
			want: []*timeslots.Slot{h.Slot(0, 8)},
		},
		{
			name:   "One block before slot boundary",
			blocks: []*timeslots.Block{h.Block(-1, 0)},
			search: h.Span(0, 8),
			filter: func(s *timeslots.Slot) bool {
				return false
			},
			want: []*timeslots.Slot{h.Slot(0, 8)},
		},
		{
			name:   "One block with overlap at start",
			blocks: []*timeslots.Block{h.Block(-1, 1)},
			search: h.Span(0, 8),
			filter: func(s *timeslots.Slot) bool {
				return false
			},
			want: []*timeslots.Slot{h.Slot(1, 8)},
		},
		{
			name:   "One block with overlap at start boundary",
			blocks: []*timeslots.Block{h.Block(0, 1)},
			search: h.Span(0, 8),
			filter: func(s *timeslots.Slot) bool {
				return false
			},
			want: []*timeslots.Slot{h.Slot(1, 8)},
		},
		{
			name:   "One block is contained in slot",
			blocks: []*timeslots.Block{h.Block(1, 5)},
			search: h.Span(0, 8),
			filter: func(s *timeslots.Slot) bool {
				return false
			},
			want: []*timeslots.Slot{h.Slot(0, 1), h.Slot(5, 8)},
		},
		{
			name:   "One block is contained in slot boundary (= One block contains slot boundary)",
			blocks: []*timeslots.Block{h.Block(0, 8)},
			search: h.Span(0, 8),
			filter: func(s *timeslots.Slot) bool {
				return false
			},
			want: []*timeslots.Slot{},
		},
		{
			name:   "One block contains slot",
			blocks: []*timeslots.Block{h.Block(-1, 9)},
			search: h.Span(0, 8),
			filter: func(s *timeslots.Slot) bool {
				return false
			},
			want: []*timeslots.Slot{},
		},
		{
			name:   "One block with overlap at end boundary",
			blocks: []*timeslots.Block{h.Block(3, 8)},
			search: h.Span(0, 8),
			filter: func(s *timeslots.Slot) bool {
				return false
			},
			want: []*timeslots.Slot{h.Slot(0, 3)},
		},
		{
			name:   "One block with overlap at end",
			blocks: []*timeslots.Block{h.Block(3, 9)},
			search: h.Span(0, 8),
			filter: func(s *timeslots.Slot) bool {
				return false
			},
			want: []*timeslots.Slot{h.Slot(0, 3)},
		},
		{
			name:   "One block after slot boundary",
			blocks: []*timeslots.Block{h.Block(8, 10)},
			search: h.Span(0, 8),
			filter: func(s *timeslots.Slot) bool {
				return false
			},
			want: []*timeslots.Slot{h.Slot(0, 8)},
		},
		{
			name:   "One block after slot",
			blocks: []*timeslots.Block{h.Block(9, 10)},
			search: h.Span(0, 8),
			filter: func(s *timeslots.Slot) bool {
				return false
			},
			want: []*timeslots.Slot{h.Slot(0, 8)},
		},
		{
			name:   "Two block are contained in slot",
			blocks: []*timeslots.Block{h.Block(1, 2), h.Block(6, 7)},
			search: h.Span(0, 8),
			filter: func(s *timeslots.Slot) bool {
				return false
			},
			want: []*timeslots.Slot{h.Slot(0, 1), h.Slot(2, 6), h.Slot(7, 8)},
		},
		{
			name:   "Two block overlaps each other are contained in slot",
			blocks: []*timeslots.Block{h.Block(1, 4), h.Block(2, 5)},
			search: h.Span(0, 8),
			filter: func(s *timeslots.Slot) bool {
				return false
			},
			want: []*timeslots.Slot{h.Slot(0, 1), h.Slot(5, 8)},
		},
		{
			name:   "Huge blocks overlap each other are contained in slot",
			blocks: h.HugeBlocks(1, 799),
			search: h.Span(0, 799),
			filter: func(s *timeslots.Slot) bool {
				return false
			},
			want: h.HugeSlots(0, 799),
		},
		{
			name:   "One block is contained in slot with filter",
			blocks: []*timeslots.Block{h.Block(1, 5)},
			search: h.Span(0, 8),
			filter: func(s *timeslots.Slot) bool {
				return s.End().Sub(s.Start()) < (2 * time.Hour)
			},
			want: []*timeslots.Slot{h.Slot(5, 8)},
		},
		{
			name:   "One block with overlap at end with filter",
			blocks: []*timeslots.Block{h.Block(3, 9)},
			search: h.Span(0, 8),
			filter: func(s *timeslots.Slot) bool {
				return s.End().Sub(s.Start()) < (4 * time.Hour)
			},
			want: []*timeslots.Slot{},
		},
		{
			name:   "One block after slot with filter",
			blocks: []*timeslots.Block{h.Block(9, 10)},
			search: h.Span(0, 8),
			filter: func(s *timeslots.Slot) bool {
				return s.End().Sub(s.Start()) > (2 * time.Hour)
			},
			want: []*timeslots.Slot{},
		},
	}
}

func TestFind(t *testing.T) {
	h := NewTestingHelper(now)
	tests := testCases(h)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := timeslots.Find(tt.blocks, tt.search, timeslots.WithFilter(tt.filter))
			if !slice.Equal(got, tt.want) {
				t.Errorf("got: %v, want: %v", slice.String(got), slice.String(tt.want))
			}
		})
	}
}

func TestFindWithMapper(t *testing.T) {
	h := NewTestingHelper(now)
	tests := testCases(h)

	mapIn := func(b *timeslots.Block) *timeslots.Block {
		return b
	}

	mapOut := func(s *timeslots.Slot) *timeslots.Slot {
		return s
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := timeslots.FindWithMapper(tt.blocks, tt.search, mapIn, mapOut, timeslots.WithFilter(tt.filter))
			if !slice.Equal(got, tt.want) {
				t.Errorf("got: %v, want: %v", slice.String(got), slice.String(tt.want))
			}
		})
	}
}

func BenchmarkFind(b *testing.B) {
	h := NewTestingHelper(now)
	tests := testCases(h)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			b.Run(tt.name, func(b *testing.B) {
				got := timeslots.Find(tt.blocks, tt.search, timeslots.WithFilter(tt.filter))
				if !slice.Equal(got, tt.want) {
					b.Errorf("got: %v, want: %v", slice.String(got), slice.String(tt.want))
				}
			})
		}
	}
}

func BenchmarkFindWithMapper(b *testing.B) {
	h := NewTestingHelper(now)
	tests := testCases(h)
	b.ResetTimer()

	mapIn := func(b *timeslots.Block) *timeslots.Block {
		return b
	}

	mapOut := func(s *timeslots.Slot) *timeslots.Slot {
		return s
	}

	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			b.Run(tt.name, func(t *testing.B) {
				got := timeslots.FindWithMapper(tt.blocks, tt.search, mapIn, mapOut, timeslots.WithFilter(tt.filter))
				if !slice.Equal(got, tt.want) {
					t.Errorf("got: %v, want: %v", slice.String(got), slice.String(tt.want))
				}
			})
		}
	}
}
