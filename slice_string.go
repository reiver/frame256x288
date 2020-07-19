package frame256x288

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/png"
	"strings"
)

func (receiver Slice) String() string {
	var buffer strings.Builder

	buffer.WriteString("IMAGE:")

	{
		var pngBuffer bytes.Buffer

		err := png.Encode(&pngBuffer, receiver)
		if nil != err {
			return fmt.Sprintf("ERROR:%s", err)
		}

		encoded := base64.StdEncoding.EncodeToString(pngBuffer.Bytes())

		buffer.WriteString(encoded)
	}

	return buffer.String()
}
