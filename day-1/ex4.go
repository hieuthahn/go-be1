package main

import (
	"fmt"
	"sort"
)

// Two Sum
func TwoSum() {
	var length, target int
	var numbers []int
	var hashMap = make(map[int]int)
	var result []int

	for length <= 0 {
		fmt.Println("Input length of numbers: ")
	  fmt.Scan(&length)
	}

	for i := 0; i < length; i++ {
		var number int
		fmt.Printf("Input number %d: ", i+1)
		fmt.Scan(&number)
		numbers = append(numbers, number)
	}

	fmt.Println("Input target: ")
	fmt.Scan(&target)

	fmt.Printf("Two Sum: %v, Target: %d\n", numbers, target)

	for i := 0; i < len(numbers); i++ {
		var minus = target - numbers[i]
		
		if (hashMap[minus] != 0) {
			result = append(result, i)
			result = append(result, hashMap[minus])
			sort.Ints(result)
			break
		}

		hashMap[numbers[i]] = i
	}

	fmt.Printf("Result: %v\n", result)
}
