package engine_test

import (
	"github.com/anydef/gol-tdd/engine"
	"testing"
)

//const succeed = "\u2713"
const fail = "\u2717"

func TestCell_HasStatus(t *testing.T) {
	var cell *engine.Cell

	cell = engine.NewCell(engine.Alive)

	if cell.State != engine.Alive {
		t.Fatalf("%s\t Cell is not Alive", fail)
	}

	cell = engine.NewCell(engine.Dead)

	if cell.State != engine.Dead {
		t.Fatalf("%s\t Cell is not Alive", fail)
	}

}

func TestCell_Dead_RemainsDead_OnNext(t *testing.T) {
	var cell *engine.Cell

	if cell = engine.NewCell(engine.Dead); cell.State != engine.Dead {
		t.Fatalf("%s\t Cell should be Dead", fail)
	}

	if state := cell.NextGeneration(); state != engine.Dead {
		t.Fatalf("%s\t Cell should be Dead", fail)
	}
}

func TestCell_Alive_WillDie_OnNext(t *testing.T) {
	cell := engine.NewCell(engine.Alive)

	if state := cell.NextGeneration(); state != engine.Dead {
		t.Fatalf("%s\t Cell should be Dead", fail)
	}
}

func TestCell_Ages_And_ChangesState(t *testing.T) {
	cell := engine.NewCell(engine.Alive)

	if state := cell.NextGeneration(); state != engine.Dead {
		t.Fatalf("%s\t Cell should be Dead", fail)
	}

	cell.Age()

	if cell.State != engine.Dead {
		t.Fatalf("%s\t Cell should be Dead", fail)
	}
}

func TestCell_Has_Neighbours_In_All_Directions(t *testing.T) {
	setup := []struct {
		neighbour *engine.Cell
		direction engine.Direction
	}{
		{engine.NewCell(engine.Alive), engine.North},
		{engine.NewCell(engine.Alive), engine.NorthEast},
		{engine.NewCell(engine.Alive), engine.East},
		{engine.NewCell(engine.Alive), engine.SouthEast},
		{engine.NewCell(engine.Alive), engine.South},
		{engine.NewCell(engine.Alive), engine.SouthWest},
		{engine.NewCell(engine.Alive), engine.West},
		{engine.NewCell(engine.Alive), engine.NorthWest},
	}

	cell := engine.NewCell(engine.Alive)

	for _, tt := range setup {
		cell.SetNeighbourAt(tt.neighbour, tt.direction)

		if cell.NeighbourAt(tt.direction) != tt.neighbour {
			t.Fatalf("%s\t Cell should have neighbour in %v", fail, tt.direction)
		}
	}
}

func TestCell_Biderectional_NeighbourRelation(t *testing.T) {
	setup := []struct {
		direction   engine.Direction
		bidirection engine.Direction
	}{
		{engine.North, engine.South},
		{engine.NorthEast, engine.SouthWest},
		{engine.East, engine.West},
		{engine.SouthEast, engine.NorthWest},
		{engine.South, engine.North},
		{engine.SouthWest, engine.NorthEast},
		{engine.West, engine.East},
		{engine.NorthWest, engine.SouthEast},
	}

	for _, tt := range setup {
		neighbour := engine.NewCell(engine.Alive)
		cell := engine.NewCell(engine.Alive)

		cell.SetNeighbourAt(neighbour, tt.direction)

		if c := cell.NeighbourAt(tt.direction); c != neighbour {
			t.Fatalf("%s\t Cell %v should have neighbour %v at %v. Got: %v", fail, cell, neighbour, tt.direction, c)
		}

		if c := neighbour.NeighbourAt(tt.bidirection); c != cell {
			t.Fatalf("%s\t Cell %+v should have neighbour %+v at %v. Got: %v", fail, neighbour, cell, tt.bidirection, c)
		}
	}
}

