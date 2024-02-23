package builder

type manualBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func (mb *manualBuilder) setWindowType(values ...string) {
	mb.windowType = values[0]
}

func (mb *manualBuilder) setDoorType(values ...string) {
	mb.doorType = values[0]
}

func (mb *manualBuilder) setFloor(values ...int) {
	mb.floor = values[0]
}

func (mb *manualBuilder) build() House {
	return House{
		windowType: mb.windowType,
		doorType:   mb.doorType,
		floor:      mb.floor,
	}
}
