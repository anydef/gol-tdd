package engine

type State int

const (
	Alive State = iota
	Dead  State = iota
)

type Cell struct {
	State State
}

func NewCell(s State) Cell {
	return Cell{s}
}
