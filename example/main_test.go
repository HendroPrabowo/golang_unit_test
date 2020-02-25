package main

import "fmt"

func ExampleCheckNumber() {
	fmt.Println(CheckNumber(100))
	// Output:
	// medium number
}

func ExampleWrongCheckNumberUnitTest() {  // This is example wrong unit test
	fmt.Println(CheckNumber(100))
	// Output:
	// big number
}