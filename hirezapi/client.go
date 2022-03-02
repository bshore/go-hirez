package hirezapi

import (
	"crypto/md5"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/bshore/go-hirez/models"
)

// APIClient is the implementation of the HiRezAPI interface.
type APIClient struct {
	DeveloperID  string
	AuthKey      string
	BasePath     string
	RespType     string
	SessionID    string
	SessionStamp string
}

// New initializes a HiRezAPI instance with devID, auth key, url, and response type.
func New(devID, key, url, respType string) (HiRezAPI, error) {
	if devID == "" {
		return nil, errors.New(`must provide a developerID (eg, 1004)`)
	}
	if key == "" {
		return nil, errors.New(`must provide an auth key (eg, 23DF3C7E9BD14D84BF892AD206B6755C)`)
	}
	api := &APIClient{
		BasePath:    url,
		RespType:    respType,
		DeveloperID: devID,
		AuthKey:     key,
	}
	return api, nil
}

// NewWithSession is like New() but it also tests connectivity with Ping and
// initializes a session for you.
func NewWithSession(devID, key, url, respType string) (HiRezAPI, error) {
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
	signature, timestamp := a.generateSignature(methodName)
	apiURL := fmt.Sprintf("%s/%s%s/%s/%s/%s/%s", a.BasePath, methodName, a.RespType, a.DeveloperID, signature, a.SessionID, timestamp)
	if path != "" {
		apiURL = fmt.Sprintf("%s/%s", apiURL, path)
	}
	return http.Get(apiURL)
}

// generateSignature takes in the requested methodName and generates an md5 hashed signature for sending a request
func (a *APIClient) generateSignature(methodName string) (string, string) {
	utcNow := time.Now().UTC().Format(models.TimeFormat)
	sigStr := fmt.Sprintf("%s%s%s%s", a.DeveloperID, methodName, a.AuthKey, utcNow)
	bs := []byte(sigStr)
	return fmt.Sprintf("%x", md5.Sum(bs)), utcNow
}

func (a *APIClient) unmarshalResponse(b []byte, v interface{}) error {
	if a.RespType == models.ResponseTypeXML {
		return xml.Unmarshal(b, v)
	}
	return json.Unmarshal(b, v)
}
