package example4

type LinuxAVI struct{}

func (l *LinuxAVI) GetAVICode() string {
	return "Linux AVI"
}

type LinuxMOV struct{}

func (l *LinuxMOV) GetMOVCode() string {
	return "Linux MOV"
}

type LinuxMP4 struct{}

func (l *LinuxMP4) GetMP4Code() string {
	return "Linux MP4"
}

type Linux struct{}

func (l *Linux) GetAVI() IAVI {
	return &LinuxAVI{}
}

func (l *Linux) GetMOV() IMOV {
	return &LinuxMOV{}
}

func (l *Linux) GetMP4() IMP4 {
	return &LinuxMP4{}
}
