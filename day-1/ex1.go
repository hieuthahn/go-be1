package main

import "fmt"

//  Viết hàm nhập 2 cạnh của hình chữ nhật, in ra diện tích, chu vi
func Rectangle() {
	var a int
	var b int

	fmt.Println("Input a: ")
	fmt.Scan(&a)

	fmt.Println("Input b: ")
	fmt.Scan(&b)

	var perimeter = rectanglePerimeter(a, b)
	var acreage = rectangleAcreage(a, b)

	fmt.Printf("Rectangle a: %d, b: %d, perimeter: %d, acreage: %d\n", a, b, perimeter, acreage)
}

func rectanglePerimeter(a, b int) int {
	return (a + b) * 2
}

func rectangleAcreage(a, b int) int {
	return a * b
}
