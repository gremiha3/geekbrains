package main

import "fmt"

var n = int(5)

func main() {
	i := 0
	f := 1

	for i < n {
		f = f * (i + 1)
		i++
	}
	fmt.Printf("Факториал %d! = %d\n", n, f)
}
