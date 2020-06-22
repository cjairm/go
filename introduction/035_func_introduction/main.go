package main

import (
	"fmt"
	"strconv"
)

type personType struct {
	firstName    string
	lastName     string
	neighborhood string
}

type weaponsType struct {
	guns   int
	arrows int
	sticks int
}

type heroeType struct {
	person      personType
	weapons     weaponsType
	superSpeed  bool
	superVision bool
	superForce  bool
	superRich   bool
}

// type humanInterface interface {
// 	greet()
// }

// func (h heroeType) greet() string {
// 	fullName := h.person.firstName + " " + h.person.lastName
// 	return fmt.Sprint("Hello... I'm a super hero and my name is: ", fullName)
// }

func main() {
	sh1 := heroeType{
		person: personType{
			firstName:    "Bruce",
			lastName:     "Wayne",
			neighborhood: "Gotham",
		},
		weapons: weaponsType{
			guns:   4,
			sticks: 1,
		},
		superRich: true,
	}

	sh2 := heroeType{
		person: personType{
			firstName:    "Oliver",
			lastName:     "Queen",
			neighborhood: "Star City",
		},
		weapons: weaponsType{
			arrows: 100,
		},
		superRich: true,
	}

	sh3 := heroeType{
		person: personType{
			firstName:    "Kal",
			lastName:     "El",
			neighborhood: "Krypton",
		},
		superForce:  true,
		superSpeed:  true,
		superVision: true,
	}

	sh4 := heroeType{
		person: personType{
			firstName:    "Harry",
			lastName:     "Lampert",
			neighborhood: "Central City",
		},
		superSpeed: true,
	}

	sh5 := heroeType{
		person: personType{
			firstName:    "Chapulin",
			lastName:     "Colorado",
			neighborhood: "Mexico",
		},
		weapons: weaponsType{
			arrows: 1,
			guns:   1,
			sticks: 1,
		},
	}

	// shs := []heroeType{sh1, sh2, sh3, sh4, sh5}

	fmt.Println(sh1.greet())
	fmt.Println(sh2.greet())
	fmt.Println(sh3.greet())
	fmt.Println(sh4.greet())
	fmt.Println(sh5.greet())
	// fmt.Println(shs)
}

func (h heroeType) sumWeapons() int {
	heroWeapons := h.weapons
	var totalWeapons int
	totalWeapons += heroWeapons.arrows
	totalWeapons += heroWeapons.arrows
	totalWeapons += heroWeapons.sticks
	return totalWeapons
}

func (h heroeType) greet() (int, string) {
	totalWeapons := h.sumWeapons()
	greet := "Hello my name is " + h.person.firstName + " " + h.person.lastName + ", and I have " + strconv.Itoa(totalWeapons) + " weapons."
	return totalWeapons, greet
}
