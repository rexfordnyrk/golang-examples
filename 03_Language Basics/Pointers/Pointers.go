package main

import "fmt"

func main()  {
	treasure := "deep in memory"
	fmt.Printf("%v \t %T \n",treasure, treasure)

	fmt.Printf("%v \t %T \n",&treasure, &treasure)

	var tr_location *string = &treasure
	*tr_location = "treasure found"

	fmt.Printf("%v %v \t %v %v %T \n", treasure, tr_location, &tr_location, *tr_location ,*tr_location)
}