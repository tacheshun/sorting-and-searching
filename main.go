package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Get the number of items and maximum item value.
	var numItems, max int
	fmt.Printf("# Items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	// Make and display an unsorted slice.
	slice := makeRandomSlice(numItems, max)
	printSlice(slice, 40)
	fmt.Println()

	// Sort and display the result.
	bubbleSort(slice)
	printSlice(slice, 40)

	// Verify that it's sorted.
	checkSorted(slice)
}

func makeRandomSlice(numitems, max int) []int {
	slice := []int{}
	for i := 0; i < numitems; i++ {
		slice = append(slice, rand.Intn(max))
	}

	return slice
}

func printSlice(slice []int, numItems int) {
	if len(slice) <= numItems {
		fmt.Println(slice)
	} else {
		fmt.Println(slice[:numItems])
	}
}

func checkSorted(slice []int) {
	prev := slice[0]
	for _, v := range slice {
		if v > prev {
			prev = v
		} else {
			fmt.Println("Slice is not Sorted")
		}
	}
	fmt.Println("The slice is sorted")
}

func bubbleSort(slice []int) {
	prev := slice[0]
	for _, v := range slice {
		if v > prev {
			prev = v
		} else {
			prev, v = v, prev
		}
	}
}
