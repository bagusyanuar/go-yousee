package request

import "mime/multipart"

type ItemImageRequest struct {
	ItemID        string                `json:"item_id" validate:"required,uuid4"`
	ImageInternal *multipart.FileHeader `form:"image_internal"`
	ImageCliet    *multipart.FileHeader `form:"image_client"`
}
