package main

import "fmt"

// Viết chương trình nhập 1 string, in ra true nếu độ dài chuỗi chia hết cho 2, và false nếu ngược lại
func IsEvenString() {
	var str string

	fmt.Println("Input string: ")
	fmt.Scan(&str)

	fmt.Printf("Is even string: %t",len(str) % 2 == 0 )
}
