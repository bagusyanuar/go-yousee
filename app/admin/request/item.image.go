package request

import "mime/multipart"

type ItemImageRequest struct {
	ItemID string                `form:"item_id" validate:"required,uuid4"`
	Image  *multipart.FileHeader `form:"image"`
	Type   uint8                 `form:"type" validate:"required"`
}
