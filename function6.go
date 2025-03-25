// Function that returns another function
package main
import "fmt"
func multiplier(factor int) func(int) int {
    return func(number int) int {
        return number * factor
    }
}

func main() {
    double := multiplier(2)
    triple := multiplier(3)

    fmt.Println("Double 5:", double(5))
    fmt.Println("Triple 5:", triple(5))
}
