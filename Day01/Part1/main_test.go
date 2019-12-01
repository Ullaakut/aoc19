package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFuelRequirementComputation(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: "12",
			want:  2,
		},
		{
			input: "14",
			want:  2,
		},
		{
			input: "1969",
			want:  654,
		},
		{
			input: "100756",
			want:  33583,
		},
		{
			input: "100756\n1969",
			want:  33583 + 654,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint("test_", i), func(t *testing.T) {
			got := solve(test.input)

			assert.Equal(t, test.want, got)
		})
	}
}
