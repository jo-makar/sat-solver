package main

import (
	"log"
	"strings"
)

func (sat *SAT) variableToString(variable uint) string {
	var s string
	for k, v := range sat.Variables {
		if v == variable {
			s = k
			break
		}
	}
	if s == "" {
		log.Panic("invalid variable")
	}
	return s
}

func (sat *SAT) literalToString(literal uint) string {
	var s string
	variable := literal >> 1
	for k, v := range sat.Variables {
		if v == variable {
			s = k
			break
		}
	}
	if s == "" {
		log.Panic("invalid literal")
	}

	if literal & 1 == 1 {
		s = "~" + s
	}
	return s
}

func (sat *SAT) clauseToString(clause []uint) string {
	var b strings.Builder
	for index, literal := range clause {
		if index > 0 {
			b.WriteString(" ")
		}
		b.WriteString(sat.literalToString(literal))
	}
	return b.String()
}

func (sat *SAT) subAssignmentToString(assignment []uint, clause []uint) string {
	var b strings.Builder
	first := true
	for _, literal := range clause {
		variable := literal >> 1
		if assignment[variable] != unassigned {
			if !first {
				b.WriteString(" ")
			}

			var s string
			for k, v := range sat.Variables {
				if v == variable {
					s = k
					break
				}
			}
			if s == "" {
				log.Panic("invalid literal")
			}

			if assignment[variable] == 0 {
				s = "~" + s
			}
			b.WriteString(s)
			first = false
		}
	}
	return b.String()
}

func (sat *SAT) assignmentToString(assignment []uint) string {
	var b strings.Builder
	first := true
	for s, variable := range sat.Variables {
		if assignment[variable] != unassigned {
			if !first {
				b.WriteString(" ")
			}
			if assignment[variable] == 0 {
				s = "~" + s
			}
			b.WriteString(s)
			first = false
		}
	}
	return b.String()
}
