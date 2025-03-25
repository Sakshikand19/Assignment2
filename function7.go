package main

import "fmt"

// Function that accepts a variable number of arguments
func sum(numbers ...int) int {
    total := 0
    for _, number := range numbers {
        total += number
    }
    return total
}

func main() {
    fmt.Println("Sum of 1, 2, 3:", sum(1, 2, 3))
    fmt.Println("Sum of 10, 20, 30, 40:", sum(10, 20, 30, 40))
}