func TestCell_CannotSet_Same_Neighbour_AtMultipleDirections(t *testing.T) {
	setup := []struct {
		direction   engine.Direction
		bidirection engine.Direction
	}{
		{engine.North, engine.South},
		{engine.NorthEast, engine.SouthWest},
		{engine.East, engine.West},
		{engine.SouthEast, engine.NorthWest},
		{engine.South, engine.North},
		{engine.SouthWest, engine.NorthEast},
		{engine.West, engine.East},
		{engine.NorthWest, engine.SouthEast},
	}

	for _, tt := range setup {
		cell := engine.NewCell(engine.Alive)
		neighbour := engine.NewCell(engine.Alive)

		cell.SetNeighbourAt(neighbour, tt.direction)

		if err := cell.SetNeighbourAt(neighbour, tt.bidirection); err == nil {
			t.Fatalf("%s\t Same neighbour in different directions should not be allowed. %v", fail, tt.direction)
		}
	}
}

type neighbour struct {
	n *engine.Cell
	d engine.Direction
}
type neighbours []neighbour

func TestRule_When_LessThan_2_Alive_Neighbours_Die_FromUnderpopulation(t *testing.T) {
	setup := []struct {
		neighbours neighbours
	}{
		{neighbours: []neighbour{},},
		{neighbours: []neighbour{{n: engine.NewCell(engine.Alive)}},},
		{neighbours: []neighbour{
			{n: engine.NewCell(engine.Alive), d: engine.North},
			{n: engine.NewCell(engine.Dead), d: engine.South},
		}},
	}

	for _, tt := range setup {

		cell := engine.NewCell(engine.Alive)

		for _, neighbour := range tt.neighbours {
			if err := cell.SetNeighbourAt(neighbour.n, neighbour.d); err != nil {
				t.Fatalf("%s\t Could not set %+v as neighbour for %+v", fail, neighbour, cell)
			}
		}

		if next := cell.NextGeneration(); next != engine.Dead {
			t.Fatalf("%s\t Cell %v should die from underpopulation", fail, cell)
		}
	}
}

func TestRule_When_2_Or_3_Alive_Neighbours_Then_Live(t *testing.T) {
	setup := []struct {
		neighbours neighbours
	}{
		{neighbours: []neighbour{
			{n: engine.NewCell(engine.Alive), d: engine.North},
			{n: engine.NewCell(engine.Alive), d: engine.NorthEast},
		}},
		{neighbours: []neighbour{
			{n: engine.NewCell(engine.Alive), d: engine.North},
			{n: engine.NewCell(engine.Alive), d: engine.NorthEast},
			{n: engine.NewCell(engine.Alive), d: engine.East},
		}},
		{neighbours: []neighbour{
			{n: engine.NewCell(engine.Alive), d: engine.North},
			{n: engine.NewCell(engine.Alive), d: engine.NorthEast},
			{n: engine.NewCell(engine.Alive), d: engine.East},
			{n: engine.NewCell(engine.Dead), d: engine.SouthEast},
		}},
	}
	for _, tt := range setup {
		cell := engine.NewCell(engine.Alive)

		for _, neighbour := range tt.neighbours {
			if err := cell.SetNeighbourAt(neighbour.n, neighbour.d); err != nil {
				t.Fatalf("%s\t Could not set %+v as neighbour for %+v", fail, neighbour, cell)
			}
		}
		if next := cell.NextGeneration(); next != engine.Alive {
			t.Fatalf("%s\t Cell %v should live", fail, cell)
		}
	}
}

func TestRule_LiveCell_WithMoraThan_3_LiveNeighbours_Dies_Of_Overpopulation(t *testing.T) {
	setup := []struct {
		neighbours neighbours
	}{
		{neighbours: []neighbour{
			{n: engine.NewCell(engine.Alive), d: engine.North},
			{n: engine.NewCell(engine.Alive), d: engine.NorthEast},
			{n: engine.NewCell(engine.Alive), d: engine.East},
			{n: engine.NewCell(engine.Alive), d: engine.SouthEast},
		}},
	}
	for _, tt := range setup {
		cell := engine.NewCell(engine.Alive)

		for _, neighbour := range tt.neighbours {
			if err := cell.SetNeighbourAt(neighbour.n, neighbour.d); err != nil {
				t.Fatalf("%s\t Could not set %+v as neighbour for %+v", fail, neighbour, cell)
			}
		}
		if next := cell.NextGeneration(); next != engine.Dead {
			t.Fatalf("%s\t Cell %v should die because of overpopulation", fail, cell)
		}
	}
}

