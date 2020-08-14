package main

import (
	"log"
	"strings"
)

// - Variables are assigned unique integers starting from zero and increasing as they are encountered
//   - Variable names can be composed of anything other than whitespace and ~ (which is used to indicate literal negation)
// - Given a variable encoded as the integer x, a positive literal is represented by 2x and a negative literal by 2x+1
// - Clauses are simply arrays/slices of encoded literals

// The preceding encoding thereby supports:
// - Tracking all variables with an array of length n and all literals with an array of length 2n
// - Fast bitwise operations for the following:
//   - Bitwise and with one determines if a literal is positive or negative
//   - Bitwise left shift of one to identify the variable associated with a literal  
//   - Bitwise xor with one to toggle between positive and negative literals
//   - Bitwise right shift of one to make a positive literal from a variable
//     - With a bitwise or with one to make a negative literal

const unassigned uint = 2

type SAT struct {
	Variables map[string]uint
	Clauses   []([]uint)
}

func (sat *SAT) ParseLine(line string) {
	log.Printf("processing line = %q", line)

	var clause []uint
	for _, literal := range strings.Fields(line) {
		var variable string
		var negated bool
		if strings.HasPrefix(literal, "~") {
			negated = true
			variable = literal[1:]
		} else {
			variable = literal
		}

		if sat.Variables == nil {
			sat.Variables = make(map[string]uint)
		}
		if _, ok := sat.Variables[variable]; !ok {
			sat.Variables[variable] = uint(len(sat.Variables))
		}

		encodedLiteral := sat.Variables[variable] << 1
		if negated {
			encodedLiteral |= 1
		}
		clause = append(clause, encodedLiteral)
	}

	if len(clause) == 0 {
		log.Panic("len(clause) == 0")
	}

	// TODO Check prior to storing to verify the clause is not a duplicate.
	//      Can be done by sorting and storing in a map (pseudo-set) temporarily.

	sat.Clauses = append(sat.Clauses, clause)
}

func (sat *SAT) Solve(solutions chan string) {
	verbose := false

	if len(sat.Clauses) == 0 {
		log.Panic("len(sat.Clauses) == 0")
	}

	// For each clause to be satisfied at least one of its literals must be satisfied.
	// Therefore arrange for each clause to watch one of its literals and ensure the following invariant is always maintained:
	//   All watched literals are either satisfied or are not yet assigned a value.
	// Then each time a variable is assigned the watched literals (or watchlist) is updated for unsatisfied clauses,
	// and for the cases where a clause cannot be satisfied the variable assignment is rolled back.

	// Initialize the watchlist such that each clause is watching its first literal.
	watchlist := make(map[uint]([]([]uint)))
	for _, clause := range sat.Clauses {
		watchlist[clause[0]] = append(watchlist[clause[0]], clause)
	}

	assignment := make([]uint, len(sat.Variables))
	for i := range assignment {
		assignment[i] = unassigned
	}

	// Maintain state for each variable such that:
	// - Zero indicates nothing has been attempted
	// - One means false has been attempted
	// - Two means true has been attempted
	// - Three means both false and true have been attempted
	const (
		falseAttempted uint = 0x01
		trueAttempted = 0x02
	)
	state := make([]uint, len(sat.Variables))

	updateWatchlist := func(unsatisfiedLiteral uint) bool {
		for len(watchlist[unsatisfiedLiteral]) > 0 {
			clause := watchlist[unsatisfiedLiteral][0]

			foundAlternative := false
			for _, literal := range clause {
				variable := literal >> 1
				positive := literal & 1  == 0
				if assignment[variable] == unassigned || (!positive && assignment[variable] == 0) || (positive && assignment[variable] == 1) {
					foundAlternative = true
					watchlist[unsatisfiedLiteral] = watchlist[unsatisfiedLiteral][1:]
					watchlist[literal] = append(watchlist[literal], clause)
					break
				}
			}

			if !foundAlternative {
				if verbose {
					log.Printf("assignment %s contradicts clause %s",
					           sat.subAssignmentToString(assignment, clause), sat.clauseToString(clause))
				}
				return false
			}
		}
		return true
	}

	var i uint = 0
	for {
		if i == uint(len(sat.Variables)) {
			solutions <- sat.assignmentToString(assignment)
			i--
			continue
		}

		madeAttempt := false
		for _, a := range []uint{falseAttempted, trueAttempted} {
			if state[i] & a == 0 {
				if verbose {
					log.Printf("attempting %s = %d", sat.variableToString(i), a >> 1)
				}
				madeAttempt = true
				state[i] |= a
				assignment[i] = a >> 1

				oppositeLiteral := i << 1 | (a >> 1)
				if updateWatchlist(oppositeLiteral) {
					i++
					break
				} else {
					assignment[i] = unassigned
				}
			}
		}

		if !madeAttempt {
			if i == 0 {
				close(solutions)
				return
			} else {
				state[i] = 0
				assignment[i] = unassigned
				i--
			}
		}
	}

	// TODO For complex problems, launch a goroutine that periodically calculates how much of the search space has been covered via log entries.
	//      This calculation can be done by summing the values of state and dividing by len(state) * 3 for a percentage of covered search space.
}
