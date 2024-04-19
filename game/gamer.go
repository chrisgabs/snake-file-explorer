package game

import "fmt"

type Game struct {
	board    Board
	snake    Snake
	gameOver bool
}

func (g *Game) InitializeGame() {
	g.gameOver = false
	g.board.InitializeBoard(20, 20)
	g.snake.InitializeSnake(g.board.cells[1][1], g.board.cells[2][1])
}

func (g *Game) StartGame() {

	// fmt.Println(g.snake.body)
	fmt.Println(g.board)
	g.snake.Move(g.GetNextCell())
	g.snake.Direction = Rightward
	g.snake.Move(g.GetNextCell())
	fmt.Println(g.board)
	// fmt.Println(g.snake.body)

	// gameTicker := time.NewTicker(time.Millisecond * 500)
	// done := make(chan bool)

	// go func() {
	// 	for {
	// 		select {
	// 		case <-done:
	// 			fmt.Println("done")
	// 			return
	// 		case <-gameTicker.C:
	// 			fmt.Println("tick")
	// 		}
	// 	}
	// }()

	// time.Sleep(time.Second * 10)
	// gameTicker.Stop()
	// done <- true
}

func (g *Game) GetNextCell() *Cell {
	xNext, yNext := g.snake.GetNextCoordinates()
	fmt.Println(xNext)
	fmt.Println(yNext)
	return g.board.cells[xNext][yNext]
}