[![Codeship Status for areatech/idealpostcodes-sdk](https://codeship.com/projects/d170a120-8f6a-0133-90ab-124821b463ac/status?branch=master)](https://codeship.com/projects/124189)

# Postcodes SDK

SDK for [ideal-postcodes.co.uk](https://ideal-postcodes.co.uk) API.

Example usage:

```go
package main

import(
	"net/http"

	postcodes "github.com/areatech/postcodes-sdk"
)

func main() {
	client := postcodes.NewClient(
		postcodes.Endpoint,
		"your_api_key",
		new(http.Client),
	)

	addresses, err := client.GetPostcode("ID11QD")

	// ...
}
```
