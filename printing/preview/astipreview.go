//go:build ELECTRON
// +build ELECTRON

package preview

import (
	"bytes"
	"encoding/base64"
	"errors"
	"image"
	"image/png"

	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
)

type AstiPreview struct {
	Asti *astilectron.Astilectron
}

func (r *AstiPreview) Name() string {
	return "Preview Printing"
}

func (r *AstiPreview) Description() string {
	return "Instead of printing show a window with the result."
}

func (r *AstiPreview) AvailableEndpoints() (map[string]string, error) {
	return map[string]string{
		"Preview Printing": "window",
	}, nil
}

func (r *AstiPreview) Print(printerEndpoint string, image image.Image, data []byte) error {
	if r.Asti == nil {
		return errors.New("not initialized")
	}

	buf := &bytes.Buffer{}
	if err := png.Encode(buf, image); err != nil {
		return err
	}

	height := image.Bounds().Max.Y
	if height < 300 {
		height = 300
	}

	var w, _ = r.Asti.NewWindow("data:image/png;base64,"+base64.StdEncoding.EncodeToString(buf.Bytes()), &astilectron.WindowOptions{
		Center: astikit.BoolPtr(true),
		Height: astikit.IntPtr(image.Bounds().Max.Y + 50),
		Width:  astikit.IntPtr(image.Bounds().Max.X + 50),
	})
	return w.Create()
}
