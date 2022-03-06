package main

import "fmt"

func main() {

	//예제 1
	a := 30 / 15
	switch a {
	case 2, 4, 6:
		fmt.Println("a -> ", a, "는 짝수")
	case 1, 3, 5:
		fmt.Println("a -> ", a, "는 홀수")
	}

	//예제 2
	switch e := "go"; e {
	case "java":
		fmt.Println("java")
	case "go":
		fmt.Println("go")
		fallthrough
	case "python":
		fmt.Println("python")
	case "ruby":
		fmt.Println("ruby")
	}
	fmt.Println("")
}
