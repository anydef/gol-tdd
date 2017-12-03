package engine

type Game struct {
	grid Grid
}

func (g *Game) Next() Grid {
	next_grid := NewGrid(g.grid.Size)

	next_grid.setCell(NewCoordinate(0, 0), Dead)

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
