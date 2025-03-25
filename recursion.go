package main             //DIRECT RECURSION
import (
	"fmt"
)
func factorial(number int) int {
	if number == 0 || number == 1{
		return 1
	}
	if number < 0{
		fmt.Println("invalid number..")
		return -1
	}
	return number*factorial((number -1))
}
	func main(){
		answer1 := factorial(0)
		fmt.Println(answer1, "\n")

		answer2 := factorial(5)
		fmt.Println(answer2, "\n")

		answer3 := factorial(-1)
		fmt.Println(answer3, "\n")

		answer4 := factorial(10)
		fmt.Println(answer4, "\n")
	
}
