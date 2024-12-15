# Golang Tiingo Client
Full-featured [Tiingo](https://www.tiingo.com/) client that offers CSV and JSON unmarshalling from Tiingo to Golang types. 

## Installation 
```shell
go get github.com/the-trader-dev/tiin-go
```

## Usage
Maximum flexibility is offered with three primary ways to use tiin-go: 
1. As a Tiingo frontend client
2. As a url builder
3. Steal the types

### Tiingo Frontend Client
Using tiin-go as a frontend client for the Tiingo api is the simplest way to use
this package. It offers a few advantages over just using the query building & type
capabilities:
1. Centralized rate limiting
2. Automatic authentication
3. Request logging
4. Automatic response body lifecycle management

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "log/slog"
    "net/http"
    "os"

    tiingo "github.com/the-trader-dev/tiin-go"
    "golang.org/x/time/rate"
)

func main() { 
    // Initialize client
    client := tiingo.NewClient(os.Getenv("YOUR_TIINGO_TOKEN"))

    // You can optionally set a rate limiter, enable logging, and change the
    // default http client
    client.RateLimiter = rate.NewLimiter(10, 1)
    client.Logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
    client.HttpClient = &http.Client{}

    // Make request
    rawBytes, err := client.DefaultEodMetadata(context.Background(), "AAPL")
    if err != nil {
        panic(err)
    }

    // Unmarshal the response
    var metadata tiingo.EodMetadata
    if err = json.Unmarshal(rawBytes, &metadata); err != nil {
        panic(err)
    }

    fmt.Println("Apple's metadata:", metadata)
}
```

### URL Builder
Using tiin-go purely as a url builder offers the most control. You pass in the
wanted parameters into the given url function and the valid corresponding url 
is returned. You then do the actual request however you want. Once the request 
is made, you can then unmarshal the response into one of the predefined types. 

NOTE: no token query param is added, you must handle authentication yourself

```go
package main

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"

    tiingo "github.com/the-trader-dev/tiin-go"
)

func main() {
    // Build url
    url := tiingo.EodMetadataUrl("AAPL", tiingo.JSON)

    // Build request
    req, err := http.NewRequest(http.MethodGet, url, nil)
    if err != nil {
        panic(err)
    }

    // Add auth
    req.Header.Set("Authorization", "Token {your_token}")

    // Make request
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        panic(err)
    }
    defer func() {
        if err = resp.Body.Close(); err != nil {
            panic(err)
        }
    }()

    // Read body
    rawBytes, err := io.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }

    // Unmarshal
    var metadata tiingo.EodMetadata
    if err = json.Unmarshal(rawBytes, &metadata); err != nil {
        panic(err)
    }

    fmt.Println("Apple's metadata:", metadata)
}
```

### Steal the types
Implemented endpoints have a one-to-one matching of Tiingo types to Golang types. 
You can handle the entire request lifecycle yourself and steal the types to make 
marshalling & unmarshalling easier.

```go
package main

import (
    "encoding/json"
    "fmt"

    "github.com/gocarina/gocsv"
    tiingo "github.com/the-trader-dev/tiin-go"
)

func main() {
    // Raw bytes from some endpoint you requested
    var jsonBytes []byte
    var csvBytes []byte

    // Unmarshal
    var jsonPrices []tiingo.EodPrice
    if err := json.Unmarshal(jsonBytes, &jsonPrices); err != nil {
        panic(err)
    }
    var csvPrices []tiingo.EodPrice
    if err := gocsv.UnmarshalBytes(csvBytes, &csvPrices); err != nil {
        panic(err)
    }

    fmt.Println("prices from json:", jsonPrices)
    fmt.Println("prices from csv:", csvPrices)
}
```

## Types 
All implemented endpoints have corresponding Golang types that allow for automatic
marshalling and unmarshalling both with CSV & JSON responses.

#### JSON
Any package that knows how to read json struct tags should be able to marshal/unmarshal
successfully. However, it has only been tested with the standard library encoding/json 
package. 
```go
err := json.Unmarshal(rawJsonBytes, &TiingoGolangType)
```

#### CSV
Any package that knows how to read csv struct tags should be able to marshal/unmarshal
successfully. However, it has only been tested with & internally uses with [gocsv](https://github.com/gocarina/gocsv).
Initially, the plan was to implement custom csv marshalling/unmarshalling to minimize 
dependencies. However, the vast quantity of types made compromising with an external
dependency a lot more attractive in terms of the overall maintenance overhead. 
```go
err := gocsv.UnmarshaBytes(rawCsvBytes, &TiingoGolangType)
```

## API Surface
### Complete
The following Tiingo endpoints have been implemented: 
- End-of-Day
- IEX
- Fundamentals
- Search

### Incomplete
The following Tiingo endpoints are not yet implemented:
- Crypto
- Forex
- Fund Fees
- Dividends
- Splits

## Contributions
Contributions are welcome! 

All I ask is that if you implement any of the needed endpoints, follow the same pattern 
as the completed ones to keep the api consistent (url builder, client method, valid csv/json parsing)