package example1_test

import (
	"testing"

	"github.com/Lornzo/GolangDesignPatterns/GolangDesignPatterns/FactoryMethod/example1"
)

func TestPlay(t *testing.T) {
	play(example1.WarriorCreator{})
	play(example1.MagicianCreator{})
	play(example1.AssassinCreator{})
}

func play(creator example1.RoleCreator) {

	var (
		role example1.Role = creator.CreateRole()
	)

	role.Attack()

}
