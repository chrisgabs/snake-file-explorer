package canvas_handler

import (
	"fmt"
	"os"
	"path/filepath"
)

var CanvasDir string

type CanvasUpdater interface {
	UpdateCanvas()
}

type CanvasHandler struct {
	Updater CanvasUpdater
}

func init() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	path := filepath.Clean(filepath.Join(dir, ".."))
	path = fmt.Sprintf("%s/canvas/", path)
	CanvasDir = path
}

func (handler CanvasHandler) CreateBlankCanvas(width, height int) {
	// 20 x 8 canvas
	totalNumFiles := width * height
	dir := CanvasDir
	for i := 0; i < totalNumFiles; i++ {
		fileName := fmt.Sprintf("%s%d.txt", dir, i)
		os.WriteFile(fileName, []byte{}, 0644)
	}
	fmt.Printf("Successfully created %d files", totalNumFiles)
}

func (handler CanvasHandler) DeleteAllFiles() {
	entries, err := os.ReadDir(CanvasDir)
	if err != nil {
		fmt.Println(err)
	}
	for _, e := range entries {
		err = os.Remove(CanvasDir + e.Name())
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("Successfully deleted all files")
}