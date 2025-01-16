package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/inancgumus/screen"
)

var size_x = 32
var size_y = 32

const alive_chance float32 = 0.3

var generations int32 = 0

var board [][]int = make([][]int, size_y)

func main() {

	initBoard(board)
	randomizeBoard(board)
	screen.Clear()
	drawBoard(board)
	time.Sleep(5 * time.Second)

	for range 1000 {

		generations++

		// Iterate and display
		iterateGameOfLife(board)
		drawBoard(board)
		// Sleep to achieve almost desired framerate
		time.Sleep(1000 / 24 * time.Millisecond)

	}

	fmt.Println("Done!")
}

func initBoard(board [][]int) {
	size_x, size_y = screen.Size()
	size_x /= 2
	println("Initializing board...")
	for y := range board {
		for range size_x {
			board[y] = append(board[y], 0)
		}
	}
}

func setAlive(board [][]int, x int, y int) {
	if len(board) > y && len(board[y]) > x {
		board[y][x] = 1
	} else {
		fmt.Println("Position illegal:", x, y)
	}

}

func setDead(board [][]int, x int, y int) {
	if len(board) > y && len(board[y]) > x {
		board[y][x] = 0
	} else {
		fmt.Println("Position illegal:", x, y)
	}
}

func randomizeBoard(board [][]int) {
	for y := range board {
		for x := range board[y] {
			rNumber := rand.Float32()
			if rNumber <= alive_chance {
				board[y][x] = 1
			} else {
				board[y][x] = 0
			}

		}
	}
}

func checkState(board [][]int, x int, y int) int {
	//println("Checking state...", x, y)
	if int(x) >= 0 && int(y) >= 0 && len(board) > int(y) && len(board[y]) > int(x) {
		return board[y][x]
	} else if int(x) >= size_x {
		return checkState(board, x-size_x, y)
	} else if int(y) >= size_y {
		return checkState(board, x, y-size_y)
	} else if int(x) < 0 {
		return checkState(board, x+size_x, y)
	} else if int(y) < 0 {
		return checkState(board, x, y+size_y)
	}
	return -1
}

func drawBoard(board [][]int) {
	screen.Clear()
	var sb strings.Builder
	println("\n")
	for y := range board {
		for x := range board[y] {
			if board[y][x] == 0 {
				sb.WriteString("- ")
			} else {
				sb.WriteString("X ")
			}

			if x == size_x-1 {
				sb.WriteString("\n")
			}
		}
	}
	fmt.Print(sb.String())
	fmt.Printf("Generation %d\n", generations)
}

func checkNeighbours(testBoard [][]int, x int, y int) int {
	//println("Checking neighbours...")
	var neighbours int = 0
	if checkState(testBoard, x+1, y) == 1 {
		neighbours++
	}
	if checkState(testBoard, x+1, y+1) == 1 {
		neighbours++
	}
	if checkState(testBoard, x, y+1) == 1 {
		neighbours++
	}
	if checkState(testBoard, x-1, y+1) == 1 {
		neighbours++
	}
	if checkState(testBoard, x-1, y) == 1 {
		neighbours++
	}
	if checkState(testBoard, x-1, y-1) == 1 {
		neighbours++
	}
	if checkState(testBoard, x, y-1) == 1 {
		neighbours++
	}
	if checkState(testBoard, x+1, y-1) == 1 {
		neighbours++
	}
	return neighbours
}

func iterateGameOfLife(board [][]int) {
	var originalState [][]int = make([][]int, len(board))
	for y := range board {
		originalState[y] = make([]int, len(board[y]))
	}
	for y := range board {
		copy(originalState[y], board[y])
	}

	for y := range originalState {
		for x := range originalState[y] {
			var neighbours int = checkNeighbours(originalState, int(x), int(y))

			if originalState[y][x] == 0 {
				// Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
				if neighbours == 3 {
					setAlive(board, x, y)
					//println(x, y, "get born, neighbours:", neighbours)
				}
			} else {

				// Any live cell with fewer than two live neighbours dies, as if by underpopulation.
				if neighbours < 2 {
					setDead(board, x, y)
					//println(x, y, "dies by underpop, neighbours:", neighbours)
				}

				// Any live cell with more than three live neighbours dies, as if by overpopulation.
				if neighbours > 3 {
					setDead(board, x, y)
					//println(x, y, "dies by overpop, neighbours:", neighbours)

				}

				// Any live cell with two or three live neighbours lives on to the next generation.
			}
		}
	}

}
