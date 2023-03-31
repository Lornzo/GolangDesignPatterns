package example3

type StyleCreator interface {
	CreateStyle() Style
}

type DarkCreator struct{}

func (d DarkCreator) CreateStyle() Style {
	return DarkStyle{}
}

type SimpleCreator struct{}

func (s SimpleCreator) CreateStyle() Style {
	return SimpleStyle{}
}

type BrightCreator struct{}

func (b BrightCreator) CreateStyle() Style {
	return BrightStyle{}
}
