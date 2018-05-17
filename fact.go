package main

import "fmt"

func Factorial(number int) int {

    if number == 1 {

        return number
    }

    return number * Factorial(number-1)
}

func main() {

    answer := Factorial(9)
    fmt.Printf("Recursive: %d\n", answer)
} 
