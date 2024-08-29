package main

import (
	"fmt"
	"math"
	"sort"
)

// Viết chương trình nhập một slice số, in ra tổng, số lớn nhất, số nhỏ nhất, trung bình cộng, slice đã được sắp xếp
func Slices() {
	var length int
	var numbers []float64
	var sum float64 = 0
	var min float64 = 0
	var max float64 = 0
	var avg float64 = 0

	for length <= 0 {
		fmt.Println("Input length of numbers: ")
		fmt.Scan(&length)
	}

	for i := 0; i < length; i++ {
		var number float64
		fmt.Printf("Input number %d: ", i+1)
		fmt.Scan(&number)
		numbers = append(numbers, number)
	}

	for i := 0; i < len(numbers); i++ {
		var number = numbers[i]
		sum = sum + number
		min = math.Min(min, float64(number))
		max = math.Max(max, float64(number))
		avg = sum / float64((len(numbers)))
	}

	sort.Float64s(numbers)

	fmt.Printf("Sum: %v, Min: %v, Max: %v, Avg: %v, Sorted: %v\n", sum, min, max, avg, numbers)
}
