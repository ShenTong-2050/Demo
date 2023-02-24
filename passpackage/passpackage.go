package passpackage

import (
	"demo/passpackage/private"
	"fmt"
	"unsafe"
)

// GetPrivatePkg 获取私有包内数据
func GetPrivatePkg() {
	// 得到 private 包内的 私有 对象 data
	var d = private.NewData()
	// 包外 访问 可访问数据 Y
	d.Y = 200
	// 利用 指针 访问并修改 private 包内 私有变量
	p := (*struct{x int})(unsafe.Pointer(d))
	p.x = 100
	fmt.Printf("%+v\n",*d)
}
