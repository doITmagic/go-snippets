package main

import (
	"fmt"
)

// Pistoler is an interface, defines all methods a gun should have
type Pistoler interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

/***************************************************************************************/

// Pistol is a struct that implements IGun interface
type Pistol struct {
	name  string
	power int
}

// setName sets the name of the gun
func (g *Pistol) setName(name string) {
	g.name = name
}

// setPower sets the power of the gun
func (g *Pistol) setPower(power int) {
	g.power = power
}

// getName returns the name of the gun
func (g *Pistol) getName() string {
	return g.name
}

// getPower returns the power of the gun
func (g *Pistol) getPower() int {
	return g.power
}

/***************************************************************************************/

type Beretta struct {
	Pistol
}

func NewBeretta() Pistoler {
	return &Beretta{
		Pistol: Pistol{
			name:  "Beretta",
			power: 20,
		},
	}
}

/***************************************************************************************/

type Colt struct {
	Pistol
}

func NewColt() Pistoler {
	return &Colt{
		Pistol: Pistol{
			name:  "Colt",
			power: 40,
		},
	}
}

/***************************************************************************************/

func getGun(gunType string) (Pistoler, error) {
	if gunType == "beretta" {
		return NewBeretta(), nil
	}
	if gunType == "colt" {
		return NewColt(), nil
	}
	return nil, fmt.Errorf("wrong gun type passed")
}

func main() {
	beretta, _ := getGun("beretta")
	fmt.Printf("Name: %s\n", beretta.getName())
	fmt.Printf("Power: %d\n", beretta.getPower())

	colt, _ := getGun("colt")
	fmt.Printf("Name: %s\n", colt.getName())
	fmt.Printf("Power: %d\n", colt.getPower())
}
