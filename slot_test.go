package timeslots_test

import (
	"fmt"
	"timeslots"
	"testing"
	"time"
)

func TestNewSlot(t *testing.T) {
	tests := []struct {
		name      string
		start     time.Time
		end       time.Time
		wantError bool
	}{
		{
			name:      "Valid slot",
			start:     now.Add(0 * time.Hour),
			end:       now.Add(1 * time.Hour),
			wantError: false,
		},
		{
			name:      "Invalid slot, start after end",
			start:     now.Add(1 * time.Hour),
			end:       now.Add(0 * time.Hour),
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			slot, err := timeslots.NewSlot(tt.start, tt.end)
			if (err != nil) != tt.wantError {
				t.Errorf("NewSlot() error = %v, wantError %v", err, tt.wantError)
			}
			if slot != nil {
				if slot.Start() != tt.start || slot.End() != tt.end {
					t.Errorf("NewSlot() = start %v, end %v; want start %v, end %v", slot.Start(), slot.End(), tt.start, tt.end)
				}
			}
		})
	}
}

func TestSlotEqual(t *testing.T) {
	slot1, _ := timeslots.NewSlot(
		now.Add(0*time.Hour),
		now.Add(8*time.Hour),
	)
	slot2, _ := timeslots.NewSlot(
		now.Add(0*time.Hour),
		now.Add(8*time.Hour),
	)
	slot3, _ := timeslots.NewSlot(
		now.Add(1*time.Hour),
		now.Add(9*time.Hour),
	)

	if !slot1.Equal(slot2) {
		t.Errorf("Slot.Equal() = false; want true")
	}

	if slot1.Equal(slot3) {
		t.Errorf("Slot.Equal() = true; want false")
	}
}

func TestSlotSmallerThan(t *testing.T) {
	slot1, _ := timeslots.NewSlot(
		now.Add(0*time.Hour),
		now.Add(8*time.Hour),
	)
	slot2, _ := timeslots.NewSlot(
		now.Add(1*time.Hour),
		now.Add(9*time.Hour),
	)

	if !slot1.SmallerThan(slot2) {
		t.Errorf("Slot.SmallerThan() = false; want true")
	}

	if slot2.SmallerThan(slot1) {
		t.Errorf("Slot.SmallerThan() = true; want false")
	}
}

func TestSlotString(t *testing.T) {
	start := now.Add(0 * time.Hour)
	end := now.Add(8 * time.Hour)
	slot, _ := timeslots.NewSlot(
		start,
		end,
	)

	want := fmt.Sprintf("%s, %s", start.Format(timeslots.TimeFormat), end.Format(timeslots.TimeFormat))
	got := slot.String()
	if got != want {
		t.Errorf("Slot.String() = %s; want %s", got, want)
	}
}
