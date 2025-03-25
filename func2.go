package main

import "fmt"

func mult(a, b int) int{
	a=a*2
	return a * b
}
func main(){
	x:=5
	y:=10
	fmt.Printf(" before: x=%d, y =%d\n",x, y)
    result:= mult(x, y)
	fmt.Printf("multiplication :  %d\n",result)
	fmt.Printf("After: x=%d, y =%d\n", x,y)
}