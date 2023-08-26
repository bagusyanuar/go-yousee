package request

type ProjectItemRequest struct {
	ProjectID   string `json:"project_id" validate:"uuid4"`
	CityID      string `json:"city_id" validate:"required,uuid4"`
	ItemID      string `json:"item_id" validate:"omitempty,uuid4"`
	PicID       string `json:"pic_id" validate:"required,uuid4"`
	VendorPrice int64  `json:"vendor_price" validate:"numeric"`
}
