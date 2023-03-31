package example5

type Medicine interface {
	GetIngredients() string
}

type PainKiller struct{}

func (p *PainKiller) GetIngredients() string {
	return "Aspirin"
}

type Antipyretics struct{}

func (a *Antipyretics) GetIngredients() string {
	return "Paracetamol"
}

type AntiInflammatory struct{}

func (a *AntiInflammatory) GetIngredients() string {
	return "Ibuprofen"
}
