// Golang program to illustrate the 
// concept of infinite recursion 
package main 

import ( 
	"fmt"
) 

// infinite recursion function 
func print_hello() { 
	
	// printing infinite times 
	fmt.Println("sakshi..") 
	print_hello() 
} 

// main function 
func main() { 
	
	// call to infinite 
	// recursive function 
	print_hello() 
} 
