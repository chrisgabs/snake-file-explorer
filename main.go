package main

import (
	"fmt"
	"os"
	"time"
)

type OsContext struct {
	canvasDir string
}

func (osCtx *OsContext) InitializeOsContext() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	filePath := fmt.Sprintf("%s/canvas/", dir)
	osCtx.canvasDir = filePath
}

type CanvasHandler struct {
	osCtx OsContext
}

func main() {
	// initialize os context
	var osCtx OsContext
	osCtx.InitializeOsContext()

	ch := CanvasHandler{
		osCtx: osCtx,
	}
	ch.createBlankCanvas()

	time.Sleep(time.Second * 5)

	ch.deleteAllFiles()
}



func (handler CanvasHandler) createBlankCanvas() {
	totalNumFiles := 20*8
	dir := handler.osCtx.canvasDir
	// 20 x 8 canvas
	for i := 0; i < totalNumFiles; i++ {
		fileName := fmt.Sprintf("%s/%d.txt", dir, i)
		fmt.Println(fileName)
		os.WriteFile(fileName, []byte{}, 0644)
	}
	fmt.Printf("Successfully created %d files", totalNumFiles)
}

func (handler CanvasHandler) deleteAllFiles() {
	entries, err := os.ReadDir(handler.osCtx.canvasDir)
	if err != nil {
		fmt.Println(err)
	}
	for _, e := range entries {
		err = os.Remove(handler.osCtx.canvasDir + e.Name())
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("Successfully deleted all files")
}