package convey

func Add(a,b int) int {
	return a+b
}

func Sub(a,b int)int  {
	return a-b
}

//斐波那契数列
func Fib(a int) int {
	if a<2 {
		return a
	}else {
		return Fib(a-1)+Fib(a-2)
	}
}
