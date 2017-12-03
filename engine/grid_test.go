package engine_test

import (
	"testing"
	"github.com/anydef/gol-tdd/engine"
)

const fail = "\u2717"

func TestGrid(t *testing.T) {
	var g engine.Grid = engine.NewGrid(1)

	if g.Size != 1 {
		t.Fatalf("%s\t Grid should be of size 1", fail)
	}
}
