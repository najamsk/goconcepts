package main

import "fmt"

type Color string
type Category string
type MaxSpeed int

type Vehicle interface {
	Drive() string
	Stop() string
}

type VehicleBuilder interface {
	TopSpeed(MaxSpeed) VehicleBuilder
	Paint(Color) VehicleBuilder
	Category(Category) VehicleBuilder
	Build() Vehicle
}

type vehicleBuilder struct {
	speed    MaxSpeed
	category Category
	color    Color
}

func (cb *vehicleBuilder) Category(c Category) VehicleBuilder {
	cb.category = c
	return cb
}
func (cb *vehicleBuilder) TopSpeed(ms MaxSpeed) VehicleBuilder {
	cb.speed = ms
	return cb
}

func (cb *vehicleBuilder) Paint(p Color) VehicleBuilder {
	cb.color = p
	return cb
}

func (cb *vehicleBuilder) Build() Vehicle {
	var vehicle Vehicle
	switch cb.category {
	case "jet":
		vehicle = &jet{
			speed: cb.speed,
			color: cb.color,
		}
	case "car":
		vehicle = &car{
			speed: cb.speed,
			color: cb.color,
		}
	}
	return vehicle
}

type car struct {
	speed MaxSpeed
	color Color
}

func (c *car) Drive() string {
	return "vroom"
}

func (c *car) Stop() string {
	return "breaking"
}

type jet struct {
	speed MaxSpeed
	color Color
}

func (j *jet) Drive() string {
	return "zoom"
}

func (j *jet) Stop() string {
	return "landing"
}

func New() VehicleBuilder {
	return &vehicleBuilder{}
}

func main() {
	vb := New()
	vb.Paint("Pink")
	vb.Category("car")
	vb.TopSpeed(200)
	girly := vb.Build()
	fmt.Printf("%T, %#v \n", girly, girly)

	vb = New()
	vb.Paint("Black")
	vb.Category("jet")
	vb.TopSpeed(20000)
	panther := vb.Build()
	fmt.Printf("%T, %#v \n", panther, panther)

}
