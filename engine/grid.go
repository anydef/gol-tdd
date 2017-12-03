package engine

type State int

const (
	Dead  State = iota
	Alive State = iota
)

type Coordinate struct {
	x int
	y int
}

type Grid struct {
	cells [][]State
	Size  int
}

func (g *Grid) IsAliveAt(c Coordinate) State {
	if c.x >= g.Size || c.y >= g.Size {
		return Dead
	}

	return g.cells[c.x][c.y]
}

func (g *Grid) SetCell(c Coordinate, s State) State {
	if c.x >= g.Size || c.y >= g.Size {
		return Dead
	}
	g.cells[c.x][c.y] = s
	return s
}

func NewGrid(size int) Grid {
	return Grid{Size: size}
}

func NewCoordinate(x int, y int) Coordinate {
	return Coordinate{x, y}
}
