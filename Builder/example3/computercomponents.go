package example3

type IComputerComponent interface {
	GetType() string
	GetName() string
}

type CpuAMD struct{}

func (cpu CpuAMD) GetType() string {
	return "CPU"
}

func (cpu CpuAMD) GetName() string {
	return "AMD"
}

type CpuIntel struct{}

func (cpu CpuIntel) GetType() string {
	return "CPU"
}

func (cpu CpuIntel) GetName() string {
	return "Intel"
}

type RamKingston struct{}

func (ram RamKingston) GetType() string {
	return "RAM"
}

func (ram RamKingston) GetName() string {
	return "Kingston 16GB DDR4"
}

type RamMicron struct{}

func (ram RamMicron) GetType() string {
	return "RAM"
}

func (ram RamMicron) GetName() string {
	return "Micron 16GB DDR4"
}

type SSDSamsung struct{}

func (ssd SSDSamsung) GetType() string {
	return "SSD"
}

func (ssd SSDSamsung) GetName() string {
	return "Samsung 1TB SSD"
}

type SSDMicron struct{}

func (ssd SSDMicron) GetType() string {
	return "SSD"
}

func (ssd SSDMicron) GetName() string {
	return "Micron 1TB SSD"
}

type NvidiaGTX1080 struct{}

func (gpu NvidiaGTX1080) GetType() string {
	return "GraphCard"
}

func (gpu NvidiaGTX1080) GetName() string {
	return "Nvidia GTX 1080"
}

type NvidiaGTX2080Ti struct{}

func (gpu NvidiaGTX2080Ti) GetType() string {
	return "GraphCard"
}

func (gpu NvidiaGTX2080Ti) GetName() string {
	return "Nvidia GTX 2080 Ti"
}
