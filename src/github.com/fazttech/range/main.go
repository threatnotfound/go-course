package main

import (
	"fmt"
)

func main() {
	ids := []int{50, 100, 77, 10}

	// loop
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// without using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// sum ids
	sum := 0
	for _, id := range ids {
		sum += id
	}

	println(sum)

	// range with maps
	emails := map[string]string{"fazt": "fazt@faztweb.com", "jesus": "jesus@gmail.com", "ryan": "ryan@gmail.com"}

	// loop the keys and values
	for k, v := range emails {
		fmt.Printf("%s: \t\t%s\n", k, v)
	}

	// loop the keys
	for k := range emails {
		fmt.Println("Name:" + k)
	}

}
