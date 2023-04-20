package example1_test

import (
	"fmt"
	"testing"

	"github.com/Lornzo/GolangDesignPatterns/AbstractFactory/example1"
)

func TestExample1(t *testing.T) {
	var (
		player        Player
		weaponFactory example1.IWeaponAbstractFactory
	)

	weaponFactory = example1.NormalSwordFactory{}
	player.Weapon = weaponFactory.CreateAttackPlusWeapon()
	player.Attack()

	weaponFactory = example1.EpicSwrodFactory{}
	player.Weapon = weaponFactory.CreateSpeedPlusWeapon()
	player.Attack()

	weaponFactory = example1.LegendaryBowFactory{}
	player.Weapon = weaponFactory.CreateScopePlusWeapon()
	player.Attack()

}

type Player struct {
	Weapon example1.IWeapon
}

func (p Player) Attack() {
	fmt.Println("player attack with", p.Weapon.GetLevel(), p.Weapon.GetType(), "weapon which has ", p.Weapon.GetCharacteristic())
}
