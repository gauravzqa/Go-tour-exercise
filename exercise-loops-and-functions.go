package main

import (
	"fmt"
)

func Sqrt(x float64) (prev float64) {
	prev = 0.0
	for i := 1; i <= (int(x)/2)+1; i++ {
		// fmt.Println(x, i)
		if z := float64(i * i); x-z >= 0 {
			prev = float64(i)
		} else {
			return prev
		}
	}
	return
}

func main() {
	fmt.Println(Sqrt(10))
}
