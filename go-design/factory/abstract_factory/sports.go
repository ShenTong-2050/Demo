package abstract_factory

// MySports 抽象接口
type MySports interface {
	MakeShoes() MyShoes
	MakeShirt() MyShirt
}

func GetMySports(t string) MySports {
	switch t {
	case "nike":
		return &Nike{}
	case "adidas":
		return &Adidas{}
	}
	return nil
}

