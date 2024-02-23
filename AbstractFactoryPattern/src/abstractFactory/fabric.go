package abstractFactory

type fabric struct {
	fabricType string
}

func (f *fabric) setFabric(fabricType string) {
	f.fabricType = fabricType
}

func (f *fabric) GetFabric() string {
	return f.fabricType
}
