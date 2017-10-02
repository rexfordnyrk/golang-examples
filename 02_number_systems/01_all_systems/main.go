package main

import "fmt"

func main(){
	x := 20
	//basically printing the number 20 in binary, decimal, octal and hexadecimal respectively
	fmt.Printf("%b %d %o %x \n", x,x,x,x)

	//basically printing the number 20 in binary, decimal, octal and hexadecimal respectively
	// but the # is introduced just to add the leading Zeros
	fmt.Printf("%#b %#d %#o %#x \n", x,x,x,x)
}
