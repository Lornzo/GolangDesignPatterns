package example1

import "fmt"

type Role interface {
	Attack()
}

type RoleWarrior struct{}

func (r RoleWarrior) Attack() {
	fmt.Println("Warrior attacks with his sword and shield")
}

type RoleMagician struct{}

func (r RoleMagician) Attack() {
	fmt.Println("Magician attacks with magic and his staff")
}

type RoleAssassin struct{}

func (r RoleAssassin) Attack() {
	fmt.Println("Assassin take a sneak attacks with his dagger")
}
