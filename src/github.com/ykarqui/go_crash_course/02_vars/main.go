package main

import "fmt"

func main() {

	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// using var
	var age int32 = 25

	const isCool = true
	// isCool = false 	--> error

	// shorthand
	name := "Yamil"
	size := 1.3

	name, email := "Karqui", "ykarq@live.com"

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T", size)
}
