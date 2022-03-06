package main

import "fmt"

func main() {
	//맵(MAp)
	// 맵 조회 할 경우에 주의 할 점

	//예제1
	map1 := map[string]int{ // int : 0, string : "", float : 0.0
		"apple":  15,
		"banana": 115,
		"oragne": 1115,
		"lemon":  0,
	}

	value1 := map1["lemon"]
	value2 := map1["kiwi"]     //기본 값 0 return
	value3, ok := map1["kiwi"] //키가 있는지 없는지에 대해 true, false return 해줌 두 번째 값에

	fmt.Println("ex1 : ", value1)
	fmt.Println("ex1 : ", value2)
	fmt.Println("ex1 : ", value3, ok)

	//예제 2

	if value, ok := map1["kiwi"]; ok {
		fmt.Println("ex2 : ", value)
	} else {
		fmt.Println("ex2 : kiwi is not exists!")
	}

	if value, ok := map1["banana"]; ok {
		fmt.Println("ex2 : ", value)
	} else {
		fmt.Println("ex2 : banana is not exists!")
	}

	if _, ok := map1["kiwi"]; !ok {
		fmt.Println("ex2 : kiwi is not exists!")
	}
}
