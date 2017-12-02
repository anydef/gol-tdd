package engine

type Status bool

const (
	Dead  Status = false
	Alive Status = true
)

type Cell struct {
	LifeStatus Status
	NextStatus Status
}

func (c *Cell) Next() {

}

type Grid [][]Cell

type Game struct {
	Side int
	Grid Grid
}

func (g *Game) Next() {
	for x := range g.Grid {
		for y := range g.Grid[x] {
			if status, err := g.TestStatusOf(x, y); status == Alive && err == nil {
				g.SeedDead(x, y)
			}
		}
	}
}

func (g *Game) SeedDead(x int, y int) error {
	return g.seed(x, y, Dead)
}

func (g *Game) SeedAlive(x int, y int) error {
	return g.seed(x, y, Alive)
}

func (g *Game) seed(x int, y int, s Status) error {
	if bitOutsideGrid(x, y, g.Side) {
		return &OutOfBoundsError{x, y}
	}
	g.cell(x, y).LifeStatus = s
	return nil
}

func (g *Game) FlipBitOnIndex(x int, y int) (Status, error) {
	if bitOutsideGrid(x, y, g.Side) {
		return Dead, &OutOfBoundsError{x, y}
	}
	cell := g.cell(x, y)
	cell.LifeStatus = !cell.LifeStatus
	return cell.LifeStatus, nil
}

func (g *Game) cell(x int, y int) *Cell {
	return &g.Grid[x][y]
}

func (g *Game) TestStatusOf(x int, y int) (Status, error) {
	if bitOutsideGrid(x, y, g.Side) {
		return Dead, &OutOfBoundsError{x, y}
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

func bitOutsideGrid(x int, y int, side int) bool {
	return (x < 0 || x >= side) || (y < 0 || y >= side)
}
