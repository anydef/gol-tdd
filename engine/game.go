package engine

type Game struct {
	grid Grid
}

func (g *Game) Next() Grid {
	side := g.grid.Size
	next_grid := NewGrid(side)

	for x := 0; x < side; x++ {
		for y := 0; y < side; y++ {
			c := NewCoordinate(x, y)
			next_grid.setCell(c, g.grid.NextGeneration(c))
		}
	}

	g.grid = next_grid

	return g.Snapshot()
}

func (g *Game) GetCell(c Coordinate) State {
	return g.grid.CellAt(c)
}

func (g *Game) Seed(c Coordinate, i State) State {
	//NewCoordinate()
	return g.grid.setCell(c, i)
}

func (g *Game) Snapshot() Grid {
	return g.grid
}


func NewGame(side int) Game {
	return Game{grid: NewGrid(side)}
}
