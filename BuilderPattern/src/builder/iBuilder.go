package builder

type IBuilder interface {
	setWindowType(values ...string)
	setDoorType(values ...string)
	setFloor(values ...int)
	build() House
}

func GetBuilder(builderType string) IBuilder {
	switch builderType {
	case "normal":
		return &normalBuilder{}
	case "igloo":
		return &iglooBuilder{}
	case "manual":
		return &manualBuilder{}
	default:
		return nil
	}
}
