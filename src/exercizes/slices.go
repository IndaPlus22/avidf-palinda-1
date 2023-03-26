package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	slice := make([][]uint8, dy)
	for y := range slice {
		row := make([]uint8, dx)
		for x := range row {
			row[x] = uint8(x * y)
		}
		slice[y] = row
	}
	return slice
}

func main() {
	pic.Show(Pic)
}
