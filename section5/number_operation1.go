package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//숫자 연산(산술, 비교)
	//타입이 같아야 가능
	//다른 타입끼리 반드시 형 변환 후 연산
	//형 변환 없을 경우 예외 발생
	// +, - *, /ㅡ >>, << ,& , ^
	/*
	   uint8       the set of all unsigned  8-bit integers (0 to 255)
	   uint16      the set of all unsigned 16-bit integers (0 to 65535)
	   uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
	   uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

	   int8        the set of all signed  8-bit integers (-128 to 127)
	   int16       the set of all signed 16-bit integers (-32768 to 32767)
	   int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
	   int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

	   float32     the set of all IEEE-754 32-bit floating-point numbers
	   float64     the set of all IEEE-754 64-bit floating-point numbers

	   complex64   the set of all complex numbers with float32 real and imaginary parts
	   complex128  the set of all complex numbers with float64 real and imaginary parts

	   byte        alias for uint8
	   rune        alias for int32
	   The value of an n-bit integer is n bits wide and represented using two's complement arithmetic.

	   There is also a set of predeclared numeric types with implementation-specific sizes:

	   uint     either 32 or 64 bits
	   int      same size as uint
	   uintptr  an unsigned integer large enough to store the uninterpreted bits of a pointer value

	*/
	var n1 uint8 = math.MaxUint8
	var n2 uint16 = math.MaxUint16
	var n3 uint32 = math.MaxUint32
	var n4 uint64 = math.MaxUint64

	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)

	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)

	n8 := uint(8)
	n9 := 8
	fmt.Println(n8)
	fmt.Println(reflect.TypeOf(n9))
}
