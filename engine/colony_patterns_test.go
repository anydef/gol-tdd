package engine_test

import (
	"testing"
	"github.com/anydef/gol-tdd/engine"
)

type colony []struct{ c engine.Coordinate }

func TestGame_Next_BlockColony_Stays_Alive(t *testing.T) {
	block_colony := colony{
		{engine.NewCoordinate(0, 0)},
		{engine.NewCoordinate(0, 1)},
		{engine.NewCoordinate(1, 0)},
		{engine.NewCoordinate(1, 1)},
	}

	var game engine.Game = engine.NewGame(2)
	for _, tt := range block_colony {
		game.Seed(tt.c, engine.Alive)
	}

	game.Next()

	for _, tt := range block_colony {
		if game.GetCell(tt.c) != engine.Alive {
			t.Fatalf("\tCell in block colony should stay alive, coordinate %+v", tt.c)
		}
	}
}

func TestGame_Blinker_Period_of_2(t *testing.T) {
	blinker_colony_period_1 := colony{
		{engine.NewCoordinate(1, 0)},
		{engine.NewCoordinate(1, 1)},
		{engine.NewCoordinate(1, 2)},
	}

	var game engine.Game = engine.NewGame(3)

	seedColony(blinker_colony_period_1, &game)

	blinker_colony_period_2 := colony{
		{engine.NewCoordinate(0, 1)},
		{engine.NewCoordinate(1, 1)},
		{engine.NewCoordinate(2, 1)},
	}

	var grid engine.Grid
	grid = game.Next()

	assertColonyPeriod(blinker_colony_period_2, grid, func(fail bool) {
		if fail {
			t.Fatalf("\t Colony cells dont stick to expected periods")
		}
	})

	grid = game.Next()

	assertColonyPeriod(blinker_colony_period_1, grid, func(fail bool) {
		if fail {
			t.Fatalf("\t Colony cells dont stick to expected periods")
		}
	})

}
func assertColonyPeriod(c colony, grid engine.Grid, f func(b bool)) {
	for _, col := range c {
		f(grid.CellAt(col.c) != engine.Alive)
	}
}
func seedColony(c colony, g *engine.Game) {
	for _, tt := range c {
		g.Seed(tt.c, engine.Alive)
	}
}
