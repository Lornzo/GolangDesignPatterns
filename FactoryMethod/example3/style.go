package example3

type Style interface {
	GetTemplate() string
}

type DarkStyle struct{}

func (d DarkStyle) GetTemplate() string {
	return "Dark template"
}

type SimpleStyle struct{}

func (s SimpleStyle) GetTemplate() string {
	return "Simple template"
}

type BrightStyle struct{}

func (b BrightStyle) GetTemplate() string {
	return "Bright template"
}
