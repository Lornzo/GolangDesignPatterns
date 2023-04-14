package example3

type IComputerBuilder interface {
	AddCPU(num int)
	AddRAM(num int)
	AddSSD(num int)
	AddGraphCard(num int)
}

type AMDBuilder struct {
	computer AMDComputer
}

func (a *AMDBuilder) AddCPU(num int) {
	for i := 0; i < num; i++ {
		a.computer.CPU = append(a.computer.CPU, CpuAMD{})
	}
}

func (a *AMDBuilder) AddRAM(num int) {
	for i := 0; i < num; i++ {
		a.computer.RAM = append(a.computer.RAM, RamKingston{})
	}
}

func (a *AMDBuilder) AddSSD(num int) {
	for i := 0; i < num; i++ {
		a.computer.SSD = append(a.computer.SSD, SSDSamsung{})
	}
}

func (a *AMDBuilder) AddGraphCard(num int) {
	for i := 0; i < num; i++ {
		a.computer.GraphCard = append(a.computer.GraphCard, NvidiaGTX1080{})
	}
}

func (a *AMDBuilder) GetAMDComputer() AMDComputer {
	return a.computer
}

type IntelBuilder struct {
	computer IntelComputer
}

func (i *IntelBuilder) AddCPU(num int) {
	for j := 0; j < num; j++ {
		i.computer.CPU = append(i.computer.CPU, CpuIntel{})
	}
}

func (i *IntelBuilder) AddRAM(num int) {
	for j := 0; j < num; j++ {
		i.computer.RAM = append(i.computer.RAM, RamMicron{})
	}
}

func (i *IntelBuilder) AddSSD(num int) {
	for j := 0; j < num; j++ {
		i.computer.SSD = append(i.computer.SSD, SSDMicron{})
	}
}

func (i *IntelBuilder) AddGraphCard(num int) {
	for j := 0; j < num; j++ {
		i.computer.GraphCard = append(i.computer.GraphCard, NvidiaGTX2080Ti{})
	}
}

func (i *IntelBuilder) GetIntelComputer() IntelComputer {
	return i.computer
}
