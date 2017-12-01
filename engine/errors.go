package engine

import "fmt"

type OutOfBoundsError struct {
	Index int
}

func (o *OutOfBoundsError) Error() string {
	return fmt.Sprintf("Index %d out of visible range", o.Index)
}
