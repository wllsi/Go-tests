// INSERTION SORT
// return arr unnecessary due to memory address assignment

package main

import "fmt"

func insertionSort(x int, arr []int) {
	for i := 0; i < x; i++ {
		j := i
		for j >= 0 && arr[j] > arr[j+1] {
			s := arr[j]
			arr[j] = arr[j+1]
			arr[j+1] = s
			j--
		}
	}
}

func main() {
	num := 5
	puzzle1 := []int{5, 6, 2, 4, 3, 1}
	insertionSort(num, puzzle1)
	fmt.Println(puzzle1)
}