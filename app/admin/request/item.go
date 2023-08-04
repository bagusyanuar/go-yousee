package request

type ItemRequest struct {
	CityID    string  `json:"city_id" validate:"required,uuid4"`
	VendorID  string  `json:"vendor_id" validate:"required,uuid4"`
	TypeID    string  `json:"type_id" validate:"required,uuid4"`
	Name      string  `json:"name" validate:"required"`
	Address   string  `json:"address" validate:"required"`
	Latitude  float64 `json:"latitude" validate:"required"`
	Longitude float64 `json:"longitude" validate:"required"`
	Position  uint8   `json:"position" validate:"numeric"`
	Width     float64 `json:"width" validate:"numeric"`
	Height    float64 `json:"height" validate:"numeric"`
}
