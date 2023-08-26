package common

import (
	"bytes"
	"encoding/base64"
	"image/png"
	"mime/multipart"
)

const (
	DateFormat     string = "2006-01-02"
	CodeDateFormat string = "20060102150405"
)

func ConvertToBase64(file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	img, err := png.Decode(src)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	e := png.Encode(&buf, img)
	if e != nil {
		return "", err
	}
	imgByte := buf.Bytes()
	imgBase64Str := "data:image/png;base64," + base64.StdEncoding.EncodeToString(imgByte)
	return imgBase64Str, nil
}
