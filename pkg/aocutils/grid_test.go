package aocutils

import (
	"bytes"
	"testing"

	"github.com/Ullaakut/disgo"
	"github.com/stretchr/testify/assert"
)

const exampleGrid = `
/->-\        
|   |  /----\
| /-+--+-\  |
| | |  | v  |
\-+-/  \-+--/
  \------/   `

// TestGridDisplay validates that given a string containing a grid,
// NewGrid parses it correctly and Display displays it properly.
func TestGridDisplay(t *testing.T) {
	output := &bytes.Buffer{}
	disgo.SetTerminalOptions(disgo.WithDefaultOutput(output))

	g := NewGrid(exampleGrid)

	g.Display()

	assert.Equal(t, exampleGrid, output.String())
}
