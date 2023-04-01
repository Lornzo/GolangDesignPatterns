package example2

type ICarEngine interface {
	Run() string
}

type RegularEngine struct{}

func (r RegularEngine) Run() string {
	return "普通的引擎冒著白煙開始運作"
}

type V8Engine struct{}

func (v V8Engine) Run() string {
	return "V8引擎轟鳴開始運作"
}
