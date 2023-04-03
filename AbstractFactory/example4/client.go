package example4

import "fmt"

type Client struct {
	System ISystemFactory
}

func (c Client) Run() {
	avi := c.System.GetAVI()
	mov := c.System.GetMOV()
	mp4 := c.System.GetMP4()

	fmt.Println("AVI:", avi.GetAVICode(), "MOV:", mov.GetMOVCode(), "MP4:", mp4.GetMP4Code())
}
