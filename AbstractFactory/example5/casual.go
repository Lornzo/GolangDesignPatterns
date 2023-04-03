package example5

type CasualToteBag struct{}

func (t CasualToteBag) GetToteBagName() string {
	return "Casual Tote Bag"
}

type CasualShoulderBag struct{}

func (s CasualShoulderBag) GetShoulderBagName() string {

	return "Casual Shoulder Bag"
}

type CasualBagFactory struct{}

func (f CasualBagFactory) CreateToteBag() IToteBag {

	return CasualToteBag{}
}

func (f CasualBagFactory) CreateShoulderBag() IShoulderBag {

	return CasualShoulderBag{}
}
