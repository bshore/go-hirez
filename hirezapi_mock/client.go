package mock

import (
	"crypto/md5"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/bshore/go-hirez/hirezapi"
	"github.com/bshore/go-hirez/models"
	"github.com/bshore/go-hirez/utils"
)

// APIClient stores data needed to interface with the HiRez API and implements the HiRezAPI interface methods.
type APIClient struct {
	DeveloperID  string
	AuthKey      string
	BasePath     string
	RespType     string
	SessionID    string
	SessionStamp string
}

// New initializes a HiRezAPI instance with devID, auth key, url, and response type.
func New(devID, key, url, respType string) (hirezapi.HiRezAPI, error) {
	if devID == "" {
		return nil, errors.New(`must provide a developerID (eg, 1004)`)
	}
	if key == "" {
		return nil, errors.New(`must provide an auth key (eg, 23DF3C7E9BD14D84BF892AD206B6755C)`)
	}
	return &APIClient{
		BasePath:    url,
		RespType:    respType,
		DeveloperID: devID,
		AuthKey:     key,
	}, nil
}

// NewWithSession is like New() but it also tests connectivity with Ping and
// initializes a session for you.
func NewWithSession(devID, key, url, respType string) (hirezapi.HiRezAPI, error) {
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

func (a *APIClient) makeRequest(methodName, path string, desiredOutput interface{}) ([]byte, error) {
	signature, timestamp := a.generateSignature(methodName)
	apiURL := fmt.Sprintf("%s/%s%s/%s/%s/%s/%s", a.BasePath, methodName, a.RespType, a.DeveloperID, signature, a.SessionID, timestamp)
	if path != "" {
		apiURL = fmt.Sprintf("%s/%s", apiURL, path)
	}
	log.Printf("Mocked Request: %s\n", apiURL)
	return utils.GenerateDesiredOutput(desiredOutput)
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
