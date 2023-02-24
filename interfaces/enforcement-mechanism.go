package interfaces

type Ner interface {
	a()
	b(int)
	c(string)	string
}

type N struct {}

func (N) a() {}

func (*N) b(int) {}

func (*N) c(s string) string {
	return s
}

func EnforcementMain() {

	var n N

	var ner Ner = &n

	println(ner.c("看看 interface 的执行机制!"))
}


