package slice

import (
	"sort"
	"strings"
)

func ShiftLeft[T any](i int, s []T) []T {
	lens := len(s)
	for j := i; j < lens-1; j++ {
		s[j] = s[j+1]
	}
	return s[:lens-1]
}

func Clone[T any](s []T) []T {
	c := make([]T, len(s))
	copy(c, s)
	return c
}

type Stringer interface {
	String() string
}

func String[T Stringer](s []T) string {
	var builder strings.Builder
	for _, v := range s {
		builder.WriteString(v.String())
		builder.WriteString("\n")
	}
	return builder.String()
}

type Comparable[T any] interface {
	Equal(T) bool
	SmallerThan(T) bool
}

func Equal[T Comparable[T]](a, b []T) bool {
	if len(a) == 0 && len(b) == 0 {
		return true
	}

	if len(a) != len(b) {
		return false
	}

	sort.Slice(a, func(i, j int) bool {
		return a[i].SmallerThan(b[j])
	})

	for i, v := range a {
		if !v.Equal(b[i]) {
			return false
		}
	}
	return true
}
