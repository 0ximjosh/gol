package core

const NUM_ALGOS = 2

type CellAlgo func(grid, buffer [][]uint8, x, y int)

func SingleCell(grid, buffer [][]uint8, x, y int) {
	n := grid[x-1][y-1] +
		grid[x-1][y+0] +
		grid[x-1][y+1] +
		grid[x+0][y-1] +
		grid[x+0][y+1] +
		grid[x+1][y-1] +
		grid[x+1][y+0] +
		grid[x+1][y+1]

	switch {
	case grid[x][y] == 0 && n == 3:
		buffer[x][y] = 1
	case n < 2 || n > 3:
		buffer[x][y] = 0
	default:
		buffer[x][y] = grid[x][y]
	}
}

func TwoCell(grid, buffer [][]uint8, x, y int) {
	n := grid[x-1][y-1] +
		grid[x-1][y+0] +
		grid[x-1][y+1] +
		grid[x+0][y-1] +
		grid[x+0][y+1] +
		grid[x+1][y-1] +
		grid[x+1][y+0] +
		grid[x+1][y+1]

	switch {
	case grid[x][y] == 0 && n == 3:
		buffer[x][y] = 1
	case n < 2:
		buffer[x][y] = 0
	case grid[x][y] == 1 && n > 3:
		buffer[x][y] = 2
	case n > 3:
		buffer[x][y] = 0
	default:
		buffer[x][y] = grid[x][y]
	}
}

func GetAlgo(index int) CellAlgo {
	switch index {
	case 1:
		return TwoCell
	default:
		return SingleCell
	}
}
