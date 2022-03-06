package main

import "fmt"

func main() {
	const (
		_ = iota
		A
		B
		C
	)

	const (
		_ = iota + 0.75*2
		DEFAULT
		SILVER
		GOLD
		PLATINUM
	)

	fmt.Println(DEFAULT, SILVER, GOLD, PLATINUM)
}
