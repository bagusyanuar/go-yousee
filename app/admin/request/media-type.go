package request

import "mime/multipart"

type MediaTypeRequest struct {
	Name string                `form:"name" validate:"required"`
	Icon *multipart.FileHeader `form:"icon"`
}
