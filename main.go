package main

import (
	"github.com/anydef/gol-tdd/engine"
	"github.com/anydef/gol-tdd/ui"
	"fmt"
	"os"
	"github.com/gdamore/tcell"
	"time"
)

func main() {
	screen, err := NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	screen.Start()
	defer screen.Stop()

	screen.SetStyle(tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorWhite))
	screen.Clear()

	width, height := screen.Size()
	game := engine.NewGame(min(width, height))

	seedBeacon(&game)

	snapshot := game.Snapshot()
	screen.RedrawGrid(snapshot.AliveCells())
	screen.Sync()

	quit := make(chan struct{})

	go func() {
		for {
			e := screen.Poll()
			switch ev := e.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape, tcell.KeyEnter:
					close(quit)
					return
				case tcell.KeyCtrlL:
					screen.Sync()
				}
			case *tcell.EventResize:
				screen.Sync()
			}
		}
	}()

ContinueLoop:
	for {
		select {
		case <-quit:
			break ContinueLoop
		case <-time.After(time.Millisecond * 500):
		}
		screen.Clear()
		next := game.Next()
		screen.RedrawGrid(next.AliveCells())
	}

}

func NewScreen() (ui.Terminal, error) {
	screen, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}
	return ui.TerminalImpl{Screen: screen}, nil
}

func min(a int, b int) int {
	if b > a {
		return a
	}
	return b
}

func seedBeacon(g *engine.Game) {
	g.Seed(engine.NewCoordinate(5, 4), engine.Alive)
	g.Seed(engine.NewCoordinate(5, 5), engine.Alive)
	g.Seed(engine.NewCoordinate(5, 6), engine.Alive)
}
