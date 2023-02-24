package interfaces

import (
	"fmt"
	"reflect"
)

type U struct {
	S
}

type S struct {}

func (U) ValU() {
	fmt.Println("this is uValue")
}
func (*U) PtrU() {
	fmt.Println("this is uPtr")
}
func (S) ValS() {
	fmt.Println("this is eVal")
}
func (*S) PtrS() {
	fmt.Println("this is ePtr")
}

func methodSet(f interface{}) {
	// 返回接口中保存值的类型
	t := reflect.TypeOf(f)
	// 遍历获取所有的方法集
	for i,n:=0,t.NumMethod(); i<n;i++ {
		m := t.Method(i)
		fmt.Println(m)
	}
}

func MethodMain() {
	var f U
	methodSet(f)	// 显示 user 方法集
	fmt.Println("------------")
	methodSet(&f)	// 显示 *user 方法集
}