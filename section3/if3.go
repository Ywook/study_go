package main

import "fmt"

func main() {
	i := 100

	// if - else if 예제 1
	if i >= 120 {
		fmt.Println("over 120")
	} else if i >= 100 && i < 120 {
		fmt.Println("over 100")
	} else if i < 100 && i >= 50 {
		fmt.Println("over 50")
	} else {
		fmt.Println(" <= 50 ")
	}

}
