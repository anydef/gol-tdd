package engine_test

import (
	"testing"
	"github.com/anydef/gol-tdd/engine"
)

const succeed = "\u2713"
const fail = "\u2717"

func TestCreateGame(t *testing.T) {
	side_len := 0
	var game engine.Game = engine.NewGame(side_len)

	if game.Side != side_len {
		t.Fatalf("%s\tGame side should be %d ", fail, side_len)
	}

	t.Logf("%s\tGame init is OK", succeed)
}

func TestInitGameWithOneDimension(t *testing.T) {
	side_len := 1
	game := engine.NewGame(side_len)

	if game.Side != side_len {
		t.Fatalf("%s\tGame side should be %d ", fail, side_len)
	}

	t.Logf("%s\tGame init is OK", succeed)
}

func TestBitStatusDefault(t *testing.T) {
	side_len := 1
	game := engine.NewGame(side_len)

	status, _ := game.GetStatusOf(0)

	if status != false {
		t.Fatalf("%s\t Bit status is not %b", fail, false)
	}
}

func TestBitStatusOfUnexistingBit(t *testing.T) {
	side_len := 1
	game := engine.NewGame(side_len)

	_, err := game.GetStatusOf(1)

	if err == nil {
		t.Fatalf("%s\tOut of bound error should be returned", fail)
	}
}

func TestStatusOnNegativeIndex(t *testing.T) {
	side_len := 1
	game := engine.NewGame(side_len)

	if _, err := game.GetStatusOf(-1); err == nil {
		t.Fatalf("%s\t Should not access item with negative index", fail)
	}

}
