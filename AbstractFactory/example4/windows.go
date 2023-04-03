package example4

type WindowsAVI struct{}

func (w *WindowsAVI) GetAVICode() string {
	return "Windows AVI"
}

type WindowsMOV struct{}

func (w *WindowsMOV) GetMOVCode() string {
	return "Windows MOV"
}

type WindowsMP4 struct{}

func (w *WindowsMP4) GetMP4Code() string {
	return "Windows MP4"
}

type Windows struct{}

func (w *Windows) GetAVI() IAVI {
	return &WindowsAVI{}
}

func (w *Windows) GetMOV() IMOV {
	return &WindowsMOV{}
}

func (w *Windows) GetMP4() IMP4 {
	return &WindowsMP4{}
}
