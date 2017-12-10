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

func TestGrid_Cell_With_3_Neighbours_Lives(t *testing.T) {
	cells := []struct {
		x int
		y int
	}{
		{0, 1},
		{0, 0},
		{1, 0},
		{1, 1},
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

func TestGrid_Cell_With_More_Than_3_Neighbours_Dies(t *testing.T) {
	cells := []struct {
		x int
		y int
	}{
		{0, 1},
		{1, 0},
		{1, 1},
		{1, 2},
		{2, 1},
	}

	var g Grid = NewGrid(3)
	for _, cell := range cells {
		g.setCell(NewCoordinate(cell.x, cell.y), Alive)
	}

	c := NewCoordinate(1, 1)
	if g.NextGeneration(c) != Dead {
		t.Fatalf("Cell with more that three neighbours at %+v should die", c)
	}

}

func TestGrid_Exactly_3_Cells_Breed(t *testing.T) {
	first_generation := []struct {
		x int
		y int
	}{
		{0, 1},
		{1, 0},
		{1, 1},
		{1, 2},
		{2, 1},
	}

	var g Grid = NewGrid(3)
	for _, cell := range first_generation {
		g.setCell(NewCoordinate(cell.x, cell.y), Alive)
	}

	second_generation := []struct {
		x int
		y int
	}{
		{0, 0},
		{0, 1},
		{0, 2},
		{1, 0},
		{1, 2},
		{2, 0},
		{2, 1},
		{2, 2},
	}

	for _, tt := range second_generation {

		c := NewCoordinate(tt.x, tt.y)
		if g.NextGeneration(c) != Alive {
			t.Fatalf("Exactly three cells have to reproduce at %+v", c)
		}
	}

}

func TestGrid_ReturnsSet_ofNone_forNilGrid(t *testing.T) {
	var g Grid = NewGrid(0)
	var c []Coordinate = g.AliveCells()
	if len(c) != 0 {
		t.Fatalf("Empty grid should return empty list of coordinates")
	}
}

func TestGrid_ReturnsSet_ofLiveCells_whenNoneAlive(t *testing.T) {
	var g Grid = NewGrid(10)
	var c []Coordinate = g.AliveCells()
	if len(c) != 0 {
		t.Fatalf("Empty grid should return empty list of coordinates")
	}
}

func TestGrid_ReturnsSet_ofLiveCells(t *testing.T) {
	cells := []struct {
		x int
		y int
	}{
		{0, 1},
		{0, 0},
		{1, 0},
	}

	var g Grid = NewGrid(10)
	for _, cell := range cells {
		g.setCell(NewCoordinate(cell.x, cell.y), Alive)
	}

	var c []Coordinate = g.AliveCells()
	if len(c) != len(cells) {
		t.Fatalf("Number of live cells=%d, should be exact as popuated cells = %d. ", len(c), len(cells))
	}
}
