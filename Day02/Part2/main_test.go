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
			input: string(realInput),
			want:  8298,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint("test_", i), func(t *testing.T) {
			got := solve(test.input)

			assert.Equal(t, test.want, got)
		})
	}
}
