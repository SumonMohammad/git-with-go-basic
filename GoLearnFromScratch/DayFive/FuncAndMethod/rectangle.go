package main

import "fmt"

type Rectangle struct {
	width  int64
	height int64
}

func (r Rectangle) Rec() int64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	if r.width == 0 || r.height == 0 {
		fmt.Println("Invalid height or width")
		return 0
	}
	return float64(2 * (r.width + r.height))
}

func main() {
	calcRec := Rectangle{width: 4, height: 6}
	calPerimeter := Rectangle{width: 0, height: 6}

	fmt.Println("ueuerueyr:", calcRec.Rec())

	fmt.Println("Periimeter:", calPerimeter.Perimeter())
}
