// SAT (satisfiability) solver based on https://sahandsaba.com/understanding-sat-by-implementing-a-simple-sat-solver-in-python.html

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	var setOnly bool

	flag.BoolVar(&setOnly, "set", false, "display only set variables in solutions")
	flag.BoolVar(&setOnly, "s", false, "display only set variables in solutions")

	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Fprintf(os.Stderr, "usage: sat-solver [flags] <file>\n")
		os.Exit(1)
	}

	sat := SAT{}

	if file, err := os.Open(flag.Args()[0]); err != nil {
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
		if setOnly {
			var b strings.Builder
			first := true
			for _, l := range strings.Fields(solution) {
				if strings.HasPrefix(l, "~") {
					continue
				}
				if !first {
					b.WriteString(" ")
				}
				b.WriteString(l)
				first = false
			}
			solution = b.String()
		}

		log.Printf("solution = %s", solution)
	}
}
