package table

func CreateGrid() [7][7]string {

	var grid [7][7]string
	return grid
}

func CountOfShipOnBoard(grid [7][7]string) int {
	count := 0

	for _, row := range grid {
		for _, location := range row {
			if location == "Ship" {
				count++
			}

		}
	}
	return count
}
