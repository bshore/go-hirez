# go-hirez - A Golang Hi-Rez API Library

Golang API Library for the Hi-Rez (Smite) developer API [PDF](https://web2.hirez.com/hirez-studios/legal/smite-api-developer-guide.pdf)

Sign up for an account [here](https://fs12.formsite.com/HiRez/form48/secure_index.html)

## Getting Started

```go
package main

import "github.com/bshore/go-hirez/hirezapi"

func main() {
  // Recommend storing these as environment variables or secrets
  devID := "1004"
  authKey := "23DF3C7E9BD14D84BF892AD206B6755C"

  // Direct struct initialization, while not exactly recommended, it is supported.
  client := &hirezapi.APIClient{
    DeveloperID: devID,
    AuthKey: authKey,
    BasePath: "http://api.smitegame.com/smiteapi.svc",
    RespType: "json",
  }

  // New initializes a HiRezAPI instance with devID, auth key, platform type, and response type.
  client, err := hirezapi.New(devID, authKey, hirezapi.URLSmitePC, hirezapi.ResponseTypeJSON)

  // ^ NOTE: Using `struct` initialization and calling `New()` both require you to call
  // `client.CreateSession()` manually, when ready to start making API calls.

  // NewWithSession is like New() but it also tests connectivity with Ping and initializes a
  // session for you. This could be used if you intend to query the API on some sort of schedule.
  client, err := hirezapi.NewWithSession(devID, authKey, hirezapi.URLSmitePC, hirezapi.ResponseTypeJSON)
}
```
