package example1

type IWeapon interface {
	GetType() string
	GetLevel() string
	GetCharacteristic() string
}

type AbstractWeapon struct {
	Type           string
	Level          string
	Characteristic string
}

func (a AbstractWeapon) GetType() string {
	return a.Type
}

func (a AbstractWeapon) GetLevel() string {
	return a.Level
}

func (a AbstractWeapon) GetCharacteristic() string {
	return a.Characteristic
}

type Sword struct {
	AbstractWeapon
}

type Bow struct {
	AbstractWeapon
}
