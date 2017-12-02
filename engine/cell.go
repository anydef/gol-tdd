package engine

import (
	"errors"
	"fmt"
)

type State int
type Direction int

const (
	_           = iota //ignored
	Dead  State = iota
	Alive State = iota
)

const (
	_                   = iota //ignored
	North     Direction = iota
	NorthEast Direction = iota
	East      Direction = iota
	SouthEast Direction = iota
	South     Direction = iota
	SouthWest Direction = iota
	West      Direction = iota
	NorthWest Direction = iota
)

var directions []Direction = []Direction{North, NorthEast, East, SouthEast, South, SouthWest, West, NorthWest}

var biDirection = map[Direction]Direction{
	North:     South,
	NorthEast: SouthWest,
	East:      West,
	SouthEast: NorthWest,
	South:     North,
	SouthWest: NorthEast,
	West:      East,
	NorthWest: SouthEast,
}

type Cell struct {
	State      State
	neighbours map[Direction]*Cell
	directions map[*Cell]Direction
}

func (c *Cell) NeighbourAt(d Direction) *Cell {
	return c.neighbours[d]
}

func (c *Cell) SetNeighbourAt(neighbour *Cell, d Direction) error {

	if c.neighbours[d] == neighbour && c.directions[neighbour] == d {
		return nil
	}

	if c.isNeighbour(neighbour, d) {
		return errors.New(fmt.Sprintf("%+v is already a neighbour at %d", neighbour, d))
	}

	c.neighbours[d] = neighbour
	c.directions[neighbour] = d

	neighbour.SetNeighbourAt(c, biDirection[d])

	return nil
}

func (c *Cell) Age() {
	c.State = c.Next()
}
func (c *Cell) isNeighbour(neighbour *Cell, d Direction) bool {
	return c.directions[neighbour] != 0 || c.neighbours[d] != nil
}

func (c *Cell) Next() State {
	var alive_neighbours int
	for _, d := range directions {
		if n := c.NeighbourAt(d); n != nil && n.State == Alive {
			alive_neighbours++
		}
	}

	if c.State == Alive {
		if alive_neighbours == 2 || alive_neighbours == 3 {
			return Alive
		}
	}
	if c.State == Dead && alive_neighbours == 3 {
		return Alive
	}

	return Dead
}

func NewCell(s State) *Cell {
	return &Cell{State: s, neighbours: make(map[Direction]*Cell), directions: make(map[*Cell]Direction)}
}
