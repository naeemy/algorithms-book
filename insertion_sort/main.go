package main

import (
	"fmt"
)

func main() {
	list := []int{5, 2, 4, 6, 1, 3}
	insertionSort(list)
}

func insertionSort(list []int) {
	for i := 1; i < len(list); i++ {
		key := list[i]
		j := i-1
		for j >= 0 && key < list[j] {
				list[j+1] = list[j]
				j--
			}
		list[j+1] = key
	}
	fmt.Println(list)
}