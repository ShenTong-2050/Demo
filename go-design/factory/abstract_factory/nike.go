package abstract_factory

type Nike struct {}

type NikeShoes struct {
	Shoe
}

type NikeShirt struct {
	Shirt
}

func (n *Nike) MakeShoes() MyShoes {
	return &NikeShoes{
		Shoe{
			logo: "nike",
			size:19,
		},
	}
}

func (n *Nike) MakeShirt() MyShirt {
	return &NikeShirt{
		Shirt{
			logo: "nike",
			size: 20,
		},
	}
}
