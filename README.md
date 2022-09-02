# About

[![Go Report Card](https://goreportcard.com/badge/github.com/bshore/go-hirez)](https://goreportcard.com/report/github.com/bshore/go-hirez)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gomods/athens.svg)](https://github.com/bshore/go-hirez)
[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/bshore/go-hirez)](https://pkg.go.dev/mod/github.com/bshore/go-hirez)

`go-hirez` is a Golang API Library for the Hi-Rez (Smite) developer API [PDF reference](https://docs.google.com/document/d/1OFS-3ocSx-1Rvg4afAnEHlT3917MAK_6eJTR6rzr-BM/edit)

Add it with `go get`
```bash
go get github.com/bshore/go-hirez
```


## Getting Started with Mocks

`go-hirez` offers a mocked version of the API that will generate response output.
You can use the mocked API to test functionality without adding to your daily API limit.

_Note: The only method that isn't mocked is `GetOrganizedMatchDetailsBatch`_

```go
package main

import (
  mock "github.com/bshore/go-hirez/hirezapi_mock"
  "github.com/bshore/go-hirez/models"
)

func main() {
  client, err := mock.New("1234", "5678", models.URLSmitePC, models.ResponseTypeJSON)
  client.StartSession()

  client, err := mock.NewWithSession("1234", "5678", models.URLSmitePC, models.ResponseTypeJSON)
}
```

## Getting Started For Real

```go
package main

import (
  "github.com/bshore/go-hirez/hirezapi"
  "github.com/bshore/go-hirez/models"
)

func main() {
  // Recommend storing these as environment variables or secrets
  devID := "yourDevID" // os.Getenv("MY_SECRET_DEV_ID")
  authKey := "yourAuthKey" // os.Getenv("MY_SECRET_AUTH_KEY")

  // NewWithSession is like New() but it also tests connectivity with `Ping()` and calls `CreateSession()`
  // for you. This could be useful if you intend to query the API on some sort of schedule.
  client, err := hirezapi.NewWithSession(devID, authKey, models.URLSmitePC, models.ResponseTypeJSON)

  // New initializes a HiRezAPI instance with devID, auth key, url, and response type.
  // Note: You will need to call `client.CreateSession()` on your own
  client, err := hirezapi.New(devID, authKey, models.URLSmitePC, models.ResponseTypeJSON)

  // Direct struct initialization is supported, though the constructor methods are recommended since
  // they avoid potential fail cases like typos in the BasePath or an unsuppored RespType
  // Note: You will need to call `client.CreateSession()` on your own
  client := &hirezapi.APIClient{
    DeveloperID: devID,
    AuthKey: authKey,
    BasePath: "https://api.smitegame.com/smiteapi.svc",
    RespType: "json",
  }
}
```

## Notes

- Written with JSON responses in mind, the XML response seems a little more involved, so XML isn't supported at this time.
- Written with SMITE in mind, the Paladins endpoints are present but have not been fully tested so there are no guarantees.
