# IdealPostcodes SDK

SDK for [IdealPostcodes](https://ideal-postcodes.co.uk)] API.

Example usage:

```go
import(
  "net/http"

  idealpostcodes "github.com/areatech/idealpostcodes-sdk"
)

// ...

client := idealpostcodes.NewClient(
  "https://api.ideal-postcodes.co.uk/v1",
  "your_api_key",
  new(http.Client),
)

addresses, err := client.GetPostcode("ID11QD")
```
