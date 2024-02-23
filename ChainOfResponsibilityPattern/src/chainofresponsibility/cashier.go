package chainofresponsibility

import "fmt"

type Cashier struct {
	next Department
}

func (c *Cashier) Execute(p *Patient) {
	if p.isPaid {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from patient patient")
	p.isPaid = true
}

func (c *Cashier) SetNext(next Department) {
	c.next = next
}
