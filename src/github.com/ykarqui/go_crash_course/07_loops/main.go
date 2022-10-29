package main

import "fmt"

func main() {
	i := 1
	// long
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// short
	for j := 1; j <= 10; j++ {
		fmt.Println(j)
	}

	// fizzbuzz
	for i := 1; i < 100; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

}
