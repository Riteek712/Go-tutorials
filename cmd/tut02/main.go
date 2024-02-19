package main

import "fmt"

type gasEngine struct {
	mgp     uint8
	gallons uint8
}
type electricEngine struct {
	kwh   uint8
	mpkwh uint8
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mgp
}

type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Printf("You can make it!")
	} else {
		fmt.Println("Need to fuel up first!")
	}
}

func main() {
	var myString = "resume"
	var indexed = myString[0]
	fmt.Printf("%v, %T\n", indexed, indexed)
	for i, v := range myString {
		fmt.Println(i, v)
	}

	var myEngine gasEngine = gasEngine{25, 15}
	canMakeIt(myEngine, 50)
}
