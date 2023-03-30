package main

import (
	"errors"
	"fmt"
	"github.com/jessevdk/go-flags"
	"strconv"
	"strings"
)

type ChildCommand struct {
	// Opt 算数类型
	Opt     string		`long:"op"`
	// Args 连续执行的参数
	Args    []string
	// Result 计算结果
	Result  int64
}

// Execute 执行子命令必须实现该方法
func (c *ChildCommand) Execute(args []string) error {
	if c.Opt != "+" && c.Opt != "-" && c.Opt != "x" && c.Opt != "/" {
		return errors.New("invalid operator")
	}
	for _,arg := range args {
		// 将参数 arg 转为 10进制 的 64位 整形
		num,err := strconv.ParseInt(arg,10,64)
		if err != nil {
			return errors.New("Invalid param " + arg)
		}
		c.Result = Calculator(c,num)
	}
	c.Args = args
	return nil
}

// Calculator 执行计算操作
func Calculator(c *ChildCommand,num int64) (int64) {
	if c.Result == 0 {
		return num
	}
	switch c.Opt {
	case "+":
		c.Result+=num
		break
	case "-":
		c.Result-=num
		break
	case "x":
		c.Result*=num
		break
	case "/":
		if num == 0 {
			return 0
		} else {
			c.Result/=num
		}
		break
	default:
		panic("something error...")
	}
	return c.Result
}

// ChildOption 定义子命令 结构体
type ChildOption struct {
	// long 传递的参数 command 子命令名称
	Calc ChildCommand `command:"calc"`
}

func main() {
	var opt ChildOption
	// 解析命令行参数
	flags.Parse(&opt)
	fmt.Printf("%s result of is %d\n",strings.Join(opt.Calc.Args,opt.Calc.Opt),opt.Calc.Result)
}