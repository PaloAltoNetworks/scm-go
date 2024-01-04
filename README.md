# scm-go
SDK for interacting with Strata Cloud Manager.

[![GoDoc](https://godoc.org/github.com/PaloAltoNetworks/scm-go?status.svg)](https://godoc.org/github.com/PaloAltoNetworks/scm-go)


## Using scm-go

To start, create a client connection, then invoke `Setup()`, and retrieve the JWT with `RefreshJwt()`.  JWTs expire after some time, but the SDK will catch the failed auth and automatically refresh the JWT when a 401 is returned from the API:

```go
package main

import (
    "context"
    "log"

    "github.com/paloaltonetworks/scm-go"
)

func main() {
    var err error
    ctx := context.TODO()

    con := scm.Client{
        ClientId: "1234-56-78-9",
        ClientSecret: "a-b-c-d",
        Scope: "tsg_id:123456789",
        CheckEnvironment: true,
    }

    if err = con.Setup(); err != nil {
        log.Fatal(err)
    } else if err = con.RefreshJwt(ctx); err != nil {
        log.Fatal(err)
    }

    log.Printf("Authenticated ok")
}
```
