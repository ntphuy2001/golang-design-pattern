package builder

type Director struct {
	builder IBuilder
}

func NewDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) SetBuilder(b IBuilder) {
	d.builder = b
}

func (d *Director) BuildHouse() House {
	if mb, ok := d.builder.(*manualBuilder); ok {
		mb.setWindowType("Manual Window")
		mb.setDoorType("Manual Door")
		mb.setFloor(10)
	} else {
		d.builder.setFloor()
		d.builder.setDoorType()
		d.builder.setWindowType()
	}

	return d.builder.build()
}
