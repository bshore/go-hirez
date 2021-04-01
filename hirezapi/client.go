package hirezapi

import (
	"crypto/md5"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// URL determinmes the baseURL to use when calling Hi-Rez API endpoints
type URL int

const (
	URLSmitePC URL = iota
	URLSmiteXBOX
	URLSmitePS4
	URLPaladinsPC
	URLPaladinsXBOX
	URLPaladinsPS4
)

func (u URL) String() string {
	return [...]string{
		"http://api.smitegame.com/smiteapi.svc",
		"http://api.xbox.smitegame.com/smiteapi.svc",
		"http://api.ps4.smitegame.com/smiteapi.svc",
		"http://api.paladins.com/paladinsapi.svc",
		"http://api.xbox.paladins.com/paladinsapi.svc",
		"http://api.ps4.paladins.com/paladinsapi.svc",
	}[u]
}

// IsURL exists to satisfy the MustBeURL interface
func (u URL) IsURL() bool {
	return true
}

// MustBeURL exists to force the constructors to use our URL type
type MustBeURL interface {
	String() string
	IsURL() bool
}

// ResponseType determines the type of the Response body: JSON or XML
type ResponseType int

const (
	ResponseTypeJSON ResponseType = iota
	ResponseTypeXML
)

func (r ResponseType) String() string {
	return [...]string{"json", "xml"}[r]
}

// IsResponseType exists to satisfy the MustBeResponseType interface
func (r ResponseType) IsResponseType() bool {
	return true
}

// MustBeResponseType exists to force the constructors to use our ResponseType type
type MustBeResponseType interface {
	String() string
	IsResponseType() bool
}

const (
	dateFormat string = "20060102" // yyyyMMdd
	timeFormat string = "20060102150405" // yyyyMMddHHmmss
)

// APIClient is the implementation of the HiRezAPI interface.
type APIClient struct {
	DeveloperID string
	AuthKey string
	BasePath string
	RespType string
	SessionID string
	SessionStamp string
}

// New initializes a HiRezAPI instance with devID, auth key, url, and response type.
func New(devID, key string, url MustBeURL, respType MustBeResponseType) (HiRezAPI, error) {
	if devID == "" {
		return nil, errors.New(`must provide a developerID (eg, 1004)`)
	}
	if key == "" {
		return nil, errors.New(`must provide an auth key (eg, 23DF3C7E9BD14D84BF892AD206B6755C)`)
	}
	api := &APIClient{
		BasePath: url.String(),
		RespType: respType.String(),
		DeveloperID: devID,
		AuthKey: key,
	}
	return api, nil
}

// NewWithSession is like New() but it also tests connectivity with Ping and
// initializes a session for you.
func NewWithSession(devID, key string, url URL, respType MustBeResponseType) (HiRezAPI, error) {
	api, err := New(devID, key, url, respType)
	if err != nil {
		return nil, err
	}
	err = api.Ping()
	if err != nil {
		return nil, err
	}
	err = api.CreateSession()
	if err != nil {
		return nil, err
	}
	return api, nil
}

func (a *APIClient) makeRequest(methodName, path string) (*http.Response, error) {
	signature, timestamp := a.GenerateSignature(methodName)
	apiURL := fmt.Sprintf("%s/%s%s/%s/%s/%s/%s", a.BasePath, methodName, a.RespType, a.DeveloperID, signature, a.SessionID, timestamp)
	if path != "" {
		apiURL = fmt.Sprintf("%s/%s", apiURL, path)
	}
	return http.Get(apiURL)
}

// generateSignature takes in the requested methodName and generates an md5 hashed signature for sending a request
func (a *APIClient) GenerateSignature(methodName string) (string, string) {
	utcNow := time.Now().UTC().Format(timeFormat)
	sigStr := fmt.Sprintf("%s%s%s%s", a.DeveloperID, methodName, a.AuthKey, utcNow)
	bs := []byte(sigStr)
	return fmt.Sprintf("%x", md5.Sum(bs)), utcNow
}