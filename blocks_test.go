package goslotify_test

import (
	"fmt"
	"goslotify"
	"testing"
	"time"
)

func TestNewBlock(t *testing.T) {

	tests := []struct {
		name    string
		start   time.Time
		end     time.Time
		wantErr bool
	}{
		{
			name:    "Valid time range",
			start:   now.Add(0 * time.Hour),
			end:     now.Add(8 * time.Hour),
			wantErr: false,
		},
		{
			name:    "Invalid time range (start after end)",
			start:   now.Add(8 * time.Hour),
			end:     now.Add(0 * time.Hour),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := goslotify.NewBlock(tt.start, tt.end)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewBlock() error = %v, wantErr = %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewBlocks(t *testing.T) {

	type Input struct {
		start time.Time
		end   time.Time
	}

	tests := []struct {
		name    string
		mapper  func(input Input) (*goslotify.Block, error)
		want    []*goslotify.Block
		wantErr bool
	}{
		{
			name: "Create multiple blocks",
			mapper: func(input Input) (*goslotify.Block, error) {
				return goslotify.NewBlock(input.start, input.end)
			},
			want: []*goslotify.Block{
				goslotify.NewBlockWithoutValidating(now.Add(1*time.Hour), now.Add(3*time.Hour)),
				goslotify.NewBlockWithoutValidating(now.Add(2*time.Hour), now.Add(4*time.Hour)),
			},
			wantErr: false,
		},
		{
			name: "Broken mapper",
			mapper: func(input Input) (*goslotify.Block, error) {
				return nil, fmt.Errorf("broken")
			},
			want:    []*goslotify.Block{},
			wantErr: true,
		},
	}

	inputs := []Input{
		{start: now.Add(1 * time.Hour), end: now.Add(3 * time.Hour)},
		{start: now.Add(2 * time.Hour), end: now.Add(4 * time.Hour)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := goslotify.NewBlocks(inputs, tt.mapper)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewBlock() error = %v, wantErr = %v", err, tt.wantErr)
			}
			if len(got) != len(tt.want) {
				t.Errorf("Expected %d blocks, got %d", len(tt.want), len(got))
			}
		})
	}
}

func TestContains(t *testing.T) {

	block := goslotify.NewBlockWithoutValidating(
		now.Add(0*time.Hour),
		now.Add(8*time.Hour),
	)

	tests := []struct {
		name  string
		other goslotify.Period
		want  bool
	}{
		{
			name: "No overlap at end, before ther other start",
			other: goslotify.NewBlockWithoutValidating(
				now.Add(-2*time.Hour),
				now.Add(-1*time.Hour),
			),
			want: false,
		},
		{
			name: "No overlap at end, end is ther other start",
			other: goslotify.NewBlockWithoutValidating(
				now.Add(-1*time.Hour),
				now.Add(0*time.Hour),
			),
			want: false,
		},
		{
			name: "Overlap at end",
			other: goslotify.NewBlockWithoutValidating(
				now.Add(-1*time.Hour),
				now.Add(1*time.Hour),
			),
			want: false,
		},
		{
			name: "Overlap at end, end is the other end",
			other: goslotify.NewBlockWithoutValidating(
				now.Add(-1*time.Hour),
				now.Add(8*time.Hour),
			),
			want: false,
		},
		{
			name: "Overlap at end, start is the other start",
			other: goslotify.NewBlockWithoutValidating(
				now.Add(0*time.Hour),
				now.Add(1*time.Hour),
			),
			want: true,
		},
		{
			name: "Overlap at start",
			other: goslotify.NewBlockWithoutValidating(
				now.Add(7*time.Hour),
				now.Add(9*time.Hour),
			),
			want: false,
		},
		{
			name: "Overlap at start, end is the other end",
			other: goslotify.NewBlockWithoutValidating(
				now.Add(7*time.Hour),
				now.Add(8*time.Hour),
			),
			want: true,
		},
		{
			name: "Overlap at start, start is the other start",
			other: goslotify.NewBlockWithoutValidating(
				now.Add(0*time.Hour),
				now.Add(9*time.Hour),
			),
			want: false,
		},
		{
			name: "No overlap at start, start is ther other end",
			other: goslotify.NewBlockWithoutValidating(
				now.Add(8*time.Hour),
				now.Add(9*time.Hour),
			),
			want: false,
		},
		{
			name: "No overlap at start, after ther other end",
			other: goslotify.NewBlockWithoutValidating(
				now.Add(9*time.Hour),
				now.Add(10*time.Hour),
			),
			want: false,
		},
		{
			name: "Is contained in the other",
			other: goslotify.NewBlockWithoutValidating(
				now.Add(-1*time.Hour),
				now.Add(9*time.Hour),
			),
			want: false,
		},
		{
			name: "Contains in the other",
			other: goslotify.NewBlockWithoutValidating(
				now.Add(1*time.Hour),
				now.Add(7*time.Hour),
			),
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := block.Contains(tt.other); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsContainedIn(t *testing.T) {

	block := goslotify.NewBlockWithoutValidating(
		now.Add(0*time.Hour),
		now.Add(8*time.Hour),
	)

	tests := []struct {
		name  string
		other goslotify.Period
		want  bool
	}{
		{
			name: "No overlap at end, before ther other start",
			other: goslotify.NewBlockWithoutValidating(
				now.Add(-2*time.Hour),
				now.Add(-1*time.Hour),
			),
			want: false,
		},
		{
			name: "No overlap at end, end is ther other start",
			other: goslotify.NewBlockWithoutValidating(
				now.Add(-1*time.Hour),
				now.Add(0*time.Hour),
			),
			want: false,
		},
		{
			name: "Overlap at end",
			other: goslotify.NewBlockWithoutValidating(
				now.Add(-1*time.Hour),
				now.Add(1*time.Hour),
			),
			want: false,
		},
		{
			name: "Overlap at end, end is the other end",
			other: goslotify.NewBlockWithoutValidating(
				now.Add(-1*time.Hour),
				now.Add(8*time.Hour),
			),
			want: true,
		},
		{
			name: "Overlap at end, start is the other start",
			other: goslotify.NewBlockWithoutValidating(
				now.Add(0*time.Hour),
				now.Add(1*time.Hour),
			),
			want: false,
		},
		{
			name: "Overlap at start",
			other: goslotify.NewBlockWithoutValidating(
				now.Add(7*time.Hour),
				now.Add(9*time.Hour),
			),
			want: false,
		},
		{
			name: "Overlap at start, end is the other end",
			other: goslotify.NewBlockWithoutValidating(
				now.Add(7*time.Hour),
				now.Add(8*time.Hour),
			),
			want: false,
		},
		{
			name: "Overlap at start, start is the other start",
			other: goslotify.NewBlockWithoutValidating(
				now.Add(0*time.Hour),
				now.Add(9*time.Hour),
			),
			want: true,
		},
		{
			name: "No overlap at start, start is ther other end",
			other: goslotify.NewBlockWithoutValidating(
				now.Add(8*time.Hour),
				now.Add(9*time.Hour),
			),
			want: false,
		},
		{
			name: "No overlap at start, after ther other end",
			other: goslotify.NewBlockWithoutValidating(
				now.Add(9*time.Hour),
				now.Add(10*time.Hour),
			),
			want: false,
		},
		{
			name: "Is contained in the other",
			other: goslotify.NewBlockWithoutValidating(
				now.Add(-1*time.Hour),
				now.Add(9*time.Hour),
			),
			want: true,
		},
		{
			name: "Contains in the other",
			other: goslotify.NewBlockWithoutValidating(
				now.Add(1*time.Hour),
				now.Add(7*time.Hour),
			),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := block.IsContainedIn(tt.other); got != tt.want {
				t.Errorf("IsContainedIn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOverlapAtStart(t *testing.T) {

	block := goslotify.NewBlockWithoutValidating(
		now.Add(0*time.Hour),
		now.Add(8*time.Hour),
	)

	tests := []struct {
		name  string
		other goslotify.Period
		want  bool
	}{
		{
			name: "No overlap at end, before ther other start",
			other: goslotify.NewBlockWithoutValidating(
				now.Add(-2*time.Hour),
				now.Add(-1*time.Hour),
			),
			want: false,
		},
		{
			name: "No overlap at end, end is ther other start",
			other: goslotify.NewBlockWithoutValidating(
				now.Add(-1*time.Hour),
				now.Add(0*time.Hour),
			),
			want: false,
		},
		{
			name: "Overlap at end",
			other: goslotify.NewBlockWithoutValidating(
				now.Add(-1*time.Hour),
				now.Add(1*time.Hour),
			),
			want: false,
		},
		{
			name: "Overlap at end, end is the other end",
			other: goslotify.NewBlockWithoutValidating(
				now.Add(-1*time.Hour),
				now.Add(8*time.Hour),
			),
			want: false,
		},
		{
			name: "Overlap at end, start is the other start",
			other: goslotify.NewBlockWithoutValidating(
				now.Add(0*time.Hour),
				now.Add(1*time.Hour),
			),
			want: false,
		},
		{
			name: "Overlap at start",
			other: goslotify.NewBlockWithoutValidating(
				now.Add(7*time.Hour),
				now.Add(9*time.Hour),
			),
			want: true,
		},
		{
			name: "Overlap at start, end is the other end",
			other: goslotify.NewBlockWithoutValidating(
				now.Add(7*time.Hour),
				now.Add(8*time.Hour),
			),
			want: true,
		},
		{
			name: "Overlap at start, start is the other start",
			other: goslotify.NewBlockWithoutValidating(
				now.Add(0*time.Hour),
				now.Add(9*time.Hour),
			),
			want: true,
		},
		{
			name: "No overlap at start, start is ther other end",
			other: goslotify.NewBlockWithoutValidating(
				now.Add(8*time.Hour),
				now.Add(9*time.Hour),
			),
			want: false,
		},
		{
			name: "No overlap at start, after ther other end",
			other: goslotify.NewBlockWithoutValidating(
				now.Add(9*time.Hour),
				now.Add(10*time.Hour),
			),
			want: false,
		},
		{
			name: "Is contained in the other",
			other: goslotify.NewBlockWithoutValidating(
				now.Add(-1*time.Hour),
				now.Add(9*time.Hour),
			),
			want: false,
		},
		{
			name: "Contains in the other",
			other: goslotify.NewBlockWithoutValidating(
				now.Add(1*time.Hour),
				now.Add(7*time.Hour),
			),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := block.OverlapAtStart(tt.other); got != tt.want {
				t.Errorf("OverlapAtStart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOverlapAtEnd(t *testing.T) {

	block := goslotify.NewBlockWithoutValidating(
		now.Add(0*time.Hour),
		now.Add(8*time.Hour),
	)

	tests := []struct {
		name  string
		other goslotify.Period
		want  bool
	}{
		{
			name: "No overlap at end, before ther other start",
			other: goslotify.NewBlockWithoutValidating(
				now.Add(-2*time.Hour),
				now.Add(-1*time.Hour),
			),
			want: false,
		},
		{
			name: "No overlap at end, end is ther other start",
			other: goslotify.NewBlockWithoutValidating(
				now.Add(-1*time.Hour),
				now.Add(0*time.Hour),
			),
			want: false,
		},
		{
			name: "Overlap at end",
			other: goslotify.NewBlockWithoutValidating(
				now.Add(-1*time.Hour),
				now.Add(1*time.Hour),
			),
			want: true,
		},
		{
			name: "Overlap at end, end is the other end",
			other: goslotify.NewBlockWithoutValidating(
				now.Add(-1*time.Hour),
				now.Add(8*time.Hour),
			),
			want: true,
		},
		{
			name: "Overlap at end, start is the other start",
			other: goslotify.NewBlockWithoutValidating(
				now.Add(0*time.Hour),
				now.Add(1*time.Hour),
			),
			want: true,
		},
		{
			name: "Overlap at start",
			other: goslotify.NewBlockWithoutValidating(
				now.Add(7*time.Hour),
				now.Add(9*time.Hour),
			),
			want: false,
		},
		{
			name: "Overlap at start, end is the other end",
			other: goslotify.NewBlockWithoutValidating(
				now.Add(7*time.Hour),
				now.Add(8*time.Hour),
			),
			want: false,
		},
		{
			name: "Overlap at start, start is the other start",
			other: goslotify.NewBlockWithoutValidating(
				now.Add(0*time.Hour),
				now.Add(9*time.Hour),
			),
			want: false,
		},
		{
			name: "No overlap at start, start is ther other end",
			other: goslotify.NewBlockWithoutValidating(
				now.Add(8*time.Hour),
				now.Add(9*time.Hour),
			),
			want: false,
		},
		{
			name: "No overlap at start, after ther other end",
			other: goslotify.NewBlockWithoutValidating(
				now.Add(9*time.Hour),
				now.Add(10*time.Hour),
			),
			want: false,
		},
		{
			name: "Is contained in the other",
			other: goslotify.NewBlockWithoutValidating(
				now.Add(-1*time.Hour),
				now.Add(9*time.Hour),
			),
			want: false,
		},
		{
			name: "Contains in the other",
			other: goslotify.NewBlockWithoutValidating(
				now.Add(1*time.Hour),
				now.Add(7*time.Hour),
			),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := block.OverlapAtEnd(tt.other); got != tt.want {
				t.Errorf("OverlapAtEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBlockString(t *testing.T) {

	start := now.Add(0 * time.Hour)
	end := now.Add(8 * time.Hour)
	span, _ := goslotify.NewBlock(
		start,
		end,
	)

	want := fmt.Sprintf("%s, %s", start.String(), end.String())
	got := span.String()
	if got != want {
		t.Errorf("Slot.String() = %s; want %s", got, want)
	}
}