func TestRule_Exactly_3_AliveNeighbours_BreedAliveCell(t *testing.T) {
	setup := []struct {
		neighbours neighbours
	}{
		{neighbours: []neighbour{
			{n: engine.NewCell(engine.Alive), d: engine.North},
			{n: engine.NewCell(engine.Alive), d: engine.NorthEast},
			{n: engine.NewCell(engine.Alive), d: engine.East},
		}},
	}
	for _, tt := range setup {
		cell := engine.NewCell(engine.Dead)

		for _, neighbour := range tt.neighbours {
			if err := cell.SetNeighbourAt(neighbour.n, neighbour.d); err != nil {
				t.Fatalf("%s\t Could not set %+v as neighbour for %+v", fail, neighbour, cell)
			}
		}
		if next := cell.NextGeneration(); next != engine.Alive {
			t.Fatalf("%s\t Cell %v should be born", fail, cell)
		}
	}
}

func TestRule_LessThan_Or_MoreThan_3_Cells_CannotBreed(t *testing.T) {
	setup := []struct {
		neighbours neighbours
	}{
		{neighbours: []neighbour{
			{n: engine.NewCell(engine.Alive), d: engine.North},
			{n: engine.NewCell(engine.Alive), d: engine.NorthEast},
		}},
		{neighbours: []neighbour{
			{n: engine.NewCell(engine.Alive), d: engine.North},
			{n: engine.NewCell(engine.Alive), d: engine.SouthEast},
			{n: engine.NewCell(engine.Alive), d: engine.West},
			{n: engine.NewCell(engine.Alive), d: engine.NorthEast},
		}},
	}
	for _, tt := range setup {
		cell := engine.NewCell(engine.Dead)

		for _, neighbour := range tt.neighbours {
			if err := cell.SetNeighbourAt(neighbour.n, neighbour.d); err != nil {
				t.Fatalf("%s\t Could not set %+v as neighbour for %+v", fail, neighbour, cell)
			}
		}
		if next := cell.NextGeneration(); next != engine.Dead {
			t.Fatalf("%s\t Less than, or more than 3 cells cannot breed", fail)
		}
	}
}

func TestRules_BlockColony_DoesntDie(t *testing.T) {
	cell_1_1 := engine.NewCell(engine.Alive)
	cell_1_2 := engine.NewCell(engine.Alive)
	cell_2_1 := engine.NewCell(engine.Alive)
	cell_2_2 := engine.NewCell(engine.Alive)

	cell_1_1.SetNeighbourAt(cell_1_2, engine.North)
	cell_1_1.SetNeighbourAt(cell_2_2, engine.NorthEast)
	cell_1_1.SetNeighbourAt(cell_2_1, engine.East)

	cell_1_2.SetNeighbourAt(cell_2_2, engine.East)
	cell_1_2.SetNeighbourAt(cell_2_1, engine.SouthEast)

	cell_2_1.SetNeighbourAt(cell_2_2, engine.North)

	for _, c := range []*engine.Cell{cell_1_1, cell_1_2, cell_2_1, cell_2_2} {
		if c.NextGeneration() != engine.Alive {
			t.Fatalf("%s\t Cell %v in block colony should live", fail, cell_1_1)
		}
	}
}

func TestCell_Cannot_BeOnesNeighbour(t *testing.T) {
	cell := engine.NewCell(engine.Alive)
	if cell.SetNeighbourAt(cell, engine.North) == nil {
		t.Fatalf("%s\t Cell %+v cannot be ones neighbour", fail, cell)
	}
}


func TestCell_ChangeStatus_On_AgeOnly(t *testing.T) {
	cell := engine.NewCell(engine.Alive)

	if cell.State != engine.Alive {
		t.Fatalf("%s\t Cell %+v status shouldn't change", fail, cell)
	}

	if cell.NextGeneration() != engine.Dead {
		t.Fatalf("%s\t Cell's %+v should die in next generation", fail, cell)
	}

	if cell.State != engine.Alive {
		t.Fatalf("%s\t Cell %+v should not die before aging", fail, cell)
	}

	cell.Age()

	if cell.State != engine.Dead {
		t.Fatalf("%s\t Cell %+v should be dead after aging", fail, cell)
	}


}