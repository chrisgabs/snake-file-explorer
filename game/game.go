package game

import (
	"fmt"
	"time"

	"github.com/eiannone/keyboard"
)

type Game struct {
	board    Board
	snake    Snake
	gameOver bool
}

func (g *Game) InitializeGame(width, height int) {
	g.gameOver = false
	g.board.InitializeBoard(width, height)
	g.snake.InitializeSnake(g.board.cells[1][1], g.board.cells[2][1])
}

func (g *Game) StartGame() {

	go g.listenToKeyboard()

	gameTicker := time.NewTicker(time.Millisecond * 500)
	done := make(chan bool)

	// TODO: do not allow changing direction opposite to current direction
	g.snake.Direction = Rightward
	g.board.GenerateFoodDeterministic(1,4)
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("done")
				return
			case <-gameTicker.C:
				fmt.Println(g.board)
				g.UpdateGame()
			}
		}
	}()

	time.Sleep(time.Second * 10000)
	gameTicker.Stop()
	done <- true
}

func (g *Game) UpdateGame() {
	nextCell, withinBounds := g.GetNextCell()
	if !withinBounds {
		fmt.Println("GAME OVER")
	}
	switch nextCell.CellType {
	case Empty:
		g.snake.Move(nextCell)
	case Food:
		g.snake.Grow(nextCell)
		g.board.GenerateFood()
	case SnakePart:
		fmt.Println("GAME OVER")
	}
}

// return true if valid next coordinate, false if outside of bounds
func (g *Game) GetNextCell() (*Cell, bool) {
	xNext, yNext := g.snake.GetNextCoordinates()
	if xNext >= g.board.height || yNext >= g.board.width {
		return nil, false
	}
	return g.board.cells[xNext][yNext], true
}

func (g *Game) listenToKeyboard() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println("Press ESC to quit")
	for {
		char , key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		directionPressed := NoDirection
		switch char {
		case 'a':
			directionPressed = Leftward
		case 'w':
			directionPressed = Upward
		case 'd':
			directionPressed = Rightward
		case 's':
			directionPressed = Downward
		}

		if (OppositeDirectionsMapping[directionPressed] != g.snake.Direction) && directionPressed != NoDirection {
			g.snake.Direction = directionPressed
		}

        if key == keyboard.KeyEsc {
			break
		}
	}
}