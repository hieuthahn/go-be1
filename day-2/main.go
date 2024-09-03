// Viết hàm tạo ra 1 struct về 1 người (gồm tên, nghề nghiệp, năm sinh) , và struct có method tính tuổi,  method kiểm tra người ấy có hợp với nghề của mình không ( năm sinh chia hết cho số chữ trong tên)
// Viết hàm có input là 1 string, trả về map có key là ký tự, và value là số lần xuất hiện của ký tự đó ở trong string
// Viết hàm giải bài toán twosum : https://leetcode.com/problems/two-sum/ sử dụng map
// Cho 1 file a.txt , trả về slice các struct là người, với các thông tin lấy được từ file đó, tên cần được in hoa tất cả, nghề nghiệp cần được viết thường tất cả ( tham khảo đọc file ở https://zetcode.com/golang/readfile/)
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	option := -1

	for option != 0 {
		fmt.Println("Please choose exercise: ")
		fmt.Println("0. Exit")
		fmt.Println("1. Exercise 1")
		fmt.Println("2. Exercise 2")
		fmt.Println("3. Exercise 3")
		fmt.Println("4. Exercise 4")
		fmt.Scan(&option)

		switch option {
			case 0:
				return
			case 1:
				Exercise1()
				option = -1
			case 2:
				Exercise2()
				option = -1
			case 3:
				Exercise3()
				option = -1
			case 4: 
				Exercise4()
				option = -1
			default: 
				fmt.Println("Invalid option")
		}
	}
}

// Exercise 1
// Start
type Person struct {
	Name string
	Job string
	Year int
}

func getAge(person Person) int {
	return time.Now().Year() - person.Year
}

func checkingSuitableJob(person Person) bool {
	if (person.Year == 0 || person.Name == "") {
		return false
	}

	return person.Year % len(person.Name) == 0
}

func Exercise1() {
	var person Person

	fmt.Println("Input name: ")
	fmt.Scan(&person.Name)

	fmt.Println("Input job: ")
	fmt.Scan(&person.Job)

	fmt.Println("Input year: ")
	fmt.Scan(&person.Year)

	fmt.Printf("Name: %s, Job: %s, Age: %d, Suitable: %v\n", person.Name, person.Job, getAge(person), checkingSuitableJob(person))
	return 
}
// End Exercise 1

// Exercise 2
// Start
func Exercise2() {
	var input string
	var mapValues map[byte]int
	mapValues = make(map[byte]int)

	fmt.Println("Input string: ")
	fmt.Scan(&input)

	for i := 0; i < len(input); i++ {
		char := input[i]
		if (mapValues[char] == 0) {
			mapValues[char] = 0
		}
		mapValues[char] = mapValues[char] + 1
	}
	
	for char, count := range mapValues {
		fmt.Printf("Char: %c, Count: %d\n", char, count)
	}
	return
}
// End Exercise 2

// Exercise 3
// Start
func Exercise3() {
	var inputLength int
	var input []int
	var target int
	var mapValues = make(map[int]int)

	fmt.Println("Input length: ")
	fmt.Scan(&inputLength)

	for i := range inputLength {
		num := 0
		fmt.Printf("Enter input index %d:\n", i)
		fmt.Scan(&num)
		input = append(input, num)
	}

	fmt.Println("Input: ", input)
	fmt.Println("Enter Target: ")
	fmt.Scan(&target)

	for i, num := range input {
		var minus = target - num
		if idx, found := mapValues[minus]; found {
			fmt.Printf("Input: %v, Target: %d, Result: %v\n", input, target, []int{idx, i})
			return
		}
		mapValues[num] = i
		fmt.Println(mapValues, mapValues[minus], minus)
	}
	fmt.Printf("Input: %v, Target: %d, Result: %v\n", input, target, []int{})
	return
}
// End exercise 3


// Exercise 4
// Start
func Exercise4() {
		personSlices := []Person{}
		f, err := os.Open("a.txt")

		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()

		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
			personArray := strings.Split(scanner.Text(), ", ")
			year, err := strconv.Atoi(personArray[1])
			
			if (err != nil) {
				log.Fatal(err)
			}

			personSlices = append(personSlices, Person{
				Name: strings.ToUpper(personArray[0]),
				Year: year,
				Job: strings.ToLower(personArray[2]),
			})
		}

		if err := scanner.Err(); err != nil {
			log.Fatal((err))
		}

		fmt.Println("Result: ", personSlices)
		return
}
// End exercise 4
