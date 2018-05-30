package main

import "fmt"

func main() {
	var arr [5]int
	var n int
	var sum int
	sum = 0

	fmt.Println("Enter The Size Of The Array")
	fmt.Scanln(&n)
	fmt.Println("Enter The Elements Of the Array")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	for i := 0; i < n; i++ {
		sum = sum + arr[i]

	}
	fmt.Println("Sum Of Array Elements", sum)

}
