[![Codeship Status for areatech/idealpostcodes-sdk](https://codeship.com/projects/d170a120-8f6a-0133-90ab-124821b463ac/status?branch=master)](https://codeship.com/projects/124189)

# IdealPostcodes SDK

SDK for [IdealPostcodes](https://ideal-postcodes.co.uk) API.

Example usage:

```go
import(
  "net/http"

  postcodes "github.com/areatech/postcodes-sdk"
)

// ...

client := idealpostcodes.NewClient(
  postcodes.Endpoint,
  "your_api_key",
  new(http.Client),
)

addresses, err := client.GetPostcode("ID11QD")
```
