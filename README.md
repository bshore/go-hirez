# About

`go-hirez` is a Golang API Library for the Hi-Rez (Smite) developer API [PDF reference](https://docs.google.com/document/d/1OFS-3ocSx-1Rvg4afAnEHlT3917MAK_6eJTR6rzr-BM/edit)

## Getting Started

```go
package main

import "github.com/bshore/go-hirez/hirezapi"

func main() {
  // Recommend storing these as environment variables or secrets
  devID := "yourDevID" // os.Getenv("MY_SECRET_DEV_ID")
  authKey := "yourAuthKey" // os.Getenv("MY_SECRET_AUTH_KEY")

  // NewWithSession is like New() but it also tests connectivity with `Ping()` and calls `CreateSession()` for you. This could be useful if you intend to query the API on some sort of schedule.
  client, err := hirezapi.NewWithSession(devID, authKey, hirezapi.URLSmitePC, hirezapi.ResponseTypeJSON)

  // New initializes a HiRezAPI instance with devID, auth key, url, and response type.
  // Note: You will need to call `client.CreateSession()` on your own
  client, err := hirezapi.New(devID, authKey, hirezapi.URLSmitePC, hirezapi.ResponseTypeJSON)

  // Direct struct initialization is supported, though the constructor methods are recommended since
  // they avoid potential fail cases like typos in the BasePath or an unsuppored RespType
  // Note: You will need to call `client.CreateSession()` on your own
  client := &hirezapi.APIClient{
    DeveloperID: devID,
    AuthKey: authKey,
    BasePath: "http://api.smitegame.com/smiteapi.svc",
    RespType: "json",
  }
}
```

## Notes/TODOs

- Written with Smite & JSON responses in mind, the `xml:""` struct tags are currently missing and has not been tested.
- There's some inconsistent validation across methods regarding Smite/Paladins/Platform specific API calls, and I'd like to iron that out.
- The API endpoint for `GetLeagueLeaderboard` is giving me some trouble...
  - Seems to always return `[]` no matter what I pass it.
- The API endpoint for `GetPlayerIDInfoForXBOXAndSwitch` is giving me some trouble...
  - Trying to get a sample XBOX gamertag with `GetChampionLeaderboard` but the request `404`s when using the Paladins XBOX endpoint.
