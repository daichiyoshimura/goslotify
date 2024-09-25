package goslotify_test

import (
	"fmt"
	"goslotify"
	"testing"
	"time"
)

type MockPeriod struct {
	start time.Time
	end   time.Time
}

func (m MockPeriod) Start() time.Time {
	return m.start
}

func (m MockPeriod) End() time.Time {
	return m.end
}

func (m MockPeriod) String() string {
	return m.start.String() + " - " + m.end.String()
}

func TestToString(t *testing.T) {
	tests := []struct {
		name   string
		input  []goslotify.Period
		output string
	}{
		{
			name: "Multiple periods",
			input: []goslotify.Period{
				MockPeriod{start: now.Add(0 * time.Hour), end: now.Add(1 * time.Hour)},
				MockPeriod{start: now.Add(2 * time.Hour), end: now.Add(3 * time.Hour)},
			},
			output: fmt.Sprintf(
				"%s - %s\n%s - %s\n",
				now.Add(0*time.Hour).String(),
				now.Add(1*time.Hour).String(),
				now.Add(2*time.Hour).String(),
				now.Add(3*time.Hour).String(),
			),
		},
		{
			name:   "Empty periods",
			input:  []goslotify.Period{},
			output: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := goslotify.ToString(tt.input)
			if result != tt.output {
				t.Errorf("expected: %v, got: %v", tt.output, result)
			}
		})
	}
}
