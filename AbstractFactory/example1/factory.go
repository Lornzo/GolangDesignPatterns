package example1

type IWeaponAbstractFactory interface {
	CreateAttackPlusWeapon() IWeapon
	CreateSpeedPlusWeapon() IWeapon
	CreateScopePlusWeapon() IWeapon
}

type NormalSwordFactory struct{}

func (n NormalSwordFactory) CreateAttackPlusWeapon() IWeapon {
	return Sword{AbstractWeapon{"Sword", "Normal", "normal attack"}}
}

func (n NormalSwordFactory) CreateSpeedPlusWeapon() IWeapon {
	return Sword{AbstractWeapon{"Sword", "Normal", "normal attack"}}
}

func (n NormalSwordFactory) CreateScopePlusWeapon() IWeapon {
	return Sword{AbstractWeapon{"Sword", "Normal", "normal attack"}}
}

type EpicSwrodFactory struct{}

func (e EpicSwrodFactory) CreateAttackPlusWeapon() IWeapon {
	return Sword{AbstractWeapon{"Sword", "Epic", "attack plus"}}
}

func (e EpicSwrodFactory) CreateSpeedPlusWeapon() IWeapon {
	return Sword{AbstractWeapon{"Sword", "Epic", "speed plus"}}
}

func (e EpicSwrodFactory) CreateScopePlusWeapon() IWeapon {
	return Sword{AbstractWeapon{"Sword", "Epic", "scope plus"}}
}

type LegendaryBowFactory struct{}

func (l LegendaryBowFactory) CreateAttackPlusWeapon() IWeapon {
	return Bow{AbstractWeapon{"Bow", "Legendary", "attack plus"}}
}

func (l LegendaryBowFactory) CreateSpeedPlusWeapon() IWeapon {
	return Bow{AbstractWeapon{"Bow", "Legendary", "speed plus"}}
}

func (l LegendaryBowFactory) CreateScopePlusWeapon() IWeapon {
	return Bow{AbstractWeapon{"Bow", "Legendary", "scope plus"}}
}
