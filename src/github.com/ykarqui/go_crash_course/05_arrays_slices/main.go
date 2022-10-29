package main

import "fmt"

func main() {
	// Arrays
	var fruitArr [2]string

	// Assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	// Declare and assign
	sportArr := [2]string{"Tennis", "Softball"}
	lengArr := []string{"go", "java", "js", "c++", "Scala"}

	fmt.Println(fruitArr[1])
	fmt.Println(sportArr)
	fmt.Println(len(sportArr))
	fmt.Println(lengArr[2:4])
}
