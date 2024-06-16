/*Create a method each_cons that accepts a list and a number n, and returns cascading subsets of the list of size n, like so:
each_cons([1,2,3,4], 2)
	#=> [[1,2], [2,3], [3,4]]

each_cons([1,2,3,4], 3)
	#=> [[1,2,3],[2,3,4]]
*/

package main

import "fmt"

func main() {
	each_cons([]int{1, 2, 3, 4}, 1)
}

func each_cons(list []int, n int) [][]int {
	for i := 0; i < len(list); i++ {
		if i + n <= len(list) {
			fmt.Print(list[i:i+n])
		}
	}
	return nil
}