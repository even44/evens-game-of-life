package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

const size_x = 16
const size_y = 16

var generations int32 = 0

var board [][]int8 = make([][]int8, size_y)

func main() {

	initBoard(board)

	randomizeBoard(board)

	drawBoard(board)

	time.Sleep(5 * time.Second)

	for range 500 {

		generations++

		// Clear screen
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()

		// Iterate and display
		iterateGameOfLife(board)
		drawBoard(board)

		// Sleep to achieve almost desired framerate
		time.Sleep(1000 / 24 * time.Millisecond)

	}

	//drawBoard(board)

	fmt.Println("Done!")
}

func initBoard(board [][]int8) {
	println("Initializing board...")
	for y := range board {
		for range size_x {
			board[y] = append(board[y], 0)
		}
	}
}

func setAlive(board [][]int8, x int, y int) {
	if len(board) > y && len(board[y]) > x {
		board[y][x] = 1
	} else {
		fmt.Println("Position illegal:", x, y)
	}

}

func setDead(board [][]int8, x int, y int) {
	if len(board) > y && len(board[y]) > x {
		board[y][x] = 0
	} else {
		fmt.Println("Position illegal:", x, y)
	}
}

func randomizeBoard(board [][]int8) {
	for y := range board {
		for x := range board[y] {
			rNumber := rand.Float32()
			if rNumber <= 0.5 {
				board[y][x] = 0
			} else {
				board[y][x] = 1
			}

			if x == size_x-1 {
				fmt.Print("\n")
			}
		}
	}
}

func checkState(board [][]int8, x int8, y int8) int8 {
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

func drawBoard(board [][]int8) {
	println("\n")
	for y := range board {
		for x := range board[y] {
			if board[y][x] == 0 {
				fmt.Print("- ")
			} else {
				fmt.Print("X ")
			}

			if x == size_x-1 {
				fmt.Print("\n")
			}
		}
	}
	fmt.Printf("Generation %d\n", generations)
}

func checkNeighbours(testBoard [][]int8, x int8, y int8) int8 {
	//println("Checking neighbours...")
	var neighbours int8 = 0
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

func iterateGameOfLife(board [][]int8) {
	var originalState [][]int8 = make([][]int8, len(board))
	for y := range board {
		originalState[y] = make([]int8, len(board[y]))
	}
	for y := range board {
		copy(originalState[y], board[y])
	}
	//print("\nPrinting original")
	//drawBoard(originalState)

	for y := range originalState {
		for x := range originalState[y] {
			var neighbours int8 = checkNeighbours(originalState, int8(x), int8(y))
			
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
	//print("\nPrinting original")
	//drawBoard(originalState)
}
