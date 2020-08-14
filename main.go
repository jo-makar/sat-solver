// SAT (satisfiability) solver based on https://sahandsaba.com/understanding-sat-by-implementing-a-simple-sat-solver-in-python.html

package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: sat-solver <file>\n")
		os.Exit(1)
	}

	sat := SAT{}

	if file, err := os.Open(os.Args[1]); err != nil {
		log.Panic(err)
	} else {
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if len(line) == 0 || strings.HasPrefix(line, "#") {
				continue
			}
			sat.ParseLine(line)
		}
		if err := scanner.Err(); err != nil {
			log.Panic(err)
		}
	}

	if len(sat.Clauses) == 0 {
		log.Panic("len(sat.Clauses) == 0")
	}

	log.Printf("problem contains %d variables (search space: %.0f) and %d clauses",
	           len(sat.Variables), math.Exp2(float64(len(sat.Variables))), len(sat.Clauses))

	solutions := make(chan string)
	go sat.Solve(solutions)
	for solution := range solutions {
		log.Printf("solution = %s", solution)
	}
}
