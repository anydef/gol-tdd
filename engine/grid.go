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

type Direction struct {
	x int
	y int
}

var directions []Direction = []Direction{
	{0, 1},
	{0, -1},
	{1, 0},
	{1, 1},
	{1, -1},
	{-1, 0},
	{-1, 1},
	{-1, -1},
}

func (g *Grid) CellAt(c Coordinate) State {
	if !g.allowedCoordinate(c.x, c.y) {
		return Dead
	}
	return g.cells[c.x][c.y]
}

func (g *Grid) isAliveAt(c Coordinate) bool {
	return g.CellAt(c) == Alive
}

func (g *Grid) setCell(c Coordinate, s State) State {
	if inNonVisibleCoordinate(c, g.Size) {
		return Dead
	}
	g.cells[c.x][c.y] = s
	return s
}

func (g *Grid) NextGeneration(c Coordinate) State {
	if !g.allowedCoordinate(c.x, c.y) {
		return Dead
	}

	var neighbours int
	for _, direction := range directions {
		x := c.x + direction.x
		y := c.y + direction.y
		if g.isAliveAt(NewCoordinate(x, y)) {
			neighbours++
		}
	}
	if neighbours == 2 {
		return Alive
	}

	if neighbours < 2 {
		return Dead
	}

	return Dead
}

func (g *Grid) allowedCoordinate(x int, y int) bool {
	return g.allowedAxis(x) && g.allowedAxis(y)
}
func (g *Grid) allowedAxis(x int) bool {
	return x >= 0 && x < g.Size
}

func NewGrid(size int) Grid {
	cells := make([][]State, size)
	for i := range cells {
		cells[i] = make([]State, size)
	}

	return Grid{Size: size, cells: cells}
}

func inNonVisibleCoordinate(c Coordinate, size int) bool {
	return c.x >= size || c.y >= size
}

func NewCoordinate(x int, y int) Coordinate {
	return Coordinate{x, y}
}
