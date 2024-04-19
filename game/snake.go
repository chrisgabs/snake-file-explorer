package game

import "snek/util"

const (
	Upward int = iota
	Downward
	Leftward
	Rightward
)

type Snake struct {
	Direction int
	body util.LinkedList[*Cell]
}

// x, y coordinates of head
func (snake *Snake) InitializeSnake(headCell, tailCell *Cell) {
	snake.Direction = Upward
	headCell.CellType = SnakePart
	tailCell.CellType = SnakePart
	snake.body.InsertFront(headCell)
	snake.body.InsertBack(tailCell)
}

func (snake *Snake) Move(nextCell *Cell) {
	nextCell.CellType = SnakePart
	snake.body.InsertFront(nextCell)
	snake.body.Tail.Data.CellType = Empty
	_ = snake.body.RemoveTail()
}

func (snake *Snake) String() string {
	return snake.body.String()
}

func (snake *Snake) GetNextCoordinates() (x, y int) {
	if snake.body.Size < 2 {
		return -1, -1
	}
	xH := snake.body.Head.Data.X
	yH := snake.body.Head.Data.Y
	switch snake.Direction{
	case Upward:
		return xH-1, yH
	case Downward:
		return xH+1, yH
	case Leftward:
		return xH, yH-1
	case Rightward:
		return xH, yH+1
	}
	panic("no next coordinate")
	// return -1, -1
}