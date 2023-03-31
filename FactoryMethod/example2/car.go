package example2

type Car interface {
	Interduce() string
}

type FamilyCar struct{}

func (f FamilyCar) Interduce() string {
	return "I am a family car, I suit driving in the city"
}

type SUVCar struct{}

func (s SUVCar) Interduce() string {
	return "I am a SUV car, I suit driving in the mountain"
}

type SportsCar struct{}

func (s SportsCar) Interduce() string {
	return "I am a sports car, I suit driving in the highway"
}
