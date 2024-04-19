package main

import (
	"fmt"
	"snek/util"
)

func main() {
	list := &util.LinkedList[int]{}
	list.InsertFront(1)
	list.InsertFront(2)
	list.InsertFront(3)
	list.InsertBack(0)
	list.InsertBack(-1)
	fmt.Println(list)
	// list = &LinkedList{}
	// list.
	// initialize os context
	// go game.StartGame()

	// canvas := handler.CanvasHandler{}
	// canvas.InitializeOsContext()
	// fmt.Printf("canvas path: %s\n", canvas.CanvasDir)

	// canvas.CreateBlankCanvas()

	// time.Sleep(time.Second * 5)

	// canvas.DeleteAllFiles()
}