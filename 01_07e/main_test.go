package main

import "testing"

func TestPush(t *testing.T) {
	s := stack{elems: []rune{1}}
	s.push(2)
	if s.elems[0] != 1 || s.elems[1] != 2 {
		t.Fatal("Wrong")
	}
}

func TestPop(t *testing.T) {
	s := stack{elems: []rune{1, 2}}
	n := s.pop()
	if *n != 2 {
		t.Fatal("Wrong")
	}
	if len(s.elems) != 1 {
		t.Fatal("Wrong")
	}
}

func TestGetOperatorType(t *testing.T) {
	op := getOperatorType('(')
	if op != openBracket {
		t.Fatalf("Expected: %v, Got: %v", openBracket, op)
	}
	op = getOperatorType(')')
	if op != closedBracket {
		t.Fatalf("Expected: %v, Got: %v", closedBracket, op)
	}
	op = getOperatorType('*')
	if op != otherOperator {
		t.Fatalf("Expected: %v, Got: %v", closedBracket, op)
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
