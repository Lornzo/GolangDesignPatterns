package example1

type RoleCreator interface {
	CreateRole() Role
}

type WarriorCreator struct{}

func (r WarriorCreator) CreateRole() Role {
	return RoleWarrior{}
}

type MagicianCreator struct{}

func (r MagicianCreator) CreateRole() Role {
	return RoleMagician{}
}

type AssassinCreator struct{}

func (r AssassinCreator) CreateRole() Role {
	return RoleAssassin{}
}
