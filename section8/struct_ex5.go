//구조체 심화(4)
/*
	Encapsulation -> Packages
	Inheritance -> Composition
	Polymorphism -> Interfaces
	Abstraction -> Embedding
*/
package main

import "fmt"

type Employee struct {
	name   string
	salary float64
	bonus  float64
}

func (e Employee) Calculate() float64 {
	return e.salary + e.bonus
}

func (e Executive) Calculate() float64 {
	return e.salary + e.bonus + e.specialBonus
}

type Executive struct {
	Employee     //임원도 직원이다. is a 관계!
	specialBonus float64
}

func main() {
	//구조체 임베디드 메소드 오버라이딩 패턴
	//예제1
	//직원
	ep1 := Employee{"kim", 2000000, 150000}
	ep2 := Employee{"park", 1500000, 200000}

	//임원
	ex := Executive{
		Employee{"lee", 5000000, 1000000},
		1000000,
	}

	fmt.Println("ex1 : ", int(ep1.Calculate()))
	fmt.Println("ex1 : ", int(ep2.Calculate()))

	//Employee 구조체 통해서 메소드 호출
	fmt.Println("ex1 : ", int(ex.Calculate())) //오버라이딩
	fmt.Println("ex1 : ", int(ex.Employee.Calculate()+ex.specialBonus))
}
