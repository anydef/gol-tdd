package ui_test

import (
	"testing"
	"github.com/anydef/gol-tdd/ui/mock"
	"github.com/anydef/gol-tdd/engine"
)

func TestTerminal_ScreenCanStart(t *testing.T) {
	terminal := mock.NewScreenMock()
	err := terminal.Start()
	if err != nil {
		t.Fatalf("Unable to start the screen")
	}
}

func TestTerminal_ScreenCanClose(t *testing.T) {
	terminal := mock.NewScreenMock()
	terminal.Stop()
}

func TestTerminal_CanDraw_AtCoordinate(t *testing.T) {
	terminal := mock.NewScreenMock()
	terminal.Draw(engine.Coordinate{})
}

func TestTerminal_CanDraw_SetOfCoordinates(t *testing.T) {
	terminal := mock.NewScreenMock()
	terminal.DrawGrid([]engine.Coordinate{})
}
