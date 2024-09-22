package slotify

import (
	"fmt"
	"slotify/internal/slice"
	"time"
)

type Period interface {
	Start() time.Time
	End() time.Time
	String() string
}

func ToString[T Period](p []T) string {
	return slice.String(p)
}

func format[T Period](p T) string {
	return fmt.Sprintf("%s, %s", p.Start().String(), p.End().String())
}

func equal[S, T Period](p S, q T) bool {
	return p.Start().Equal(q.Start()) && p.End().Equal(q.End())
}
