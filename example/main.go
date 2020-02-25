package main

import "fmt"

func CheckNumber(n int) string {
	if n <=0 {
		return "small number"
	}
	if n > 0 && n <= 100 {
		return "medium number"
	}else {
		return "big number"
	}
}

func main() {
	fmt.Println(CheckNumber(100))
}