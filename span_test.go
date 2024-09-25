package goslotify_test

import (
	"fmt"
	"goslotify"
	"testing"
	"time"
)

func TestNewSpan(t *testing.T) {
	tests := []struct {
		name      string
		start     time.Time
		end       time.Time
		wantError bool
	}{
		{
			name:      "Valid span",
			start:     now.Add(0 * time.Hour),
			end:       now.Add(8 * time.Hour),
			wantError: false,
		},
		{
			name:      "Invalid span, start after end",
			start:     now.Add(8 * time.Hour),
			end:       now.Add(0 * time.Hour),
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			span, err := goslotify.NewSpan(tt.start, tt.end)
			if (err != nil) != tt.wantError {
				t.Errorf("NewSpan() error = %v, wantError %v", err, tt.wantError)
			}
			if span != nil {
				if span.Start() != tt.start || span.End() != tt.end {
					t.Errorf("NewSpan() = start %v, end %v; want start %v, end %v", span.Start(), span.End(), tt.start, tt.end)
				}
			}
		})
	}
}

func TestSpanClone(t *testing.T) {
	span, _ := goslotify.NewSpan(
		now.Add(0*time.Hour),
		now.Add(8*time.Hour),
	)
	clone := span.Clone()

	if !clone.Start().Equal(span.Start()) || !clone.End().Equal(span.End()) {
		t.Errorf("Clone() = start %v, end %v; want start %v, end %v", clone.Start(), clone.End(), span.Start(), span.End())
	}
}

func TestSpanRemain(t *testing.T) {
	tests := []struct {
		name   string
		start  time.Time
		end    time.Time
		expect bool
	}{
		{
			name:   "Remaining period",
			start:  now.Add(0 * time.Hour),
			end:    now.Add(8 * time.Hour),
			expect: true,
		},
		{
			name:   "No remaining period",
			start:  now.Add(0 * time.Hour),
			end:    now.Add(0 * time.Hour),
			expect: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			span, _ := goslotify.NewSpan(tt.start, tt.end)
			if span.Remain() != tt.expect {
				t.Errorf("Remain() = %v, want %v", span.Remain(), tt.expect)
			}
		})
	}
}

func TestSpanShorten(t *testing.T) {
	span, _ := goslotify.NewSpan(
		now.Add(0*time.Hour),
		now.Add(8*time.Hour),
	)
	block := goslotify.NewBlockWithoutValidating(
		now.Add(1*time.Hour),
		now.Add(6*time.Hour),
	)

	span.Shorten(block)

	expectedStart := block.End()

	if !span.Start().Equal(expectedStart) {
		t.Errorf("Shorten() = start %v, want %v", span.Start(), expectedStart)
	}
}

func TestSpanDrop(t *testing.T) {
	span, _ := goslotify.NewSpan(
		now.Add(0*time.Hour),
		now.Add(8*time.Hour),
	)

	span.Drop()

	if !span.Start().Equal(span.End()) {
		t.Errorf("Drop() = start %v, end %v; want start and end to be equal", span.Start(), span.End())
	}
}

func TestSpanToSlot(t *testing.T) {
	span, _ := goslotify.NewSpan(
		now.Add(0*time.Hour),
		now.Add(8*time.Hour),
	)

	slot := span.ToSlot()

	if !slot.Start().Equal(span.Start()) || !slot.End().Equal(span.End()) {
		t.Errorf("ToSlot() = start %v, end %v; want start %v, end %v", slot.Start(), slot.End(), span.Start(), span.End())
	}
}

func TestSpanString(t *testing.T) {
	start := now.Add(0 * time.Hour)
	end := now.Add(8 * time.Hour)
	span, _ := goslotify.NewSpan(
		start,
		end,
	)

	want := fmt.Sprintf("%s, %s", start.String(), end.String())
	got := span.String()
	if got != want {
		t.Errorf("Slot.String() = %s; want %s", got, want)
	}
}
