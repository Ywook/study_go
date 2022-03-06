package main

import "fmt"

func multiply(x, y int) (r int) {
	r = x * y
	return r
}

func sum(x, y int) (r int) {
	r = x + y
	return r
}

func main() {
	//함수 고급
	//함수를 변수에 할당
	f := []func(int, int) int{multiply, sum}

	a := f[0](10, 10)
	b := f[1](10, 10)

	fmt.Println(a)
	fmt.Println(b)

	var f1 func(int, int) int = multiply
	f2 := sum

	fmt.Println("ex2 : ", f1(10, 10))
	fmt.Println("ex2 2 ", f2(10, 10))

	//예제3 (맴에 할당)
	m := map[string]func(int, int) int{
		"mul_func": multiply,
		"sum_func": sum,
	}

	fmt.Println("ex3 : ", m["mul_func"](10, 10))
	fmt.Println("ex3 : ", m["sum_func"](10, 10))
}
