package example4

type IMP4 interface {
	GetMP4Code() string
}

type IMOV interface {
	GetMOVCode() string
}

type IAVI interface {
	GetAVICode() string
}

type ISystemFactory interface {
	GetMP4() IMP4
	GetMOV() IMOV
	GetAVI() IAVI
}
