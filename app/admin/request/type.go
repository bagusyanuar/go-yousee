package request

import "mime/multipart"

type TypeRequest struct {
	Name string                `json:"name" validate:"required"`
	Icon *multipart.FileHeader `json:"icon"`
}
