package abstractFactory

import "fmt"

type ISportsFactory interface {
	MakeShoe() IShoe
	MakeShort() IShort
}

func GetSportsFactory(brand string) (ISportsFactory, error) {
	switch brand {
	case "adidas":
		return &adidas{}, nil
	case "nike":
		return &nike{}, nil
	default:
		return nil, fmt.Errorf("Wrong brand type passed")
	}
}
