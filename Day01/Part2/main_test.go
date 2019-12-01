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
			input: "1969",
			want:  966,
		},
		{
			input: "100756",
			want:  50346,
		},
		{
			input: "100756\n1969",
			want:  50346 + 966,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint("test_", i), func(t *testing.T) {
			got := solve(test.input)

			assert.Equal(t, test.want, got)
		})
	}
}
