package main

import (
	"BuilderPattern/src/builder"
	"fmt"
)

func main() {
	normalBuilder := builder.GetBuilder("normal")
	iglooBuilder := builder.GetBuilder("igloo")
	manualBuilder := builder.GetBuilder("manual")

	director := builder.NewDirector(normalBuilder)
	normalHouse := director.BuildHouse()
	fmt.Println(normalHouse)

	director = builder.NewDirector(iglooBuilder)
	iglooHouse := director.BuildHouse()
	fmt.Println(iglooHouse)

	director = builder.NewDirector(manualBuilder)
	manualHouse := director.BuildHouse()
	fmt.Println(manualHouse)
}
