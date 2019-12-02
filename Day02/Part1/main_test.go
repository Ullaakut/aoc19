package main

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOpCodeComputation(t *testing.T) {
	realInput, err := ioutil.ReadFile("../input.txt")
	require.NoError(t, err)

	tests := []struct {
		input string
		want  int
	}{
		{
			input: "1,9,10,3,2,3,11,0,99,30,40,50",
			want:  3500,
		},
		{
			input: "1,0,0,0,99",
			want:  2,
		},
		{
			input: "2,3,0,3,99",
			want:  2,
		},
		{
			input: "2,4,4,5,99,0",
			want:  2,
		},
		{
			input: "1,1,1,4,99,5,6,0,99",
			want:  30,
		},
		{
			input: "1,9,10,3,2,3,11,0,99,30,40,50",
			want:  3500,
		},
		{
			input: string(realInput),
			want:  3562624,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint("test_", i), func(t *testing.T) {
			got := solve(test.input)

			assert.Equal(t, test.want, got)
		})
	}
}
