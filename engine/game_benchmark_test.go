package engine_test

import (
	"testing"
	"github.com/anydef/gol-tdd/engine"
)

func BenchmarkBlinker_3x3(b *testing.B) {
	colony := colony{
		{engine.NewCoordinate(1, 0)},
		{engine.NewCoordinate(1, 1)},
		{engine.NewCoordinate(1, 2)},
	}
	var game engine.Game = engine.NewGame(3)
	seedColony(colony, &game)
	for i := 0; i < b.N; i ++ {
		game.Next()
	}
}

func BenchmarkSingle_1x1(b *testing.B) {
	colony := colony{
		{engine.NewCoordinate(0, 0)},
	}
	var game engine.Game = engine.NewGame(1)
	seedColony(colony, &game)
	for i := 0; i < b.N; i ++ {
		game.Next()
	}
}

func BenchmarkSingle_2x2(b *testing.B) {
	colony := colony{
		{engine.NewCoordinate(0, 0)},
	}
	var game engine.Game = engine.NewGame(2)
	seedColony(colony, &game)
	for i := 0; i < b.N; i ++ {
		game.Next()
	}
}

func Benchmark16Cells(b *testing.B) {
	colony := colony{
		{engine.NewCoordinate(0, 0)},
	}
	var game engine.Game = engine.NewGame(4)
	seedColony(colony, &game)
	for i := 0; i < b.N; i ++ {
		game.Next()
	}
}
