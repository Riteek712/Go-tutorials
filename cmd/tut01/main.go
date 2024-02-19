package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("Hello world 2!")
	var intNum int = 32767
	fmt.Println(intNum)

	// Typecasting
	var floatnum32 float32 = 10.1
	var intNum32 = 2
	var result1 float32 = floatnum32 + float32(intNum32)
	fmt.Println(result1)
	var myString string = "Hello" + " " + "World"
	fmt.Println(myString)

	fmt.Println(utf8.RuneCountInString("&"))

	var myRune rune = 'a'
	fmt.Println(myRune)

	var1, var2 := 11, 2
	fmt.Println(var1, var2)

	var result, remainder, err = intDivision(var1, var2)
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Printf("the result of integer division is %v with remainder %v.", result, remainder)
	var result2, remainder2, err2 = intDivision(var1, 0)
	if err2 != nil {
		fmt.Printf(err.Error())
	}
	fmt.Printf("the result of integer division is %v with remainder %v.", result2, remainder2)
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Cannot Divide by Zero.")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
