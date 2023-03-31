package example5

type MedicineFactory interface {
	MakeMedicine() Medicine
}

type PainKillerFactory struct{}

func (p *PainKillerFactory) MakeMedicine() Medicine {
	return &PainKiller{}
}

type AntipyreticsFactory struct{}

func (a *AntipyreticsFactory) MakeMedicine() Medicine {
	return &Antipyretics{}
}

type AntiInflammatoryFactory struct{}

func (a *AntiInflammatoryFactory) MakeMedicine() Medicine {
	return &AntiInflammatory{}
}
