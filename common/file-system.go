package common

import (
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

func (fs *FileSystem) UploadAndResize(dst string, width uint) error {

	src, err := fs.File.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// buffer, err := io.ReadAll(src)
	// if err != nil {
	// 	return err
	// }

	// resized, err := bimg.NewImage(buffer).Resize(int(width), 0)
	// if err != nil {
	// 	return err
	// }

	// compressed, err := bimg.NewImage(resized).Process(bimg.Options{Quality: 30})
	// if err != nil {
	// 	return err
	// }
	// return bimg.Write(dst, compressed)
	img, err := png.Decode(src)
	if err != nil {
		return err
	}

	m := resize.Resize(width, 0, img, resize.Lanczos3)
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	png.Encode(out, m)

	return err
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
