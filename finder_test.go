package goslotify_test

import (
	"goslotify"
	"goslotify/internal/slice"
	"testing"
)

type testCase struct {
	name   string
	blocks []*goslotify.Block
	search *goslotify.Span
	want   []*goslotify.Slot
}

func testCases(h *TestingHelper) []testCase {
	return []testCase{
		{
			name:   "No blocks",
			blocks: []*goslotify.Block{},
			search: h.Span(0, 1),
			want:   []*goslotify.Slot{h.Slot(0, 1)},
		},
		{
			name:   "Nil span",
			blocks: []*goslotify.Block{},
			search: nil,
			want:   []*goslotify.Slot{},
		},
		{
			name:   "Empty span",
			blocks: []*goslotify.Block{h.Block(0, 1)},
			search: h.Span(0, 0),
			want:   []*goslotify.Slot{},
		},
		{
			name:   "One block before slot",
			blocks: []*goslotify.Block{h.Block(-2, -1)},
			search: h.Span(0, 8),
			want:   []*goslotify.Slot{h.Slot(0, 8)},
		},
		{
			name:   "One block before slot boundary",
			blocks: []*goslotify.Block{h.Block(-1, 0)},
			search: h.Span(0, 8),
			want:   []*goslotify.Slot{h.Slot(0, 8)},
		},
		{
			name:   "One block with overlap at start",
			blocks: []*goslotify.Block{h.Block(-1, 1)},
			search: h.Span(0, 8),
			want:   []*goslotify.Slot{h.Slot(1, 8)},
		},
		{
			name:   "One block with overlap at start boundary",
			blocks: []*goslotify.Block{h.Block(0, 1)},
			search: h.Span(0, 8),
			want:   []*goslotify.Slot{h.Slot(1, 8)},
		},
		{
			name:   "One block is contained in slot",
			blocks: []*goslotify.Block{h.Block(1, 5)},
			search: h.Span(0, 8),
			want:   []*goslotify.Slot{h.Slot(0, 1), h.Slot(5, 8)},
		},
		{
			name:   "One block is contained in slot boundary (= One block contains slot boundary)",
			blocks: []*goslotify.Block{h.Block(0, 8)},
			search: h.Span(0, 8),
			want:   []*goslotify.Slot{},
		},
		{
			name:   "One block contains slot",
			blocks: []*goslotify.Block{h.Block(-1, 9)},
			search: h.Span(0, 8),
			want:   []*goslotify.Slot{},
		},
		{
			name:   "One block with overlap at end boundary",
			blocks: []*goslotify.Block{h.Block(3, 8)},
			search: h.Span(0, 8),
			want:   []*goslotify.Slot{h.Slot(0, 3)},
		},
		{
			name:   "One block with overlap at end",
			blocks: []*goslotify.Block{h.Block(3, 9)},
			search: h.Span(0, 8),
			want:   []*goslotify.Slot{h.Slot(0, 3)},
		},
		{
			name:   "One block after slot boundary",
			blocks: []*goslotify.Block{h.Block(8, 10)},
			search: h.Span(0, 8),
			want:   []*goslotify.Slot{h.Slot(0, 8)},
		},
		{
			name:   "One block after slot",
			blocks: []*goslotify.Block{h.Block(9, 10)},
			search: h.Span(0, 8),
			want:   []*goslotify.Slot{h.Slot(0, 8)},
		},
		{
			name:   "Two block are contained in slot",
			blocks: []*goslotify.Block{h.Block(1, 2), h.Block(6, 7)},
			search: h.Span(0, 8),
			want:   []*goslotify.Slot{h.Slot(0, 1), h.Slot(2, 6), h.Slot(7, 8)},
		},
		{
			name:   "Two block overlaps each other are contained in slot",
			blocks: []*goslotify.Block{h.Block(1, 4), h.Block(2, 5)},
			search: h.Span(0, 8),
			want:   []*goslotify.Slot{h.Slot(0, 1), h.Slot(5, 8)},
		},
		{
			name:   "Huge blocks overlap each other are contained in slot",
			blocks: h.HugeBlocks(1, 799),
			search: h.Span(0, 799),
			want:   h.HugeSlots(0, 799),
		},
	}
}

func TestFind(t *testing.T) {
	h := NewTestingHelper(now)
	tests := testCases(h)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := goslotify.Find(tt.blocks, tt.search)
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
				got := goslotify.Find(tt.blocks, tt.search)
				if !slice.Equal(got, tt.want) {
					b.Errorf("got: %v, want: %v", slice.String(got), slice.String(tt.want))
				}
			})
		}
	}
}
