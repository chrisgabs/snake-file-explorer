package canvas_handler

import (
	"fmt"
	"os"
)

type CanvasHandler struct {
	canvasDir string
}

func (handler *CanvasHandler) InitializeOsContext() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	filePath := fmt.Sprintf("%s/canvas/", dir)
	handler.canvasDir = filePath
}

func (handler CanvasHandler) CreateBlankCanvas() {
	// 20 x 8 canvas
	totalNumFiles := 20 * 8
	dir := handler.canvasDir
	for i := 0; i < totalNumFiles; i++ {
		fileName := fmt.Sprintf("%s%d.txt", dir, i)
		fmt.Println(fileName)
		os.WriteFile(fileName, []byte{}, 0644)
	}
	fmt.Printf("Successfully created %d files", totalNumFiles)
}

func (handler CanvasHandler) DeleteAllFiles() {
	entries, err := os.ReadDir(handler.canvasDir)
	if err != nil {
		fmt.Println(err)
	}
	for _, e := range entries {
		err = os.Remove(handler.canvasDir + e.Name())
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("Successfully deleted all files")
}