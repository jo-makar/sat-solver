// Convert a Sudoku puzzle into a SAT problem
//
// Let pijk represent row i, line j and value k (all of range [1,9]).
// These also serve as inputs for the initial puzzle values.

package main

import (
	"fmt"
	"strings"
)

func main() {
	// For each position, allow exactly one value
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			var b strings.Builder
			for k := 1; k <= 9; k++ {
				if k > 1 {
					b.WriteString(" ")
				}
				b.WriteString(fmt.Sprintf("p%d%d%d", i, j, k))
			}
			fmt.Printf(b.String() + "\n")

			for k := 1; k <= 9; k++ {
				for l := k+1; l <= 9; l++ {
					fmt.Printf("~p%d%d%d ~p%d%d%d\n", i,j,k, i,j,l)
				}
			}
		}
	}

	// For each line, allow exactly one instance of each value
	for j := 1; j <= 9; j++ {
		for k := 1; k <= 9; k++ {
			for i := 1; i <= 9; i++ {
				for l := i+1; l <= 9; l++ {
					fmt.Printf("~p%d%d%d ~p%d%d%d\n", i,j,k, l,j,k)
				}
			}
		}
	}

	// For each row, allow exactly one instance of each value
	for i := 1; i <= 9; i++ {
		for k := 1; k <= 9; k++ {
			for j := 1; j <= 9; j++ {
				for l := j+1; l <= 9; l++ {
					fmt.Printf("~p%d%d%d ~p%d%d%d\n", i,j,k, i,l,k)
				}
			}
		}
	}

	// For each box, allow exactly one instance of each value
	boxes := [9][9][2]uint{
		{{1,1}, {2,1}, {3,1}, {1,2}, {2,2}, {3,2}, {1,3}, {2,3}, {3,3}},
		{{4,1}, {5,1}, {6,1}, {4,2}, {5,2}, {6,2}, {4,3}, {5,3}, {6,3}},
		{{7,1}, {8,1}, {9,1}, {7,2}, {8,2}, {9,2}, {7,3}, {8,3}, {9,3}},

		{{1,4}, {2,4}, {3,4}, {1,5}, {2,5}, {3,5}, {1,6}, {2,6}, {3,6}},
		{{4,4}, {5,4}, {6,4}, {4,5}, {5,5}, {6,5}, {4,6}, {5,6}, {6,6}},
		{{7,4}, {8,4}, {9,4}, {7,5}, {8,5}, {9,5}, {7,6}, {8,6}, {9,6}},

		{{1,7}, {2,7}, {3,7}, {1,8}, {2,8}, {3,8}, {1,9}, {2,9}, {3,9}},
		{{4,7}, {5,7}, {6,7}, {4,8}, {5,8}, {6,8}, {4,9}, {5,9}, {6,9}},
		{{7,7}, {8,7}, {9,7}, {7,8}, {8,8}, {9,8}, {7,9}, {8,9}, {9,9}},
	}
	for box := 0; box < len(boxes[box]); box++ {
		for k := 1; k <= 9; k++ {
			for a := 0; a < len(boxes[box]); a++ {
				for b := a+1; b < len(boxes[box]); b++ {
					i1, j1 := boxes[box][a][0], boxes[box][a][1]
					i2, j2 := boxes[box][b][0], boxes[box][b][1]
					fmt.Printf("~p%d%d%d ~p%d%d%d\n", i1,j1,k, i2,j2,k)
				}
			}
		}
	}
}
