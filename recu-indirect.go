package main
import (
	"fmt"
)
func Print_one(n int){
	if n>= 0{
		fmt.Println("In first function :",n)
		print_two(n-1)
	}
}
func print_two(n int ){
	if n>= 0{
		fmt.Println("in second function ",n)
		Print_one(n-1)
	}
}
func main(){
	Print_one(10)
	Print_one(-1)
}