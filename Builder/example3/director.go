package example3

type IComputerDirector interface {
	MakeComputer(builder IComputerBuilder)
}

type OfficeDirector struct{}

func (o OfficeDirector) MakeComputer(builder IComputerBuilder) {
	builder.AddCPU(1)
	builder.AddRAM(1)
	builder.AddSSD(1)
}

type GamingDirector struct{}

func (g GamingDirector) MakeComputer(builder IComputerBuilder) {
	builder.AddCPU(2)
	builder.AddRAM(2)
	builder.AddSSD(2)
	builder.AddGraphCard(2)
}
