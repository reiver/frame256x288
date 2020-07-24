package frame256x288

import (
	"github.com/reiver/go-rgba32"

	"image"
	"image/color"
	"image/draw"
)

type Slice []uint8

func (receiver Slice) at(x, y int) []uint8 {
	if nil == receiver {
		return nil
	}
	if x < 0 || Width <= x {
		return nil
	}
	if y < 0 || Height <= y {
		return nil
	}

	offset := receiver.PixOffset(x,y)

	low  := offset
	high := low + rgba32.ByteSize
	if rgba32.ByteSize != (high-low) {
		return nil
	}

	p := receiver[low:high]

	return rgba32.Slice(p)
}

func (receiver Slice) At(x, y int) color.Color {
	p := receiver.at(x,y)
	return rgba32.Slice(p)
}

func (receiver Slice) Bounds() image.Rectangle {
	const x = 0
	const y = 0

	// [x,x+Width) and [y,y+Height)
	return image.Rectangle{
		Min: image.Point{
			X: x,
			Y: y,
		},
		Max: image.Point{
			X: x+Width,
			Y: y+Height,
		},
	}
}

func (receiver Slice) ColorModel() color.Model {
	return color.NRGBAModel
}

func (receiver Slice) Draw(img image.Image) error {
	if nil == receiver {
		return errNilReceiver
	}

	rect := img.Bounds()

	draw.Draw(receiver, rect, img, rect.Min, draw.Over)

	return nil
}

// Dye changes the color of all the pixels / pels in this from to ‘color’.
func (receiver Slice) Dye(c color.Color) error {
	if nil == receiver {
		return errNilReceiver
	}

	var r,g,b,a uint8
	{
		switch casted := c.(type) {
		case rgba32.Slice:
			r = casted[rgba32.OffsetRed]
			g = casted[rgba32.OffsetGreen]
			b = casted[rgba32.OffsetBlue]
			a = casted[rgba32.OffsetAlpha]
		default:
			rr,gg,bb,aa := casted.RGBA()

			r = uint8((rr*0xff)/0xffff)
			g = uint8((gg*0xff)/0xffff)
			b = uint8((bb*0xff)/0xffff)
			a = uint8((aa*0xff)/0xffff)
		}
	}

	for y:=0; y<Height; y++ {
		for x:=0; x<Width; x++ {
			receiver.set(x,y, r,g,b,a)
		}
	}

	return nil
}

func (receiver Slice) PixOffset(x int, y int) int {
	return y*(Width*Depth) + x*Depth
}

func (receiver Slice) set(x, y int, r, g, b, a uint8) {
	if nil == receiver {
		return
	}

	p := receiver.at(x,y)

	p[rgba32.OffsetRed]   = r
	p[rgba32.OffsetGreen] = g
	p[rgba32.OffsetBlue]  = b
	p[rgba32.OffsetAlpha] = a
}

// Set helps Slice fit the Go built-in draw.Image interface.
//
// Set will change the color of the pixel / pel, in this frame,
// at (‘x’,‘y’) to ‘color’.
func (receiver Slice) Set(x, y int, c color.Color) {
	if nil == receiver {
		return
	}
	if nil == c {
		return
	}

	u32r, u32g, u32b, u32a := c.RGBA()

	u8r := uint8((u32r*0xff)/0xffff)
	u8g := uint8((u32g*0xff)/0xffff)
	u8b := uint8((u32b*0xff)/0xffff)
	u8a := uint8((u32a*0xff)/0xffff)

	receiver.set(x,y, u8r, u8g, u8b, u8a)
}
