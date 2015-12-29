package idealpostcodes

// Address ...
type Address struct {
	Postcode                string      `json:"postcode"`
	PostcodeInward          string      `json:"postcode_inward"`
	PostcodeOutward         string      `json:"postcode_outward"`
	PostTown                string      `json:"post_town"`
	DependantLocality       string      `json:"dependant_locality"`
	DoubleDependantLocality string      `json:"double_dependant_locality"`
	Thoroughfare            string      `json:"thoroughfare"`
	DependantThoroughfare   string      `json:"dependant_thoroughfare"`
	BuildingNumber          string      `json:"building_number"`
	BuildingName            string      `json:"building_name"`
	SubBuildingName         string      `json:"sub_building_name"`
	POBox                   string      `json:"po_box"`
	DepartmentName          string      `json:"department_name"`
	OrganisationName        string      `json:"organisation_name"`
	UDPRN                   int64       `json:"udprn"`
	PostcodeType            string      `json:"postcode_type"`
	SuOrganisationIndicator string      `json:"su_organisation_indicator"`
	DeliveryPointSuffix     string      `json:"delivery_point_suffix"`
	Line1                   string      `json:"line_1"`
	Line2                   string      `json:"line_2"`
	Line3                   string      `json:"line_3"`
	Premise                 string      `json:"premise"`
	Country                 string      `json:"country"`
	County                  string      `json:"county"`
	AdministrativeCounty    string      `json:"administrative_county"`
	PostalCounty            string      `json:"postal_county"`
	TraditionalCounty       string      `json:"traditional_county"`
	District                string      `json:"district"`
	Ward                    string      `json:"ward"`
	Longitude               interface{} `json:"longitude"` // number or empty string
	Latitude                interface{} `json:"latitude"`  // number or empty string
	Eastings                interface{} `json:"eastings"`  // number or empty string
	Northings               interface{} `json:"northings"` // number or empty string
}
