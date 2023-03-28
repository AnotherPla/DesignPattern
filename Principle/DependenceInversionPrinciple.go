package principle

import "fmt"

// --------------抽象层---------------
type Car interface {
	Run()
}
type Driver interface {
	Drive(car Car)
}

// --------------实现层---------------
type Tom struct{}

func (tom *Tom) Drive(car Car) {
	fmt.Println("tom is driving")
	car.Run()
}

type Jerry struct{}

func (jerry *Jerry) Drive(car Car) {
	fmt.Println("jerry is driving")
	car.Run()
}

type BYD struct{}

func (byd *BYD) Run() {
	fmt.Print("BYD is running\n")
}

type Benz struct{}

func (benz *Benz) Run() {
	fmt.Println("Benz is running")
}

// --------------逻辑层---------------
func DIP() {
	var benz, byd Car
	benz = new(Benz)
	byd = new(BYD)

	var tom Driver
	tom = new(Tom)
	tom.Drive(byd)

	var jerry Driver
	jerry = new(Jerry)
	jerry.Drive(benz)
}
