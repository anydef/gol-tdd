package engine

type State int

const (
	Dead  State = iota
	Alive State = iota
)

type Coordinate struct {
	X int
	Y int
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
	if !g.isCoordinateAllowed(c.X, c.Y) {
		return Dead
	}
	return g.cells[c.X][c.Y]
}

func (g *Grid) isAliveAt(x int, y int) bool {
	return g.CellAt(NewCoordinate(x, y)) == Alive
}

func (g *Grid) setCell(c Coordinate, s State) State {
	if inNonVisibleCoordinate(c, g.Size) {
		return Dead
	}
	g.cells[c.X][c.Y] = s
	return s
}

func (g *Grid) NextGeneration(c Coordinate) State {
	if !g.isCoordinateAllowed(c.X, c.Y) {
		return Dead
	}

	var neighbours int
	for _, direction := range directions {
		x := c.X + direction.x
		y := c.Y + direction.y
		if g.isAliveAt(x, y) {
			neighbours++
		}
	}
	if neighbours == 2 && g.isAliveAt(c.X, c.Y) {
		return Alive
	}

	if neighbours == 3 {
		return Alive
	}

	if neighbours < 2 {
		return Dead
	}

	return Dead
}

func (g *Grid) isCoordinateAllowed(x int, y int) bool {
	return g.isAxisAllowed(x) && g.isAxisAllowed(y)
}

func (g *Grid) isAxisAllowed(x int) bool {
	return x >= 0 && x < g.Size
}

func (g *Grid) AliveCells() []Coordinate {
	var c []Coordinate
	for x, xAxis := range g.cells {
		for y, cell := range xAxis {
			if cell == Alive {
				c = append(c, Coordinate{x, y})
			}

		}
	}
	return c
}

func NewGrid(size int) Grid {
	cells := make([][]State, size)
	for i := range cells {
		cells[i] = make([]State, size)
	}

	return Grid{Size: size, cells: cells}
}

func inNonVisibleCoordinate(c Coordinate, size int) bool {
	return c.X >= size || c.Y >= size
}

func NewCoordinate(x int, y int) Coordinate {
	return Coordinate{x, y}
}
