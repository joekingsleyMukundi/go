package main

import "fmt"

const (
	SmallLift = iota
	StandardLift
	LargLift
)

type Lift int

type LiftPicker interface {
	PickLift() Lift
}

type Motrcycle string
type Car string
type Truck string

func (m Motrcycle) String() string {
	return fmt.Sprintf("Motorcycle: %v", string(m))
}
func (c Car) String() string {
	return fmt.Sprintf("Motorcycle: %v", string(c))
}
func (t Truck) String() string {
	return fmt.Sprintf("Motorcycle: %v", string(t))
}

func (m Motrcycle) PickLift() Lift {
	return SmallLift
}
func (c Car) PickLift() Lift {
	return StandardLift
}
func (t Truck) PickLift() Lift {
	return LargLift
}

func sendToLift(p LiftPicker) {
	switch p.PickLift() {
	case SmallLift:
		fmt.Printf("send %v to small lift\n", p)
	case StandardLift:
		fmt.Printf("send %v to standard lift\n", p)
	case LargLift:
		fmt.Printf("send %v to large lift\n", p)
	}

}
func main() {
	car := Car("Sporty")
	truck := Truck("MountainCrushers")
	mortorcycle := Motrcycle("Croozer")

	sendToLift(car)
	sendToLift(truck)
	sendToLift(mortorcycle)
}
