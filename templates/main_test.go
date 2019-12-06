package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolve(t *testing.T) {
	realInput, err := ioutil.ReadFile("../input.txt")
	require.NoError(t, err)

	tests := []struct {
		content string
		want    int
	}{
		{
			content: strings.TrimSpace(string(realInput)),
			want:    4242,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint("test_", i), func(t *testing.T) {
			got := solve(test.content)

			assert.Equal(t, test.want, got)
		})
	}
}
