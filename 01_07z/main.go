package main

import (
	"flag"
	"log"
)

// operatorType represents the type of operator in an expression
type operatorType rune

const (
	openBracket operatorType = iota
	closeBracket
	otherOperator
)

// bracketPairs is the map legal bracket pairs
var bracketPairs = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
}

// getOperatorType return the operator type of the given operator.
func getOperatorType(op rune) operatorType {
	for ob, cb := range bracketPairs {
		switch op {
		case ob:
			return openBracket
		case cb:
			return closeBracket
		}
	}
	return otherOperator
}

// stack is a simple LIFO stack implementation using a slice.
type stack[T any] struct {
	elements []T
}

// push adds a new element to the stack.
func (s *stack[T]) push(e T) {
	s.elements = append(s.elements, e)
}

// pop removes the last element from the stack.
func (s *stack[T]) pop() *T {
	if len(s.elements) == 0 {
		return nil
	}
	n := len(s.elements) - 1
	last := s.elements[n]
	s.elements = s.elements[:n]
	return &last
}

// isBalanced returns whether the given expression
// has balanced brackets.
func isBalanced(expr string) bool {
	s := stack[rune]{}
	for _, c := range expr {
		switch getOperatorType(c) {
		case openBracket:
			s.push(c)
		case closeBracket:
			last := s.pop()
			if last == nil || bracketPairs[*last] != c {
				return false
			}
		}
	}
	return len(s.elements) == 0
}

// printResult prints whether the expression is balanced.
func printResult(expr string, balanced bool) {
	if balanced {
		log.Printf("%s is balanced.\n", expr)
		return
	}
	log.Printf("%s is not balanced.\n", expr)
}

func main() {
	expr := flag.String("expr", "", "The expression to validate brackets on.")
	flag.Parse()
	printResult(*expr, isBalanced(*expr))
}
