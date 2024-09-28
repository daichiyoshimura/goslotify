package timeslots

import (
	"fmt"
	"strings"
	"time"
)

const TimeFormat = "2006-01-02 15:04:05"

// This is an interface representing a period. Block, Span, and Slot all implement the Period interface.
type Period interface {
	Start() time.Time
	End() time.Time
}

// Represents the start time and end time as strings.
func ToString[T Period](p []T) string {
	var builder strings.Builder
	for _, v := range p {
		builder.WriteString(format(v))
		builder.WriteString("\n")
	}
	return builder.String()
}

func format[T Period](p T) string {
	return fmt.Sprintf("%s, %s", p.Start().Format(TimeFormat), p.End().Format(TimeFormat))
}

func equal[S, T Period](p S, q T) bool {
	return p.Start().Equal(q.Start()) && p.End().Equal(q.End())
}
