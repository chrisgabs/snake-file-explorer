package game

import (
	"fmt"
	"os"
	canvasHandler "snek/canvas_handler"
	"sort"
	"strconv"
	"strings"
)

var cellTypeFileTypeMapping map[int]string

func init() {
	cellTypeFileTypeMapping = make(map[int]string)
	cellTypeFileTypeMapping[Empty] = ".txt"
	cellTypeFileTypeMapping[SnakePart] = ".rar"
	cellTypeFileTypeMapping[Food] = ".docx"
}

type SnakeCanvasUpdater struct {
	board Board
}

// Assumes that canvas is already populated
func (handler SnakeCanvasUpdater) UpdateCanvas() {
	fmt.Println(handler.board)
	
	dir := canvasHandler.CanvasDir
	entries, err := os.ReadDir(dir)
	sort.Slice(entries, func(i,j int) bool{
		iVal, _ := strconv.Atoi(strings.Split(entries[i].Name(), ".")[0])
		jVal, _ := strconv.Atoi(strings.Split(entries[j].Name(), ".")[0])
    	return iVal < jVal
	})
	if err != nil {
		panic("panic reading dir")
	}


	var newFileName string
	cells := handler.board.cells

	for row := 0; row < handler.board.height; row++ {
		position := row*handler.board.width
		for col := 0; col < handler.board.width; col++ {
			existingFileName := fmt.Sprintf("%s%s", dir, entries[position].Name())
			newFileName = fmt.Sprintf("%s%d%s", dir, position, cellTypeFileTypeMapping[cells[row][col].CellType])
			updateFile(existingFileName, newFileName)
			position++
		}
	}

}

func updateFile(existingFileName string, newFileName string) {
	if existingFileName == newFileName {
		return
	}else {
		os.Rename(existingFileName, newFileName)
	}
}