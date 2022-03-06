// 자료형 : 포인터(2)

package main

import (
	"fmt"
	_ "fmt"
)

func main() {
	//예제 1
	i := 7
	p := &i

	fmt.Println("ex1 : ", i, *p, &i, p)

	*p++

	fmt.Println("ex1 : ", i, *p, &i, p)
}
