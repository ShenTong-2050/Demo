package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
)

type Option struct {
	Verbose 			[]bool				`short:"v" long:"verbose" description:"show verbose debug message"`
	FlagInt 			int					`short:"i" long:"intFlag" default:"2" description:"int flag description"`
	SliceInt			[]int				`long:"sliceInt" description:"slice int flag"`
	FlagFloat   		float64				`long:"FlagFloat" description:"flag float description"`
	SliceFloat			[]float64			`long:"sliceFloat" description:"slice float description"`
	FlagString  		string				`long:"string" description:"flag string description"`
	SliceString 		[]string 			`long:"sliceString" description:"slice string description"`
	SlicePtrString   	[]*string			`long:"ptrString" description:"slice ptr string description"`
	Call                func(val string)	`long:"call" description:"function description"`
	IntMap				map[string]int		`long:"intMap" description:"int map description"`
}



func main() {
	var option Option

	// 命令行函数类型
	option.Call = func(str string) {
		fmt.Println("in callback",str)
	}

	flags.Parse(&option)

	fmt.Println(option.Verbose)
	fmt.Printf("int flag %v\n",option.FlagInt)
	fmt.Printf("slice int %v\n",option.SliceInt)
	fmt.Printf("flag float %v\n",option.FlagFloat)
	fmt.Printf("slice float %v\n",option.SliceFloat)
	fmt.Printf("flag string %v\n",option.FlagString)
	fmt.Printf("slice string %v\n",option.SliceString)

	fmt.Println("slice ptr string ")
	for i:=0; i<len(option.SlicePtrString); i++ {
		fmt.Printf("%d slice ptr string %v\n",i,*option.SlicePtrString[i])
	}

	fmt.Printf("int map %v\n",option.IntMap)
}
