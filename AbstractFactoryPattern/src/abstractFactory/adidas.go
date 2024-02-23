package abstractFactory

type adidas struct{}

func (a *adidas) MakeShoe() IShoe {
	return &adidasShoe{
		shoe: shoe{
			logo: "Adidas",
			size: 14,
		},
	}
}

func (a *adidas) MakeShort() IShort {
	return &adidasShort{
		short: short{
			logo: "Adidas",
			size: 14,
		},
	}
}
