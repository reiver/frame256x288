package frame256x288_test

import (
	"github.com/reiver/go-frame256x288"

	"image"

	"testing"
)

func TestSlice_imageimage(t *testing.T) {

	// THIS IS WHAT ACTUALLY MATTERS!
	var x image.Image = frame256x288.Slice{}

	if nil == x {
		t.Error("This should never happen")
	}
}
