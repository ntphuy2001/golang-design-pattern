package builder

type normalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func (nb *normalBuilder) setWindowType(values ...string) {
	nb.windowType = "Wooden Window"
}

func (nb *normalBuilder) setDoorType(values ...string) {
	nb.doorType = "Wooden Door"
}

func (nb *normalBuilder) setFloor(values ...int) {
	nb.floor = 2
}

func (nb *normalBuilder) build() House {
	return House{
		windowType: nb.windowType,
		doorType:   nb.doorType,
		floor:      nb.floor,
	}
}
