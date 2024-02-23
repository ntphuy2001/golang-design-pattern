package main

import (
	"AbstractFactoryPattern/src/abstractFactory"
	"fmt"
)

func main() {
	adidasFactory, _ := abstractFactory.GetSportsFactory("adidas")
	nikeFactory, _ := abstractFactory.GetSportsFactory("nike")

	adidasShoe := adidasFactory.MakeShoe()
	adidasShort := adidasFactory.MakeShort()

	nikeShoe := nikeFactory.MakeShoe()
	nikeShort := nikeFactory.MakeShort()

	printShoeDetails(adidasShoe)
	printShoeDetails(nikeShoe)
	printShirtDetails(adidasShort)
	printShirtDetails(nikeShort)

}

func printShoeDetails(s abstractFactory.IShoe) {
	fmt.Println("Shoe:")
	fmt.Printf("Logo: %s", s.GetLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.GetSize())
	fmt.Println()
}

func printShirtDetails(s abstractFactory.IShort) {
	fmt.Println("Short:")
	fmt.Printf("Logo: %s", s.GetLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.GetSize())
	fmt.Println()
}
