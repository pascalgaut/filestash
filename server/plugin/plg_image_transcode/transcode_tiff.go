package plg_image_transcode

import (
	. "github.com/pascalgaut/filestash/server/common"
	_ "golang.org/x/image/tiff"
	"image"
	"image/jpeg"
	"io"
)

func transcodeTiff(reader io.Reader) (io.ReadCloser, string, error) {
	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, "", err
	}

	r, w := io.Pipe()
	go func() {
		err := jpeg.Encode(w, img, &jpeg.Options{Quality: 80})
		w.Close()
		if err != nil {
			Log.Debug("plg_image_transcode::tiff jpeg encoding error '%s'", err.Error())
		}
	}()
	return NewReadCloserFromReader(r), "image/jpeg", nil
}
