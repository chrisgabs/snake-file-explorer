package canvas_handler

import (
	"fmt"
	"os"
	"path/filepath"
)

type CanvasHandler struct {
	CanvasDir string
}

func (handler *CanvasHandler) InitializeOsContext() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	path := filepath.Clean(filepath.Join(dir, ".."))
	path = fmt.Sprintf("%s/canvas/", path)
	handler.CanvasDir = path
}

func (handler CanvasHandler) CreateBlankCanvas() {
	// 20 x 8 canvas
	totalNumFiles := 20 * 8
	dir := handler.CanvasDir
	for i := 0; i < totalNumFiles; i++ {
		fileName := fmt.Sprintf("%s%d.txt", dir, i)
		os.WriteFile(fileName, []byte{}, 0644)
	}
	fmt.Printf("Successfully created %d files", totalNumFiles)
}

func (handler CanvasHandler) DeleteAllFiles() {
	entries, err := os.ReadDir(handler.CanvasDir)
	if err != nil {
		fmt.Println(err)
	}
	for _, e := range entries {
		err = os.Remove(handler.CanvasDir + e.Name())
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("Successfully deleted all files")
}