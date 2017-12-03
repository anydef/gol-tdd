package main

import (
	"github.com/anydef/gol-tdd/engine"
)

func main() {
	game := engine.NewGame(10000)
	for i := 0; i < 1; i++ {
		game.Next()
	}
}
