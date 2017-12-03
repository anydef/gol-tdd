package engine_test

import (
	"testing"
	"github.com/anydef/gol-tdd/engine"
)

func TestGame_initialized(t *testing.T) {
	var game engine.Game = engine.NewGame(10)
	var grid engine.Grid = game.Snapshot()

	if grid.Size != 10 {
		t.Fatalf("Grid size is not initialized")
	}
}

func TestGame_seed_cell(t *testing.T) {
	var game engine.Game = engine.NewGame(1)
	var grid engine.Grid = game.Snapshot()
	if grid.Size != 1 {
		t.Fatalf("Grid size is not initialized")
	}

	var c engine.Coordinate = engine.NewCoordinate(0, 0)

	game.Seed(c, engine.Alive)

	if game.GetCell(c) != engine.Alive {
		t.Fatalf("Cell at %v should be Alive", c)
	}
}

func TestGame_default_dead_cell(t *testing.T) {
	var game engine.Game = engine.NewGame(1)
	var grid engine.Grid = game.Snapshot()
	if grid.Size != 1 {
		t.Fatalf("Grid size is not initialized")
	}

	var c engine.Coordinate = engine.NewCoordinate(0, 0)

	if game.GetCell(c) != engine.Dead {
		t.Fatalf("Cell at %v should be Dead", c)
	}
}

func TestGame_cell_outside_grid_is_dead(t *testing.T) {
	var game engine.Game = engine.NewGame(1)

	var c engine.Coordinate = engine.NewCoordinate(1, 1)

	if game.GetCell(c) != engine.Dead {
		t.Fatalf("Cell at %v should be Dead", c)
	}

}

func TestRules_single_cell_dies(t *testing.T) {
	var game engine.Game = engine.NewGame(1)

	var c engine.Coordinate = engine.NewCoordinate(0, 0)
	game.Seed(c, engine.Alive)

	if game.GetCell(c) != engine.Alive {
		t.Fatalf("Cell at %v should be Alive", c)
	}

	game.Next()

	if game.GetCell(c) != engine.Dead {
		t.Fatalf("Cell at %v should be Dead", c)
	}
}
