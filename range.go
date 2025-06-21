package main

import "fmt"

func main() {
	var a float32
	a = 3.1415926535
	fmt.Printf("Value: %f and type: %T\n", a, a)

	var b float64
	b = 2.718281828
	fmt.Printf("Value: %f and type: %T\n", b, b)

	var c int16
	c = 12
	fmt.Printf("Value: %d and type: %T\n", c, c)

	var e int64
	e = 1221
	fmt.Printf("Value: %d and type: %T\n", e, e)

	var ak string
	ak = "Aditya"
	fmt.Printf("Value: %s and type: %T\n", ak, ak)

}
