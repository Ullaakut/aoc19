package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOrbitCounts(t *testing.T) {
	realInput, err := ioutil.ReadFile("../input.txt")
	require.NoError(t, err)

	tests := []struct {
		content string
		want    int
	}{
		{
			content: "COM)B\nB)C\nC)D\nD)E\nE)F\nB)G\nG)H\nD)I\nE)J\nJ)K\nK)L",
			want:    42,
		},
		{
			content: strings.TrimSpace(string(realInput)),
			want:    223251,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint("test_", i), func(t *testing.T) {
			got := solve(test.content)

			assert.Equal(t, test.want, got)
		})
	}
}
