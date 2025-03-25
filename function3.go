// Function with no parameters but returns a value
package main
import "fmt"
func getGreeting() string {
    return "Hello from Go!"
}
func main() {
   // greeting := getGreeting()
    fmt.Println(getGreeting())
}
