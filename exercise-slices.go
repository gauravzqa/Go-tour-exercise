package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	tr := make([][]uint8, dy)
	for i := range tr {
		tr[i] = make([]uint8, dx)
	}
	for i := uint8(0); i < uint8(dy); i++ {
		for j := uint8(0); j < uint8(dx); j++ {
			tr[i][j] = (i * j)
		}
	}
	return tr
}

func main() {
	pic.Show(Pic)
}
