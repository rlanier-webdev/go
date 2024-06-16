/*Write a function that returns the total surface area and volume of a box as an array: [area, volume]
SA = 2(LW + LH + WH)
V = L x W x H
*/

package main

import "fmt"

func main() {
	var (
		w int = 5
		h int = 3
		d int = 5
	)

	s := getSurfaceArea(w,h,d)
	v := getVolume(w,h,d)
	fmt.Println([2]int{s,v})
}

func getSurfaceArea(w,h,d int) int {
	return 2*(w*h + h*d + w*d)
}

func getVolume(w,h,d int) int {
	return w*h*d
}