package decimal

import (
	"fmt"
)

// DecimalMain golang int 与 float 计算、加减乘除
func DecimalMain() {
	mathBigPackage()
}

func mathBigPackage() {
	var a float64 = 5
	var b float64 = 2

	fmt.Println(a/b)
}
