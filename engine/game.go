package engine

type Status bool

const Dead Status = false
const Alive Status = true

type Cell struct {
	LifeStatus Status
}

type Game struct {
	Side int
	Grid [][]Cell
}

func (g *Game) FlipBitOnIndex(x int, y int) (Status, error) {

	if bitNotOnIndex(x, y, g.Side) {
		return Dead, &OutOfBoundsError{x}
	}
	cell := g.cell(x, y)
	cell.LifeStatus = !cell.LifeStatus
	return cell.LifeStatus, nil
}

func (g *Game) cell(x int, y int) *Cell {
	return &g.Grid[x][y]
}

func (g *Game) GetStatusOf(x int, y int) (Status, error) {
	if bitNotOnIndex(x, y, g.Side) {
		return Dead, &OutOfBoundsError{x}
	}
	return g.cell(x, y).LifeStatus, nil
}

func NewGame(side_len int) Game {
	grid := make([][]Cell, side_len)
	for i := range grid {
		grid[i] = make([]Cell, side_len)
	}
	return Game{Side: side_len, Grid: grid}
}

func bitNotOnIndex(x int, y int, side int) bool {
	return (x < 0 || x >= side) || (y < 0 || y >= side)
}
