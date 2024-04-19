package game

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