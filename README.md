[![Codeship Status for areatech/idealpostcodes-sdk](https://codeship.com/projects/d170a120-8f6a-0133-90ab-124821b463ac/status?branch=master)](https://codeship.com/projects/124189)

# Postcodes SDK

SDK for [IdealPostcodes.co.uk](https://ideal-postcodes.co.uk) API. This library will first try to use the open data from [Postcodes.io](http://postcodes.io/) and fallback to [IdealPostcodes.co.uk](https://ideal-postcodes.co.uk).

Example usage:

```go
import(
  "net/http"

  postcodes "github.com/areatech/postcodes-sdk"
)

// ...

client := postcodes.NewClient(
  postcodes.Endpoint,
  "your_api_key",
  new(http.Client),
)

addresses, err := client.GetPostcode("ID11QD")
```
