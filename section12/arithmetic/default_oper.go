// 두 숫자의 사칙연산 계산 제공 패키지(1)
//X, Y 2개의 Integer 구조체
package arithmetic

type Numbers struct {
	X int
	Y int
}

//x,y 합을 계산해서 반환
func (o *Numbers) Plus() int {
	return o.X + o.Y
}

func (o *Numbers) Minus() int {
	return o.X - o.Y
}

func (o *Numbers) Multi() int {
	return o.X * o.Y
}

func (o *Numbers) Divide() int {
	return o.X / o.Y
}
