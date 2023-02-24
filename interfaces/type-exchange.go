package interfaces

import (
	"fmt"
	"runtime"
)

type data int

type X int

type FuncString func() string

func (d data) String() string {
	return fmt.Sprintf("data:%d",d)
}

func (f FuncString) String() string {
	return f()
}

//ChangeMain 转换类型入口
func ChangeMain() {

	var d data = 5

	var x interface{} = d

	if n,ok := x.(fmt.Stringer); ok {		// 转换为更具体的接口类型
		fmt.Println(n)
	}

	if d2,ok := x.(data); ok {				// 转换为原始类型
		fmt.Println(d2)
	}

	fmt.Println("-----------")

	var n interface{} = func(x int) string {
		return fmt.Sprintf("d:%d",x)
	}

	switch v := n.(type) {		// 局部变量 v 是类型转换后的结果
	case nil:
		println("nil")
	case *int:
		println(*v)
	case func(int) string:
		println(v(100))
	case fmt.Stringer:
		fmt.Println(v)
	default:
		println("unknown")
	}

	fmt.Println("------------")

	var t = FuncString(func() string {
		return "hello world"
	})
	fmt.Println(t)

	fmt.Println("--------")

	fmt.Println(runtime.NumCPU())
}









