// Anonymous function assigned to a variable
package main
import "fmt"
func main() {
    greet := func(name string) {
        fmt.Println("Hello,", name)
    }

    greet("sakshi")
    greet("komal")
}
