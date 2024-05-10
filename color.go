package main

type RGB struct {
	R, G, B uint8
}

func (c RGB) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R)
	r |= r << 8
	g = uint32(c.G)
	g |= g << 8
	b = uint32(c.B)
	b |= b << 8
	a = uint32(255)
	a |= a << 8
	return
}

func (c RGB) Scaled(f float64) RGB {
	c.R = uint8(float64(c.R) * f)
	c.G = uint8(float64(c.G) * f)
	c.B = uint8(float64(c.B) * f)
	return c
}

func (a RGB) MixSub(b RGB) RGB {
	a.R = uint8(uint16(a.R) * uint16(b.R) / 255)
	a.G = uint8(uint16(a.G) * uint16(b.G) / 255)
	a.B = uint8(uint16(a.B) * uint16(b.B) / 255)
	return a
}

func (a RGB) MixAdd(b RGB) RGB {
	a.R = 255 - uint8((float64(255-a.R)/255*float64(255-b.R)/255)*255)
	a.G = 255 - uint8((float64(255-a.G)/255*float64(255-b.G)/255)*255)
	a.B = 255 - uint8((float64(255-a.B)/255*float64(255-b.B)/255)*255)
	return a
}
