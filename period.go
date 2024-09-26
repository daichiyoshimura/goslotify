package goslotify

import (
	"fmt"
	"strings"
	"time"
)

// This is an interface representing a period. Block, Span, and Slot all implement the Period interface.
type Period interface {
	Start() time.Time
	End() time.Time
}

// Represents the start time and end time as strings.
func ToString[T Period](p []T) string {
	var builder strings.Builder
	for _, v := range p {
		builder.WriteString(fmt.Sprintf("%s, %s", v.Start().String(), v.End().String()))
		builder.WriteString("\n")
	}
	return builder.String()
}

func format[T Period](p T) string {
	return fmt.Sprintf("%s, %s", p.Start().String(), p.End().String())
}

func equal[S, T Period](p S, q T) bool {
	return p.Start().Equal(q.Start()) && p.End().Equal(q.End())
}
