package main

import (
	cor "ChainOfResponsibilityPattern/src/chainofresponsibility"
)

func main() {

	cashier := &cor.Cashier{}

	//Set next for medical department
	medical := &cor.Medical{}
	medical.SetNext(cashier)

	//Set next for doctor department
	doctor := &cor.Doctor{}
	doctor.SetNext(medical)

	//Set next for reception department
	reception := &cor.Reception{}
	reception.SetNext(doctor)

	patient := &cor.Patient{Name: "abc"}
	//Patient visiting
	reception.Execute(patient)

}
