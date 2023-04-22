package example5

type IProfessionDirector interface {
	SetActor(actor IActorBuilder)
	Make()
}

type WarriorDirector struct {
	Actor IActorBuilder
}

func (w *WarriorDirector) SetActor(actor IActorBuilder) {
	w.Actor = actor
}

func (w *WarriorDirector) Make() {
	w.Actor.SetStrength(10)
	w.Actor.SetAgility(5)
	w.Actor.SetIntelligence(1)
}

type WizardDirector struct {
	Actor IActorBuilder
}

func (w *WizardDirector) SetActor(actor IActorBuilder) {
	w.Actor = actor
}

func (w *WizardDirector) Make() {
	w.Actor.SetStrength(1)
	w.Actor.SetAgility(5)
	w.Actor.SetIntelligence(10)
}
