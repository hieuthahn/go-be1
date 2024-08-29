package main

import "fmt"

func main() {
	var option int
	
	fmt.Println("Please choose exercise: ")
	fmt.Println("1. Rectangle")
	fmt.Println("2. Is Even String")
	fmt.Println("3. Slices")
	fmt.Println("4. Two Sum")
	fmt.Scan(&option)

	switch option {
		case 1:
			Rectangle()
		case 2:
			IsEvenString()
		case 3:
			Slices()
		case 4:
			TwoSum()
		default: 
			fmt.Println("Invalid option")
	}
}
