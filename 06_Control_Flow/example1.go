package main

import "fmt"

func main(){

	/*Write a program that prints out all the numbers between 1 and 100 that are
	evenly divisible by 3 (i.e., 3, 6, 9, etc.).*/
	for i:= 1; i < 100; i++ {
		if i % 3 == 0{
			fmt.Printf("%v \n", i)
		}
	}
}
