package engine_test

import (
	//"testing"
	"github.com/anydef/gol-tdd/engine"
	"testing"
)

const succeed = "\u2713"
const fail = "\u2717"

func TestCell_HasStatus(t *testing.T) {
	var cell engine.Cell

	cell = engine.NewCell(engine.Alive)

	if cell.State != engine.Alive {
		t.Fatalf("%s\t Cell is not Alive", fail)
	}

	cell = engine.NewCell(engine.Dead)

	if cell.State != engine.Dead {
		t.Fatalf("%s\t Cell is not Alive", fail)
	}

}
