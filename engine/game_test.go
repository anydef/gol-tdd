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

func TestInitGame_WithOneDimension(t *testing.T) {
	side_len := 1
	game := engine.NewGame(side_len)

	if game.Side != side_len {
		t.Fatalf("%s\tGame side should be %d ", fail, side_len)
	}

	t.Logf("%s\tGame init is OK", succeed)
}

func TestBitStatus_Default(t *testing.T) {
	side_len := 1
	game := engine.NewGame(side_len)

	status, _ := game.GetStatusOf(0, 0)

	if status != engine.Dead {
		t.Fatalf("%s\t Bit status is not %v", fail, engine.Dead)
	}
}

func TestBitStatus_OfNotExistingBit(t *testing.T) {
	side_len := 1
	game := engine.NewGame(side_len)

	if _, err := game.GetStatusOf(1, 0); err == nil {
		t.Fatalf("%s\tOut of bound error should be returned", fail)
	}

	if _, err := game.GetStatusOf(-1, 0); err == nil {
		t.Fatalf("%s\t Should not access item outside", fail)
	}
}

func TestFlipBit_OnNotExistingIndex(t *testing.T) {
	side_len := 1
	game := engine.NewGame(side_len)

	if status, err := game.FlipBitOnIndex(1, 0); status != engine.Dead && err != nil {
		t.Fatalf("%s\t Should not allow flipping bit outside grid", fail)
	}

	if status, err := game.FlipBitOnIndex(-1, 0); status != engine.Dead && err != nil {
		t.Fatalf("%s\t Should not allow flipping bit outside grid", fail)
	}
}

func TestFlipBit_On_1x1_Grid(t *testing.T) {
	side_len := 1
	game := engine.NewGame(side_len)

	if status, err := game.GetStatusOf(0, 0); status != engine.Dead || err != nil {
		t.Fatalf("%s\t Initial bit is %v", fail, engine.Dead)
	}

	if status, err := game.FlipBitOnIndex(0, 0); status != engine.Alive && err != nil {
		t.Fatalf("%s\t Flipped bit should be %v", fail, engine.Alive)
	}

	if status, err := game.FlipBitOnIndex(0, 0); status != engine.Dead && err != nil {
		t.Fatalf("%s\t Flipped bit should be %v", fail, engine.Dead)
	}
}

func TestBitsFlip_On_NxN_Grid(t *testing.T) {
	side_len := 2
	game := engine.NewGame(side_len)

	for x := 0; x < side_len; x++ {
		for y := 0; y < side_len; y++ {
			if status, err := game.GetStatusOf(x, y); status != engine.Dead || err != nil {
				t.Fatalf("%s\t Initial bit is %v", fail, engine.Dead)
			}

			if status, err := game.FlipBitOnIndex(x, y); status != engine.Alive || err != nil {
				t.Fatalf("%s\t Initial bit is %v", fail, engine.Alive)
			}

			if status, err := game.GetStatusOf(x, y); status != engine.Alive || err != nil {
				t.Fatalf("%s\t Status should remain alive %v", fail, engine.Alive)
			}

		}
	}

}

