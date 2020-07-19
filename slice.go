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

func (receiver Slice) Draw(img image.Image, x int, y int) error {
	if nil == receiver {
		return errNilReceiver
	}

	var rect image.Rectangle
	{
		bounds := img.Bounds()

		width  := bounds.Max.X - bounds.Min.X
		height := bounds.Max.Y - bounds.Min.Y

		rect = image.Rectangle{
			Min: image.Point{
				X:x,
				Y:y,
			},
			Max: image.Point{
				X:x+width,
				Y:y+height,
			},
		}
	}

	draw.Draw(receiver, rect, img, image.ZP, draw.Src)

	return nil
}

func (receiver Slice) PixOffset(x int, y int) int {
	return y*(Width*Depth) + x*Depth
}

func (receiver Slice) Set(x, y int, c color.Color) {
	if nil == receiver {
		return
	}
	if nil == c {
		return
	}

	p := receiver.at(x,y)

	u32r, u32g, u32b, u32a := c.RGBA()

	u8r := uint8((u32r*0xff)/0xffff)
	u8g := uint8((u32g*0xff)/0xffff)
	u8b := uint8((u32b*0xff)/0xffff)
	u8a := uint8((u32a*0xff)/0xffff)

	p[rgba32.OffsetRed]   = u8r
	p[rgba32.OffsetGreen] = u8g
	p[rgba32.OffsetBlue]  = u8b
	p[rgba32.OffsetAlpha] = u8a
}
