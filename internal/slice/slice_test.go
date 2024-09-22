package slice_test

import (
	"reflect"
	"slotify/internal/slice"
	"testing"
)

type mockStringer struct {
	value string
}

func (m mockStringer) String() string {
	return m.value
}

type mockComparable struct {
	value int
}

func (m mockComparable) Equal(other mockComparable) bool {
	return m.value == other.value
}

func (m mockComparable) SmallerThan(other mockComparable) bool {
	return m.value < other.value
}

func TestShiftLeft(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		index  int
		output []int
	}{
		{
			name:   "Shift from start",
			input:  []int{1, 2, 3, 4, 5},
			index:  0,
			output: []int{2, 3, 4, 5},
		},
		{
			name:   "Shift from middle",
			input:  []int{1, 2, 3, 4, 5},
			index:  2,
			output: []int{1, 2, 4, 5},
		},
		{
			name:   "Shift from end",
			input:  []int{1, 2, 3, 4, 5},
			index:  4,
			output: []int{1, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := slice.ShiftLeft(tt.index, tt.input)
			if !reflect.DeepEqual(result, tt.output) {
				t.Errorf("expected: %v, got: %v", tt.output, result)
			}
		})
	}
}

func TestClone(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output []int
	}{
		{
			name:   "Clone slice",
			input:  []int{1, 2, 3, 4, 5},
			output: []int{1, 2, 3, 4, 5},
		},
		{
			name:   "Empty slice",
			input:  []int{},
			output: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := slice.Clone(tt.input)
			if !reflect.DeepEqual(result, tt.output) {
				t.Errorf("expected: %v, got: %v", tt.output, result)
			}
		})
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		name   string
		input  []mockStringer
		output string
	}{
		{
			name: "Stringify slice",
			input: []mockStringer{
				{value: "one"},
				{value: "two"},
				{value: "three"},
			},
			output: "one\ntwo\nthree\n",
		},
		{
			name:   "Empty slice",
			input:  []mockStringer{},
			output: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := slice.String(tt.input)
			if result != tt.output {
				t.Errorf("expected: %v, got: %v", tt.output, result)
			}
		})
	}
}

func TestEqual(t *testing.T) {
	tests := []struct {
		name   string
		a, b   []mockComparable
		output bool
	}{
		{
			name: "Equal slices",
			a: []mockComparable{
				{value: 1},
				{value: 2},
				{value: 3},
			},
			b: []mockComparable{
				{value: 1},
				{value: 2},
				{value: 3},
			},
			output: true,
		},
		{
			name: "Different slices",
			a: []mockComparable{
				{value: 1},
				{value: 2},
			},
			b: []mockComparable{
				{value: 1},
				{value: 3},
			},
			output: false,
		},
		{
			name: "Different length slices",
			a: []mockComparable{
				{value: 1},
			},
			b: []mockComparable{
				{value: 1},
				{value: 3},
			},
			output: false,
		},
		{
			name:   "Empty slices",
			a:      []mockComparable{},
			b:      []mockComparable{},
			output: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := slice.Equal(tt.a, tt.b)
			if result != tt.output {
				t.Errorf("expected: %v, got: %v", tt.output, result)
			}
		})
	}
}
