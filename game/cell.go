package game

import "fmt"

const (
	Empty int = iota
	SnakePart
	Food
)

type Cell struct {
	CellType int
	X        int
	Y        int
}

func (c Cell) String() string {
	return fmt.Sprintf("%v", c.CellType)
}