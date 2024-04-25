package main

import "testing"

func TestPush(t *testing.T) {
	s := stack[rune]{elements: []rune{1}}
	s.push(2)
	if s.elements[0] != 1 || s.elements[1] != 2 {
		t.Fatal("Wrong")
	}
}

func TestPop(t *testing.T) {
	s := stack[rune]{elements: []rune{1, 2}}
	n := s.pop()
	if *n != 2 {
		t.Fatal("Wrong")
	}
	if len(s.elements) != 1 {
		t.Fatal("Wrong")
	}
}

func TestGetOperatorType(t *testing.T) {
	op := getOperatorType('(')
	if op != openBracket {
		t.Fatalf("Expected: %v, Got: %v", openBracket, op)
	}
	op = getOperatorType(')')
	if op != closeBracket {
		t.Fatalf("Expected: %v, Got: %v", closeBracket, op)
	}
	op = getOperatorType('*')
	if op != otherOperator {
		t.Fatalf("Expected: %v, Got: %v", closeBracket, op)
	}
}

func TestIsBalanced(t *testing.T) {
	expr := "2*(5+2)"
	if !isBalanced(expr) {
		t.Fatalf("Expected: %v, Got: %v", true, false)
	}
	expr = "{5+[2*(5+2)]}"
	if !isBalanced(expr) {
		t.Fatalf("Expected: %v, Got: %v", true, false)
	}
	expr = "2*(5+2"
	if isBalanced(expr) {
		t.Fatalf("Expected: %v, Got: %v", false, true)
	}
	expr = ")"
	if isBalanced(expr) {
		t.Fatalf("Expected: %v, Got: %v", false, true)
	}
}
