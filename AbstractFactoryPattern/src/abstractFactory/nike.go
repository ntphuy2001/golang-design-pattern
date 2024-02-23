package abstractFactory

type nike struct{}

func (n *nike) MakeShoe() IShoe {
	return &nikeShoe{
		shoe: shoe{
			logo: "Nike",
			size: 12,
		},
	}
}

func (n *nike) MakeShort() IShort {
	return &nikeShort{
		short: short{
			logo: "Nike",
			size: 12,
		},
	}
}
