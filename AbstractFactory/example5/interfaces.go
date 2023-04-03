package example5

type IToteBag interface {
	GetToteBagName() string
}

type IShoulderBag interface {
	GetShoulderBagName() string
}

type IBagFactory interface {
	CreateToteBag() IToteBag
	CreateShoulderBag() IShoulderBag
}
