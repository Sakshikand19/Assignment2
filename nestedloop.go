package main

import "fmt"

func main() {
  array := [][]int{{14, 32, 33}, {54, 15, 65}, {73, 28, 59}}
  largest := array[0][0]

  for i := 0; i < len(array); i++ {
    for j := 0; j < len(array[0]); j++ {
      if array[i][j] > largest {
        largest = array[i][j]
      }
    }
  }

  fmt.Println("Largest element:", largest)
}