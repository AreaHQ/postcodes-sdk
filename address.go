package idealpostcodes

// Address ...
type Address struct {
	Postcode                string      `json:"postcode"`
	PostcodeInward          string      `json:"postcode_inward,omitempty"`
	PostcodeOutward         string      `json:"postcode_outward,omitempty"`
	PostTown                string      `json:"post_town,omitempty"`
	DependantLocality       string      `json:"dependant_locality,omitempty"`
	DoubleDependantLocality string      `json:"double_dependant_locality,omitempty"`
	Thoroughfare            string      `json:"thoroughfare,omitempty"`
	DependantThoroughfare   string      `json:"dependant_thoroughfare,omitempty"`
	BuildingNumber          string      `json:"building_number,omitempty"`
	BuildingName            string      `json:"building_name,omitempty"`
	SubBuildingName         string      `json:"sub_building_name,omitempty"`
	POBox                   string      `json:"po_box,omitempty"`
	DepartmentName          string      `json:"department_name,omitempty"`
	OrganisationName        string      `json:"organisation_name,omitempty"`
	UDPRN                   int64       `json:"udprn,omitempty"`
	PostcodeType            string      `json:"postcode_type,omitempty"`
	SuOrganisationIndicator string      `json:"su_organisation_indicator,omitempty"`
	DeliveryPointSuffix     string      `json:"delivery_point_suffix,omitempty"`
	Line1                   string      `json:"line_1"`
	Line2                   string      `json:"line_2,omitempty"`
	Line3                   string      `json:"line_3,omitempty"`
	Premise                 string      `json:"premise,omitempty"`
	Country                 string      `json:"country,omitempty"`
	County                  string      `json:"county,omitempty"`
	AdministrativeCounty    string      `json:"administrative_county,omitempty"`
	PostalCounty            string      `json:"postal_county,omitempty"`
	TraditionalCounty       string      `json:"traditional_county,omitempty"`
	District                string      `json:"district,omitempty"`
	Ward                    string      `json:"ward,omitempty"`
	Longitude               interface{} `json:"longitude,omitempty"` // number or empty string
	Latitude                interface{} `json:"latitude,omitempty"`  // number or empty string
	Eastings                interface{} `json:"eastings,omitempty"`  // number or empty string
	Northings               interface{} `json:"northings,omitempty"` // number or empty string
}
