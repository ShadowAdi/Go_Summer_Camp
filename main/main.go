package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func returnString(a, b string) (string, string) {
	return b, a
}

func main() {
	var ans, ans1 string
	ans, ans1 = returnString("Adi", "Lucky")
	fmt.Println(ans)
	fmt.Println(ans1)

}
