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
			content: strings.TrimSpace(string(realInput)),
			input:   1,
			want:    16209841,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint("test_", i), func(t *testing.T) {
			got := solve(test.content, test.input)

			assert.Equal(t, test.want, aocutils.Atoi(got))
		})
	}
}
