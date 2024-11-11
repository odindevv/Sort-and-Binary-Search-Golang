package main

import "fmt"

func main() {
	numbers := []int{35, 12, 45, 1, 5, 99, 120, 32, 6, 2, 9}
	fmt.Printf("Arreglo sin ordenar:\n%v\n", numbers)
	sortedNumbers := bubbleSort(numbers)
	targetIndex := binarySearch(numbers, 45)
	fmt.Printf("Arreglo ordenado con el algoritmo Bubble Sort:\n%v\n", sortedNumbers)
	fmt.Printf("Búsqueda de número utilizando el algoritmo Binary Search, target: %d\nÍndice en el arreglo: %d\n", 45, targetIndex)

}

func bubbleSort(nums []int) []int {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
				swapped = true
			}
		}
	}

	return nums
}

func binarySearch(nums []int, target int) int {
	start := 0
	end := len(nums) - 1

	for start <= end {
		middle := (start + end) / 2
		if nums[middle] == target {
			return middle
		} else if nums[middle] > target {
			end = middle - 1
		} else {
			start = middle + 1
		}
	}

	return -1

}
