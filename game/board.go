package game

import (
	"math/rand"
)

type Board struct {
	width, height int
	cells         [][]*Cell
}

func (board *Board) NewBoard(width, height int) {
	board.cells = make([][]*Cell, height)
	for i := range board.cells {
		board.cells[i] = make([]*Cell, width)
	}

	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			board.cells[row][col] = &Cell{
				CellType: Empty,
				X:        row,
				Y:        col,
			}
		}
	}
}

// TODO: catch case where there are no more emptry cells (win condition)
func (board *Board) GenerateFood() (x int, y int) {
	xRandom := rand.Intn(board.height)
	yRandom := rand.Intn(board.width)

	for {
		if board.cells[xRandom][yRandom].CellType == Empty {
			board.cells[xRandom][yRandom].CellType = Food
			return xRandom, yRandom
		}
		xRandom = rand.Intn(board.height)
		yRandom = rand.Intn(board.width)
	}
}