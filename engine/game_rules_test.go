package engine_test

import (
	"testing"
	"github.com/anydef/gol-tdd/engine"
)

func TestRule_Cell_with_0_neighbours_dies(t *testing.T) {
	var game engine.Game = engine.NewGame(10)

	coordinate := engine.NewCoordinate(1, 1)
	game.Seed(coordinate, engine.Alive)
	game.Next()

	grid := game.Snapshot()

	if grid.CellAt(coordinate) != engine.Dead {
		t.Fatalf("Single cell should die")
	}
}

func TestRule_Cell_with_1_neighbour_dies(t *testing.T) {
	var game engine.Game = engine.NewGame(10)

	game.Seed(engine.NewCoordinate(1, 1), engine.Alive)
	game.Seed(engine.NewCoordinate(1, 2), engine.Alive)

	game.Next()

	grid := game.Snapshot()

	if grid.CellAt(engine.NewCoordinate(1, 1)) != engine.Dead {
		t.Fatalf("Single cell should die")
	}

	if grid.CellAt(engine.NewCoordinate(1, 2)) != engine.Dead {
		t.Fatalf("Single cell should die")
	}

}
