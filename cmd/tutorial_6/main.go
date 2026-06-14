package main

import (
	"fmt"
)

// this file is explaining about struct, and interface

type gundamType struct {
	model string
	pilot
	manufacturer string
}

type zakuType struct {
	model string
	pilot
	manufacturer string
}

func (g gundamType) useWeapon() string {
	return "Laser Beam"
}

func (z zakuType) useWeapon() string {
	return "Thermal Axe"
}

// interface is a contract, that is used to define the methods or properties that a struct must have
type mobileSuit interface {
	useWeapon() string
}

// and it can be useful for function parameter type
// the function can accept any struct that implements the interface
func useWeapon(robot mobileSuit) string {
	return robot.useWeapon()
}

type pilot struct {
	name   string
	age    uint8
	gender string
}

// method on struct
func (p pilot) sayHello(gundam string) {
	p.age++ // modify only the copy of the struct
	fmt.Printf("Hello, I am %v, I am %v years old, and I am the pilot of %v\n", p.name, p.age, gundam)
}

// the (p pilot) is called a receiver, that is referring to the pilot struct
// the receiver is used to make the function tied and have a copy of the fields of the struct, so
// the function can use the fields of the struct, if we manipulate the field inside the function
// it has nothing to do with the original struct

// pointer receiver
func (p *pilot) aging() {
	p.age++ // modify the original value
}

// here, the (p *pilot) is pointer receiver. This one is different, if a normal receiver only copy the struct
// so it can use the field, then the pointer receiver is directly tied to the struct fields and value
// For the example above, the age field that we are accessing in the aging function is the age field of the struct
// so the original value will be replaced by the new value that we modify inside the function

func main() {
	newGundam := gundamType{
		model: "RX-93-v2 Hi-v",
		pilot: pilot{
			name:   "Amuro Ray",
			age:    20,
			gender: "male",
		},
		manufacturer: "Anaheim Electronics",
	}
	newZaku := zakuType{
		model: "Zaku Mk II",
		pilot: pilot{
			name:   "Char Aznable",
			age:    20,
			gender: "male",
		},
		manufacturer: "Zeonic Company",
	}

	fmt.Println(newGundam.model, newGundam.pilot.name, newGundam.manufacturer)

	// you can also make a struct inside a variable, but it is not reusable
	newPilot := struct {
		name   string
		age    uint8
		gender string
	}{
		name:   "Kira Yamato",
		age:    20,
		gender: "male",
	}

	fmt.Println(newPilot.name, newPilot.age, newPilot.gender)
	newGundam.pilot.sayHello(newGundam.model) // there is also p.age++ in this function
	// but it is not modifying the original value, it is modifying the copy of the struct
	fmt.Println("Without pointer receiver:", newGundam.pilot.age)

	newGundam.pilot.aging()
	fmt.Println("With pointer receiver:", newGundam.pilot.age)

	fmt.Println(useWeapon(newGundam))
	fmt.Println(useWeapon(newZaku))

}
