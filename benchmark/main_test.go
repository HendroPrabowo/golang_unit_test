package main

import "testing"

func BenchmarkIsPalindrome(b *testing.B) {
	input := "kasurnababanrusak"
	for i:=0; i<b.N; i++ {
		IsPalindrome(input)
	}
}

func BenchmarkIsPalindromeLoop(b *testing.B) {
	input := "kasurnababanrusak"
	for i:=0; i<b.N; i++ {
		IsPalindromeLoop(input)
	}
}
