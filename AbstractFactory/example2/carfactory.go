package example2

type ICarFactory interface {
	GetProduceType() string
	CreateEngine() ICarEngine
	CreateBody() ICarBody
	CreateWheel() ICarWheel
}

type VanFactory struct{}

func (v VanFactory) GetProduceType() string {
	return "廂型車"
}

func (v VanFactory) CreateEngine() ICarEngine {
	return RegularEngine{}
}

func (v VanFactory) CreateBody() ICarBody {
	return VanBody{}
}

func (v VanFactory) CreateWheel() ICarWheel {
	return RegularWheel{}
}

type SportCarFactory struct{}

func (s SportCarFactory) GetProduceType() string {
	return "跑車"
}

func (s SportCarFactory) CreateEngine() ICarEngine {
	return V8Engine{}
}

func (s SportCarFactory) CreateBody() ICarBody {
	return SportCarBody{}
}

func (s SportCarFactory) CreateWheel() ICarWheel {
	return LuxuryWheel{}
}
