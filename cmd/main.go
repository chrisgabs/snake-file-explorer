package main

import (
	handler "snek/handler"
	"time"
)


func main() {
	// initialize os context

	canvas := handler.CanvasHandler{}
	canvas.InitializeOsContext()
	canvas.CreateBlankCanvas()

	time.Sleep(time.Second * 5)

	canvas.DeleteAllFiles()
}