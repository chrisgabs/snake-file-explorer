package main

import (
	game "snek/game"
)

func main() {
	// list = &LinkedList{}
	// list.
	// initialize os context
	g := &game.Game{}
	
	g.InitializeGame(10, 10)
	g.StartGame()
	// go game.StartGame()

	// canvas := handler.CanvasHandler{}
	// canvas.InitializeOsContext()
	// fmt.Printf("canvas path: %s\n", canvas.CanvasDir)

	// canvas.CreateBlankCanvas()

	// time.Sleep(time.Second * 5)

	// canvas.DeleteAllFiles()
}