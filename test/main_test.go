package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	expectation := true
	actual := IsPalindrome("kayak")
	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}