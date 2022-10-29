package main

import "fmt"

func main() {
	// Define map
	//emails := make(map[string]string)
	//
	// assign key values
	//emails["Alice"] = "alice@gmail.com"
	//emails["Bob"] = "bob@gmail.com"
	//emails["Carl"] = "cc@gmail.com"

	// Declare map and add K-V
	emails := map[string]string{"Bob": "bob@gmail.com", "Carl": "cc@gmail.com"}

	fmt.Println(emails)
	fmt.Println(emails["Bob"])

	// delete from map
	delete(emails, "Bob")
	fmt.Println(emails)

}
