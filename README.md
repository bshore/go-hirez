# About

`go-hirez` is a Golang API Library for the Hi-Rez (Smite) developer API [PDF reference](https://web2.hirez.com/hirez-studios/legal/smite-api-developer-guide.pdf)

## Getting Started

```go
package main

import "github.com/bshore/go-hirez/hirezapi"

func main() {
  // Recommend storing these as environment variables or secrets
  devID := "yourDevID" // os.Getenv("MY_SECRET_DEV_ID")
  authKey := "yourAuthKey" // os.Getenv("MY_SECRET_AUTH_KEY")

  // New initializes a HiRezAPI instance with devID, auth key, url, and response type.
  client, err := hirezapi.New(devID, authKey, hirezapi.URLSmitePC, hirezapi.ResponseTypeJSON)

  // NewWithSession is like New() but it also tests connectivity with Ping and initializes a
  // session for you. This could be used if you intend to query the API on some sort of schedule.
  client, err := hirezapi.NewWithSession(devID, authKey, hirezapi.URLSmitePC, hirezapi.ResponseTypeJSON)

  // Direct struct initialization is supported, though the constructor methods are recommended since they avoid
  // potential fail cases like typos in the BasePath or an unsuppored RespType
  client := &hirezapi.APIClient{
    DeveloperID: devID,
    AuthKey: authKey,
    BasePath: "http://api.smitegame.com/smiteapi.svc",
    RespType: "json",
  }
}
```

### **NOTE:** By using `struct` initialization or `New()`, you will need to call `client.CreateSession()` on your own once you are ready to begin making API calls.
