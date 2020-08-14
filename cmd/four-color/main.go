// Convert the Four color theorem (https://en.wikipedia.org/wiki/Four_color_theorem) into a SAT problem.
//
// Each region is numbered (1 to n) and neighbors are provided one per line.
// The variables are ri, bi, gi, yi to indicate region i is colored red, blug, green or yellow.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: four-color <file>\n")
		os.Exit(1)
	}

	neighbors := make(map[uint]([]uint))

	if file, err := os.Open(os.Args[1]); err != nil {
		log.Panic(err)
	} else {
		has := func(s []uint, x uint) bool {
			for _, e := range s {
				if e == x {
					return true
				}
			}
			return false
		}

		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if len(line) == 0 {
				continue
			}

			//
			// Parse the input
			//

			log.Printf("processing line = %q", line)
			pair := strings.Fields(line)
			if len(pair) != 2 {
				log.Panic("len(pair) != 2")
			}

			var p [2]int
			for i := 0; i < 2; i++ {
				if x, err := strconv.Atoi(pair[i]); err != nil {
					log.Panic(err)
				} else if x < 1 {
					log.Panic("p[%d] < 1", i)
				} else {
					p[i] = x
				}
			}
			if p[0] == p[1] {
				log.Panic("p[0] == p[1]")
			}

			//
			// Ensure that each pairing is stored only once
			//

			if has(neighbors[uint(p[0])], uint(p[1])) || has(neighbors[uint(p[1])], uint(p[0])) {
				log.Printf("%d %d pair previously stored", p[0], p[1])
			} else {
				neighbors[uint(p[0])] = append(neighbors[uint(p[0])], uint(p[1]))
			}
		}
		if err := scanner.Err(); err != nil {
			log.Panic(err)
		}
	}

	if len(neighbors) == 0 {
		log.Panic("len(neighbors) == 0")
	}

	regions := make(map[uint]bool)
	for k, v := range neighbors {
		regions[k] = true
		for _, w := range v {
			regions[w] = true
		}
	}

	n := 0
	for _, v := range neighbors {
		n += len(v)
	}
	log.Printf("problem contains %d regions and %d neighbors", len(regions), n)

	//
	// Convert to a SAT problem
	//

	// For each region, assign exactly one color
	for i := range regions {
		fmt.Printf("r%d b%d g%d y%d\n", i, i, i, i)
		fmt.Printf("~r%d ~b%d\n", i, i)
		fmt.Printf("~r%d ~g%d\n", i, i)
		fmt.Printf("~r%d ~y%d\n", i, i)
		fmt.Printf("~b%d ~g%d\n", i, i)
		fmt.Printf("~b%d ~y%d\n", i, i)
		fmt.Printf("~g%d ~y%d\n", i, i)
		fmt.Printf("\n")
	}

	// For each pair of neighbors, ensure they are not the same color
	for k, v := range neighbors {
		for _, w := range v {
			fmt.Printf("~r%d ~r%d\n", k, w)
			fmt.Printf("~b%d ~b%d\n", k, w)
			fmt.Printf("~g%d ~g%d\n", k, w)
			fmt.Printf("~y%d ~y%d\n", k, w)
			fmt.Printf("\n")
		}
	}
}
