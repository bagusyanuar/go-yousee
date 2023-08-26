package common

import (
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"mime/multipart"
	"os"

	"github.com/nfnt/resize"
)

type FileSystem struct {
	File *multipart.FileHeader
}

func (fs *FileSystem) Upload(dst string) error {
	src, err := fs.File.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

func (fs *FileSystem) UploadAndResize(dst string, width uint, ext string) error {

	src, err := fs.File.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// img, err := png.Decode(src)
	// if err != nil {
	// 	return err
	// }

	img, err := fs.decodeImage(src, ext)
	if err != nil {
		return err
	}

	m := resize.Resize(width, 0, img, resize.Lanczos3)
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	// png.Encode(out, m)
	fs.encodeImage(out, m, ext)

	return nil
}

func (fs *FileSystem) decodeImage(file multipart.File, ext string) (image.Image, error) {
	var (
		img image.Image
		err error
	)
	switch ext {
	case ".jpeg", ".jpg":
		img, err = jpeg.Decode(file)
		if err != nil {
			return img, err
		}
		return img, nil
	case ".png":
		img, err := png.Decode(file)
		if err != nil {
			return img, err
		}
		return img, nil
	default:
		return nil, errors.New("unknown image type")
	}
}

func (fs *FileSystem) encodeImage(file *os.File, img image.Image, ext string) {

	switch ext {
	case ".jpeg", ".jpg":
		jpeg.Encode(file, img, nil)
	case ".png":
		png.Encode(file, img)
	default:
	}
}

func (fs *FileSystem) CheckPath(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	} else {
		return err
	}
	return nil
}
