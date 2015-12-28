package idealpostcodes

// ClientInterface defines exported methods
type ClientInterface interface {
	// Exported methods
	GetPostcode(postcode string) ([]*Address, error)
}
