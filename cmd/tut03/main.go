package main

import "fmt"

func main() {
	var p *int32 = new(int32)
	var i int32
	fmt.Printf("The value of  p points to is %v", *p)
	fmt.Printf("\nThe value of i is: %v", i)

	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\nThe memory location of t1 is: %p", &thing1)
	var result [5]float64 = square(&thing1)
	fmt.Printf("\nThe result is: %v", result)

}

func square(thing2 *[5]float64) [5]float64 {
	fmt.Printf("\nThe memory location of t2 is: %p", thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return *thing2
}
