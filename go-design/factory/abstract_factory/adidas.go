package abstract_factory

type Adidas struct {}

type AdidasShoe struct {
	Shoe
}

type AdidasShirt struct {
	Shirt
}

func (a *Adidas) MakeShoes() MyShoes {
	return &AdidasShoe{
		Shoe{
			logo: "adidas",
			size: 14,
		},
	}
}

func (a *Adidas) MakeShirt() MyShirt {
	return &AdidasShirt{
		Shirt{
			logo: "adidas",
			size: 17,
		},
	}
}




