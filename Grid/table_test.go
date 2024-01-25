package table

import (
	"testing"
)

func TestCreateGrid(t *testing.T){
	grid := CreateGrid()
	assertGrisIsCorrectSize(t, grid, 7, 7)
	}


func assertGrisIsCorrectSize(t *testing.T, grid [7][7]string, expectedRows int, expectedCols int){
	gridSizeCols := len(grid)
	if gridSizeCols != expectedCols{
		t.Error("Grid is wrong size. Expected exact size of 7, got: ", gridSizeCols)
	}

	gridSizeRows:= len(grid[0])
	if gridSizeRows!= expectedRows{
		t.Error("Grid has wrong number of rows. Want 7, but was", gridSizeRows)
	}
}




