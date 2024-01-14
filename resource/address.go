package resource

import "github.com/maantje/postcode/.gen/model"

type AddressResource struct {
	Postcode     string  `json:"postcode"`
	HouseNumber  int32   `json:"house_number"`
	Addition     string  `json:"addition"`
	Extension    string  `json:"extension"`
	Municipality string  `json:"municipality"`
	City         string  `json:"city"`
	Street       string  `json:"street"`
	LongStreet   string  `json:"long_street"`
	ShortStreet  string  `json:"short_street"`
	Area         float32 `json:"area"`
	Usage        string  `json:"usage"`
	BuiltIn      int32   `json:"built_in"`
	Latitude     float32 `json:"latitude"`
	Longitude    float32 `json:"longitude"`
}

func NewAddressResource(a *model.Addresses) AddressResource {
	return AddressResource{
		Postcode:     *a.Postcode,
		HouseNumber:  *a.HouseNumber,
		Addition:     *a.Addition,
		Extension:    *a.Extension,
		Municipality: *a.Municipality,
		City:         *a.City,
		Street:       *a.Street,
		LongStreet:   *a.LongStreet,
		ShortStreet:  *a.ShortStreet,
		Area:         *a.Area,
		Usage:        *a.Usage,
		BuiltIn:      *a.BuiltIn,
		Latitude:     *a.Latitude,
		Longitude:    *a.Longitude,
	}
}
