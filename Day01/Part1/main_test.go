package main

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFuelRequirementComputation(t *testing.T) {
	realInput, err := ioutil.ReadFile("../input.txt")
	require.NoError(t, err)

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
		{
			input: string(realInput),
			want:  3563458,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint("test_", i), func(t *testing.T) {
			got := solve(test.input)

			assert.Equal(t, test.want, got)
		})
	}
}
