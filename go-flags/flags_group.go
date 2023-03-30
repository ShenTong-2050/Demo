package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"log"
	"os"
)

type GroupOption struct {
	BasicGroup		BasicOption		`description:"basic data type group" group:"basic"`
	SliceGroup 		SliceOption		`description:"slice data type group" group:"slice"`
}

type BasicOption struct {
	IntFlag		int			`short:"i" required:"true" long:"intFlag" description:"intFlag description" default:"1"`
	StrFlag   	string		`short:"s" required:"true" long:"strFlag" description:"strFlag description" default:"a"`
	FloatFlag	float64		`short:"f" required:"true" long:"floatFlag" description:"floatFlag description" default:"3.12"`
}

type SliceOption struct {
	SliceInt	[]int		`short:"m" required:"true" long:"sliceInt" description:"sliceInt description" default:"2"`
	SliceStr	[]string	`short:"n" required:"true" long:"sliceStr" description:"sliceStr description" default:"b"`
	SliceFloat	[]float64	`short:"o" required:"true" long:"sliceFloat" description:"sliceFloat description" default:"3.55"`
}

func main() {
	var opt GroupOption
	p := flags.NewParser(&opt,flags.Default)
	_,err := p.ParseArgs(os.Args[1:])
	if err != nil {
		log.Fatal("Parse error : ",err)
	}

	basicGroup := p.Command.Group.Find("basic")
	for _,basicOption := range basicGroup.Options() {
		fmt.Printf("name:%v value:%v\n",basicOption.LongNameWithNamespace(),basicOption.Value())
	}

	sliceGroup := p.Command.Group.Find("slice")
	for _,sliceOption := range sliceGroup.Options() {
		fmt.Printf("name:%v value:%v\n",sliceOption.LongNameWithNamespace(),sliceOption.Value())
	}
}