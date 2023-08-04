package request

type VendorRequest struct {
	Name    string `json:"name" validate:"required"`
	Address string `json:"address" validate:"required"`
	Phone   string `json:"phone" validate:"required"`
	Brand   string `json:"brand" validate:"required"`
}
