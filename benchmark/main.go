package main

func IsPalindrome(text string) bool {
	if len(text) <= 1 {
		return true
	}
	if len(text) == 2 && text[0] == text[1] {
		return true
	}
	if text[0] != text[len(text)-1] {
		return false
	}
	return IsPalindrome(text[1 : len(text)-1])
}

func IsPalindromeLoop(text string) bool {
	n := len(text)/2
	for i:=0; i<n; i++{
		if text[i] != text[len(text)-1-i] {
			return false
		}
	}
	return true
}