package frame256x288_test

import (
	"github.com/reiver/go-frame256x288"

	"github.com/reiver/go-dast8x8"

	"image"
	"image/color"

	"testing"
)

func TestSlice_Draw(t *testing.T) {

	var buffer [frame256x288.ByteSize]uint8

	var frame frame256x288.Slice = frame256x288.Slice(buffer[:])

	if err := frame.Dye(color.RGBA{0,111,184,255}); nil != err {
		t.Errorf("Received an error but did not expect one.")
		t.Logf("ERROR: (%T) %s", err, err)
		return
	}

	{
		var img image.Image = dast8x8.Sprite(0)

		if err := frame.Draw(img); nil != err {
			t.Errorf("Received an error but did not expect one.")
			t.Logf("ERROR: (%T) %s", err, err)
			return
		}
	}

	{
		expected := "IMAGE:iVBORw0KGgoAAAANSUhEUgAAAQAAAAEgCAIAAADUvDoHAAADkElEQVR4nOzYsa3TQByA8bP12uzCCFQRC7w5PEtgBOjeApEr6NkmAxgFp4gYIJb4fj9Zyl33L+67JH4by3WMsV7ex1/n5WNAxvx8+v9Zw3/vbf84j9u2bdM0HT0PvNQjgHWcfv349vP71zHG599HDwWvMq3jtH8D7PvH1j8BGub9uK/jtD/n5ePTdjt6KniRaSzXcflyXz4voGFeL+/73b///rk/XgSRMbnvKZuPHgCOJADSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApP0JAAD//9JMK6ywQ6wwAAAAAElFTkSuQmCC"
		actual := frame.String()

		if expected != actual {
			t.Errorf("The actual serialized image is not what was expected.")
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			return
		}
	}
}
