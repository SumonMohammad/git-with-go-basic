package main

type PaymentGateway interface {
	Pay(amount float64) error
}
type Bksah struct {
}

func (b Bksah) Pay(amount float64) error {
	b.Pay(amount)
	return nil
}

type Nagad struct {
}

func (n Nagad) Pay(amount float64) error {
	n.Pay(amount)
	return nil
}

func PaymentProcess(g PaymentGateway, amount float64) {
	g.Pay(500)
}

func main() {
	bkash := Bksah{}
	nagad := Nagad{}

	PaymentProcess(bkash, 500)
	PaymentProcess(nagad, 800)
}
