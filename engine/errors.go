package engine

import "fmt"

type OutOfBoundsError struct {
	X_Axis int
	Y_Axis int
}

func (o *OutOfBoundsError) Error() string {
	return fmt.Sprintf("X_Axis %d out of visible range", o.X_Axis)
}
