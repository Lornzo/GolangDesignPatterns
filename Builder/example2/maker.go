package example2

type IBeverageMaker interface {
	SetBuilder(builder IBeverageBuilder)
	Make()
}

type HelfSugerMaker struct {
	builder IBeverageBuilder
}

func (h *HelfSugerMaker) SetBuilder(builder IBeverageBuilder) {
	h.builder = builder
}

func (h *HelfSugerMaker) Make() {
	h.builder.SetSugar(5)
}

type HighTemperatureMaker struct {
	builder IBeverageBuilder
}

func (h *HighTemperatureMaker) SetBuilder(builder IBeverageBuilder) {
	h.builder = builder
}

func (h *HighTemperatureMaker) Make() {
	h.builder.SetTemperature(100)
}
