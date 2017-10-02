package main

import "fmt"

func main(){

	// printing the first 100 numbers in the various number systems using loop

	for i:=0 ; i < 100; i++ {
		fmt.Printf("%#b \t %#d \t %#o \t %#X \n", i,i,i,i)
	}
}
