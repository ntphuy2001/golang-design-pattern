package builder

type iglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func (ib *iglooBuilder) setWindowType(values ...string) {
	ib.windowType = "Snow Window"
}

func (ib *iglooBuilder) setDoorType(values ...string) {
	ib.doorType = "Snow Door"
}

func (ib *iglooBuilder) setFloor(values ...int) {
	ib.floor = 1
}

func (ib *iglooBuilder) build() House {
	return House{
		windowType: ib.windowType,
		doorType:   ib.doorType,
		floor:      ib.floor,
	}
}
