package main

import "fmt"

func main() {
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a) - 2])

	var p [3]int = [3]int{1, 2, 3}
	fmt.Println(p)

	q := [...]int{1, 2, 3}
	fmt.Println(q)
}