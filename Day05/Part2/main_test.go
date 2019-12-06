package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/Ullaakut/aoc19/pkg/aocutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOpcodeDiagnostics(t *testing.T) {
	realInput, err := ioutil.ReadFile("../input.txt")
	require.NoError(t, err)

	tests := []struct {
		content string
		input   int
		want    int
	}{
		{
			content: `3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99`,
			input:   7,
			want:    999,
		},
		{
			content: `3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99`,
			input:   8,
			want:    1000,
		},
		{
			content: `3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99`,
			input:   9,
			want:    1001,
		},
		{
			content: strings.TrimSpace(string(realInput)),
			input:   5,
			want:    8834787,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint("test_", i), func(t *testing.T) {
			got := solve(test.content, test.input)

			assert.Equal(t, test.want, aocutils.Atoi(got))
		})
	}
}
