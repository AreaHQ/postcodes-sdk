package idealpostcodes

// GetPostcodeResponse ...
type GetPostcodeResponse struct {
	Result  []*Address `json:"result"`
	Code    int        `json:"code"`
	Message string     `json:"message"`
}
