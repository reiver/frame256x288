package frame256x288_test

import (
	"github.com/reiver/go-frame256x288"

	"github.com/reiver/go-rgba32"

	"math/rand"
	"image/color"
	"time"

	"testing"
)

func TestSlice_Set(t *testing.T) {

	randomness := rand.New(rand.NewSource( time.Now().UTC().UnixNano() ))

	const testLimit int = 20

	for testNumber:=0; testNumber<testLimit; testNumber++ {

		var buffer [frame256x288.ByteSize]uint8

		for offset, actual := range buffer {
			if expected := uint8(0); expected != actual {
				t.Errorf("For test #%d & offset=%d, actual value for uninitialized buffer is not what was expected.", testNumber, offset)
				t.Logf("EXPECTED: %d", expected)
				t.Logf("ACTUAL:   %d", actual)
				return
			}
		}


		var frame frame256x288.Slice = frame256x288.Slice(buffer[:])

		for y:=0; y<frame256x288.Height; y++ {
			for x:=0; x<frame256x288.Width; x++ {
				c := frame.At(x,y)

				aR, aG, aB, aA := c.RGBA()

				var eR, eG, eB, eA uint32 = 0,0,0,0

				if eR != aR || eG != aG || eB != aB || eA != aA {
					t.Errorf("For test #%d & (x,y)=(%d,%d), actual value for uninitialized buffer is not what was expected.", testNumber, x,y)
					t.Logf("EXPECTED (r,g,b,a)=(%d,%d,%d,%d)", eR, eG, eB, eA)
					t.Logf("ACTUAL   (r,g,b,a)=(%d,%d,%d,%d)", aR, aG, aB, aA)
					{
						p := c.(rgba32.Slice)
						t.Logf("ACTUAL: rgba(%d,%d,%d,%d)", p[0], p[1], p[2], p[3])
					}
					return
				}
			}
		}



		r := uint8(randomness.Intn(256))
		g := uint8(randomness.Intn(256))
		b := uint8(randomness.Intn(256))
		a := uint8(255)
		t.Logf("For test #%d, randomly chosen color is: rgba(%d,%d,%d,%d)", testNumber, r,g,b,a)


		xx := randomness.Intn(frame256x288.Width)
		yy := randomness.Intn(frame256x288.Height)
		switch testNumber {
		case 0:
			xx = 0
			yy = 0
		case 1:
			xx = frame256x288.Width-1
			yy = 0
		case 2:
			xx = 0
			yy = frame256x288.Height-1
		case 3:
			xx = frame256x288.Width-1
			yy = frame256x288.Height-1
		}
		t.Logf("For test #%d, randomly chosen position is: (x,y)=(%d,%d)", testNumber, xx,yy)


		{
			var c color.Color = color.NRGBA{
				R: r,
				G: g,
				B: b,
				A: a,
			}

			{
				rr, gg, bb, aa := c.RGBA()
				t.Logf("For test #%d, randomly chosen color is: (r,g,b,a)=(%d,%d,%d,%d)", testNumber, rr,gg,bb,aa)
			}

			frame.Set(xx, yy, c)
		}


		for y:=0; y<frame256x288.Height; y++ {
			for x:=0; x<frame256x288.Width; x++ {
				c := frame.At(x,y)

				aR, aG, aB, aA := c.RGBA()

				var eR, eG, eB, eA uint32
				if x == xx && y == yy {
					eR = uint32(r) * (0xffff/0xff)
					eG = uint32(g) * (0xffff/0xff)
					eB = uint32(b) * (0xffff/0xff)
					eA = uint32(a) * (0xffff/0xff)
				} else {
					eR, eG, eB, eA = 0,0,0,0
				}

				if eR != aR || eG != aG || eB != aB || eA != aA {
					t.Errorf("For test #%d & (x,y)=(%d,%d), actual value for buffer is not what was expected.", testNumber, x,y)
					t.Logf("EXPECTED (r,g,b,a)=(%d,%d,%d,%d)", eR, eG, eB, eA)
					t.Logf("ACTUAL   (r,g,b,a)=(%d,%d,%d,%d)", aR, aG, aB, aA)
					{
						p := c.(rgba32.Slice)
						t.Logf("ACTUAL: rgba(%d,%d,%d,%d)", p[0], p[1], p[2], p[3])
					}
					return
				}
			}
		}
	}
}
