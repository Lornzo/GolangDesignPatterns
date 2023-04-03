package example5

type LuxuryToteBag struct{}

func (t LuxuryToteBag) GetToteBagName() string {
	return "Luxury Tote Bag"
}

type LuxuryShoulderBag struct{}

func (s LuxuryShoulderBag) GetShoulderBagName() string {
	return "Luxury Shoulder Bag"
}

type LuxuryBagFactory struct{}

func (f LuxuryBagFactory) CreateToteBag() IToteBag {
	return LuxuryToteBag{}
}

func (f LuxuryBagFactory) CreateShoulderBag() IShoulderBag {
	return LuxuryShoulderBag{}
}
