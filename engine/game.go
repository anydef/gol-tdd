package engine

type Game struct {
	grid Grid
}

func (g *Game) GetCell(c Coordinate) State {
	return g.grid.IsAliveAt(c)
}

func (g *Game) Seed(c Coordinate, i State) State {
	return g.grid.SetCell(c, i)
}

func (g *Game) Snapshot() Grid {
	return g.grid
}

func NewGame(side int) Game {
	cells := make([][]State, side)
	for i := range cells {
		cells[i] = make([]State, side)
	}
	return Game{grid: Grid{Size: side, cells: cells}}
}
