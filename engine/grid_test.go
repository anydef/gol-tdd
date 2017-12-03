package engine

import (
	"testing"
)

const fail = "\u2717"

func TestGrid(t *testing.T) {
	var g Grid = NewGrid(1)

	if g.Size != 1 {
		t.Fatalf("%s\t Grid should be of size 1", fail)
	}
}

func TestGrid_NextGenerationCell(t *testing.T) {
	c := NewCoordinate(1, 1)
	var g Grid = NewGrid(5)

	if g.setCell(c, Alive) != Alive {
		t.Fatalf("Cell at %+v should be Alive", c)
	}

	if g.NextGeneration(c) != Dead {
		t.Fatalf("Cell at %+v should die", c)
	}
}

func TestGrid_NextGeneration_Cell_OnTheEdges(t *testing.T) {
	setup := []struct {
		x int
		y int
	}{
		{0, 0},
		{0, 1},
		{1, 0},
		{1, 1},
	}

	for _, tt := range setup {
		var g Grid = NewGrid(2)
		c := NewCoordinate(tt.x, tt.y)

		if g.setCell(c, Alive) != Alive {
			t.Fatalf("Cell at %+v should be Alive", c)
		}

		if g.NextGeneration(c) != Dead {
			t.Fatalf("Cell at %+v should die", c)
		}
	}
}

func TestGrid_Cell_With_1_Neighbour_Dies(t *testing.T) {
	cells := []struct {
		x int
		y int
	}{
		{0, 1},
		{0, 0},
	}

	var g Grid = NewGrid(2)
	for _, cell := range cells {
		g.setCell(NewCoordinate(cell.x, cell.y), Alive)
	}
	for _, cell := range cells {
		c := NewCoordinate(cell.x, cell.y)
		if g.NextGeneration(c) != Dead {
			t.Fatalf("Cell with one neighbour at %+v should die", c)
		}
	}
}


func TestGrid_Cell_With_2_Neighbours_Lives(t *testing.T) {
	cells := []struct {
		x int
		y int
	}{
		{0, 1},
		{0, 0},
		{1, 0},
	}

	var g Grid = NewGrid(2)
	for _, cell := range cells {
		g.setCell(NewCoordinate(cell.x, cell.y), Alive)
	}
	for _, cell := range cells {
		c := NewCoordinate(cell.x, cell.y)
		if g.NextGeneration(c) != Alive {
			t.Fatalf("Cell with two neighbours at %+v should live", c)
		}
	}
}
