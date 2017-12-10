package engine

type Game struct {
	Grid Grid
}

func (g *Game) Next() Grid {
	side := g.Grid.Size
	next_grid := NewGrid(side)

	for x := 0; x < side; x++ {
		for y := 0; y < side; y++ {
			c := NewCoordinate(x, y)
			next_grid.setCell(c, g.Grid.NextGeneration(c))
		}
	}

	g.Grid = next_grid

	return g.Snapshot()
}

func (g *Game) GetCell(c Coordinate) State {
	return g.Grid.CellAt(c)
}

func (g *Game) Seed(c Coordinate, i State) State {
	//NewCoordinate()
	return g.Grid.setCell(c, i)
}

func (g *Game) Snapshot() Grid {
	return g.Grid
}


func NewGame(side int) Game {
	return Game{Grid: NewGrid(side)}
}
