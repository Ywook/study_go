package main

import "fmt"

func main() {
	//Swith 뒤 Expression 생략 가능
	//case 뒤 표현식 Expression 사용 가능
	//자동 break 때문에 fallthrouth 존재
	//Type 분기 -> 값이 아닌 변수 Type 분기 가능

	a := -7
	switch {
	case a < 0:
		fmt.Println(a, "는 음수")
	case a == 0:
		fmt.Println(a, "는 0")
	case a > 0:
		fmt.Println(a, "는 양수")
	}

	switch b := 27; {
	case b < 0:
		fmt.Println(b, "는 음수")
	case b == 0:
		fmt.Println(b, "는 0")
	case b > 0:
		fmt.Println(b, "는 양수")
	}

	//예제 3
	switch c := "go"; c {
	case "go":
		fmt.Println("Go!")
	case "java":
		fmt.Println("Java!")
	default:
		fmt.Println("X")
	}

	//예제 4
	switch c := "go"; c + "lang" {
	case "golang":
		fmt.Println("golang")
	case "java":
		fmt.Println("Java!")

	default:
		fmt.Println("X")
	}

	//예제 5
	switch i, j := 20, 30; {
	case i < j:
		fmt.Println("i < j")
	case i == j:
		fmt.Println("i==j")
	}
}
