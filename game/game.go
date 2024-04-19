package game

import (
	"fmt"
	"time"
)

type Node struct {
	L         *Node
	R         *Node
	T         *Node
	D         *Node
	Type      string // board, snek, food
	Direction string
}

type Game struct {
	board Node
}

func (g *Game) initializeBoard() {
	for {
		// how
	}
}

func (g *Game) StartGame() {
	gameTicker := time.NewTicker(time.Millisecond * 500)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("done")
				return
			case <-gameTicker.C:
				fmt.Println("tick")
			}
		}
	}()

	time.Sleep(time.Second * 10)
	gameTicker.Stop()
	done <- true
}
