package utils

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"mime/multipart"
)

type Image struct {
}

type ImageInfo struct {
	Width  int
	Height int
}

func (u *Image) Decode(file multipart.File) (*ImageInfo, error) {
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	info := &ImageInfo{
		Width:  img.Bounds().Dx(),
		Height: img.Bounds().Dy(),
	}

	return info, nil
}
