package main

import (
	"fmt"
)

type FrameBuffer struct {
	X, Y   int
	Bpp    int    // Byte per pixel
	buffer []byte // Byte stream
}

func NewFrameBuffer(x, y int, bpp int) *FrameBuffer {
	var fb FrameBuffer
	fb.X, fb.Y = x, y
	fb.Bpp = bpp
	fb.buffer = make([]byte, x*y*bpp)
	return &fb
}

func NewFrameBufferFromBuffer(x, y int, buff []byte) *FrameBuffer {
	return &FrameBuffer{
		X: x,
		Y: y,
		Bpp: len(buff) / (x * y),
		buffer : buff,
	}
}

func (fb *FrameBuffer) CheckValid() bool {
	return len(fb.buffer) >= fb.X * fb.Y * fb.Bpp
}

func main() {
	// ASSIGN_START OMIT
	var fbRGB565 FrameBuffer
	fbRGB565.X, fbRGB565.Y = 640, 480
	fbRGB565.Bpp = 2

	fbRGB888 := FrameBuffer{640, 480, 3, nil}

	fbGray := &FrameBuffer{
		X:   640,
		Y:   480,
		Bpp: 1,
	}
	// ASSIGN_END OMIT

	fmt.Println(fbRGB888, fbRGB565, fbGray)
}
