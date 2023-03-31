package example2

type CarCreator interface {
	CreateCar() Car
}

type FamilyCarCreator struct{}

func (f FamilyCarCreator) CreateCar() Car {
	return FamilyCar{}
}

type SUVCarCreator struct{}

func (s SUVCarCreator) CreateCar() Car {
	return SUVCar{}
}

type SportsCarCreator struct{}

func (s SportsCarCreator) CreateCar() Car {
	return SportsCar{}
}
