package example1

type IDirector interface {
	SetBuilder(builder IBuilder)
	LetBuilderBuildingHouse()
}

type SimpleHouseDirector struct {
	builder IBuilder
}

func (s *SimpleHouseDirector) SetBuilder(builder IBuilder) {
	s.builder = builder
}

func (s *SimpleHouseDirector) LetBuilderBuildingHouse() {
	s.builder.BuildRoom(1)
	s.builder.BuildBedRoom(1)
	s.builder.BuildKitchen(1)
	s.builder.BuildBathRoom(1)
}

type LuxuryHouseDirector struct {
	builder IBuilder
}

func (l *LuxuryHouseDirector) SetBuilder(builder IBuilder) {
	l.builder = builder
}

func (l *LuxuryHouseDirector) LetBuilderBuildingHouse() {
	l.builder.BuildRoom(2)
	l.builder.BuildBedRoom(2)
	l.builder.BuildKitchen(1)
	l.builder.BuildLivingRoom(1)
	l.builder.BuildBathRoom(2)
	l.builder.BuildGarage(1)
}
