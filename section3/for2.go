package main

import "fmt"

func main() {
	//예제 1
	sum1 := 0

	for i := 0; i <= 100; i++ {
		sum1 += i //sum = sum + 1

	}

	fmt.Println("ex1 : ", sum1)

	//예제2
	sum2, i := 0, 0

	for i <= 100 {
		sum2 += i
		i++
		// j := i++ <- 에러발생. Go 에서 수치연산은 반환 값이 없음.
	}
	fmt.Println(sum2)

	//예제3
	sum3, i := 0, 0

	for {
		if i > 100 {
			break
		}

		sum3 += i
		i++
	}

	//예제4
	for i, j := 0, 0; i <= 10 && j < 20; i, j = i+1, j+10 {
		fmt.Println("ex4 : ", i, j)
	}

	//에러 발생
	/*
		for i, j := 0, 0; i <= 10 && j < 20; i++, j += 10 { <- go 에서 수치연산은 반환 값이 없기 때문에 syntax에러 발생
			fmt.Println("ex4 : ", i, j)
		}
	*/
}
