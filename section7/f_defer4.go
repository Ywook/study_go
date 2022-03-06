//함수 Defer(4)

package main

import "fmt"

func a() {
	defer end(start("b")) //중첩 함수 주의!! 그리고 웬만하면 defer문에 쓰지 않는게 좋음
	fmt.Println("in a")
}

func start(msg string) string {
	fmt.Println("start : ", msg)
	return msg
}

func end(t string) {
	fmt.Println("end : ", t)
}

func main() {

	a()
}
