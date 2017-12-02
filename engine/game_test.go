package engine_test

import (
	"testing"
	"github.com/anydef/gol-tdd/engine"
)

const succeed = "\u2713"
const fail = "\u2717"

func testAlive(g engine.Game, x int, y int, t *testing.T) {
	statusTest(g, x, y, engine.Alive, t)
}

func testDead(g engine.Game, x int, y int, t *testing.T) {
	statusTest(g, x, y, engine.Dead, t)
}

func statusTest(g engine.Game, x int, y int, s engine.Status, t *testing.T) {
	if status, err := g.TestStatusOf(x, y); status != s || err != nil {
		t.Fatalf("%s\t Wrong status %v of a seeded cell", fail, status)
	}
}

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

	status, _ := game.TestStatusOf(0, 0)

	if status != engine.Dead {
		t.Fatalf("%s\t Bit status is not %v", fail, engine.Dead)
	}
}

func TestBitStatus_OfNotExistingBit(t *testing.T) {
	side_len := 1
	game := engine.NewGame(side_len)

	if _, err := game.TestStatusOf(1, 0); err == nil {
		t.Fatalf("%s\tOut of bound error should be returned", fail)
	}

	if _, err := game.TestStatusOf(-1, 0); err == nil {
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

	if status, err := game.TestStatusOf(0, 0); status != engine.Dead || err != nil {
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
			if status, err := game.TestStatusOf(x, y); status != engine.Dead || err != nil {
				t.Fatalf("%s\t Initial bit is %v", fail, engine.Dead)
			}

			if status, err := game.FlipBitOnIndex(x, y); status != engine.Alive || err != nil {
				t.Fatalf("%s\t Initial bit is %v", fail, engine.Alive)
			}

			if status, err := game.TestStatusOf(x, y); status != engine.Alive || err != nil {
				t.Fatalf("%s\t Status should remain alive %v", fail, engine.Alive)
			}

		}
	}
}

func TestGame_1x1_SeedAliveOutOfBounds(t *testing.T) {
	side_len := 1
	game := engine.NewGame(side_len)

	if err := game.SeedAlive(1, 1); err == nil {
		t.Fatalf("%s\t Should not allow setting cell outside of a grid", fail)
	}
}

func TestGame_1x1_SeedAlive(t *testing.T) {
	side_len := 1
	game := engine.NewGame(side_len)

	if err := game.SeedAlive(0, 0); err != nil {
		t.Fatalf("%s\t Unable to seed a cell", fail)
	}

	testAlive(game, 0, 0, t)
}

func TestGame_1x1_OneCellDiesWhenHasZeroNeighbours(t *testing.T) {
	game := engine.NewGame(1)

	var x, y int

	if err := game.SeedAlive(x, y); err != nil {
		t.Fatalf("%s\t Unable to seed a cell", fail)
	}

	testAlive(game, x, y, t)

	game.Next()

	if status, err := game.TestStatusOf(x, y); status != engine.Dead || err != nil {
		t.Fatalf("%s\t Cell with zero neighbours should die", fail)
	}
}

func TestGame_NxN_DeadDoesntBecomeAliveOnNext(t *testing.T) {
	var x, y int
	for i := 1; i < 5; i++ {
		game := engine.NewGame(i)

		testDead(game, x, y, t)

		game.Next()

		testDead(game, x, y, t)
	}
}

func TestGame_2x2_CellWith_LessThan_TwoNeighbours_Dies(t *testing.T) {
	game := engine.NewGame(2)

	alive_cells := []struct {
		x int
		y int
	}{
		{0, 0},
		{0, 1},
	}

	for _, tt := range alive_cells {
		game.SeedAlive(tt.x, tt.y)
	}

	for _, tt := range alive_cells {
		testAlive(game, tt.x, tt.y, t)
	}

	game.Next()

	for _, tt := range alive_cells {
		testDead(game, tt.x, tt.y, t)
	}

}

//func TestGame_2x2_CellWith_ThreeNeighbours_StaysAlive(t *testing.T) {
//	game := engine.NewGame(2)
//
//	alive_cells := []struct {
//		x int
//		y int
//	}{
//		{0, 0},
//		{0, 1},
//		{1, 0},
//		{1, 1},
//	}
//
//	for _, tt := range alive_cells {
//		game.SeedAlive(tt.x, tt.y)
//	}
//
//	for _, tt := range alive_cells {
//		testAlive(game, tt.x, tt.y, t)
//	}
//
//	game.Next()
//
//	for _, tt := range alive_cells {
//		testAlive(game, tt.x, tt.y, t)
//	}
//
//}
