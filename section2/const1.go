package main

import "fmt"

func main() {
	const a string = "Test1"
	const b = "Test2"
	const c int32 = 10 * 10
	const d string = "Test3"
	//const a = getHeight() <- 상수는 함수를 값으로 받을 수 없음.
	const e = 35.6
	const f = false

	fmt.Println("a : ", a, "b : ", b, "c : ", c, "d: ", d, "e : ", e, "f : ", f)
}
