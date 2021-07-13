package main

import "fmt"

func main() {
	var b byte = 3
	fmt.Println(b)

	i := 3 //interferncia no tipo
	i += 3 // i = i + 3
	i -= 3 // i = i - 3
	i /= 3 // i = i / 3
	i *= 3 // i = i * 3
	i %= 3 // i = i % 3

	fmt.Println(i)

	x, y := 1, 2
	fmt.Println(x, y)

	x, y = 2, 1
	fmt.Println(x, y)

}
