package main

import "fmt"

func main() {
    array1 := []int{11, 20, 30, 40, 59}
    array2 := []int{20, 40, 30, 6, 59,123,32}

    fmt.Println("Common Elements:")
    
    for i := 0; i < len(array1); i++ {        // Outer loop iterating over array1
        for j := 0; j < len(array2); j++ {    // Inner loop iterating over array2
            if array1[i] == array2[j] {
                fmt.Println(array1[i])        // Print common element
            }
        }
    }
}