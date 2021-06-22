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

type APIClient struct {
	DeveloperID  string
	AuthKey      string
	BasePath     string
	RespType     string
	SessionID    string
	SessionStamp string
	Logger 			 *log.Logger
}

// New initializes a HiRezAPI instance with devID, auth key, url, and response type.
func New(devID, key string, url models.MustBeURL, respType models.MustBeResponseType) (hirezapi.HiRezAPI, error) {
	if devID == "" {
		return nil, errors.New(`must provide a developerID (eg, 1004)`)
	}
	if key == "" {
		return nil, errors.New(`must provide an auth key (eg, 23DF3C7E9BD14D84BF892AD206B6755C)`)
	}
	api := &APIClient{
		BasePath:    url.String(),
		RespType:    respType.String(),
		DeveloperID: devID,
		AuthKey:     key,
		Logger:      log.Default(),
	}
	return api, nil
}

// NewWithSession is like New() but it also tests connectivity with Ping and
// initializes a session for you.
func NewWithSession(devID, key string, url models.URL, respType models.MustBeResponseType) (hirezapi.HiRezAPI, error) {
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

func (a *APIClient) NoLogging() {
	a.Logger = nil
}

func (a *APIClient) makeRequest(methodName, path string, desiredOutput interface{}) ([]byte, error) {
	signature, timestamp := a.generateSignature(methodName)
	apiURL := fmt.Sprintf("%s/%s%s/%s/%s/%s/%s", a.BasePath, methodName, a.RespType, a.DeveloperID, signature, a.SessionID, timestamp)
	if path != "" {
		apiURL = fmt.Sprintf("%s/%s", apiURL, path)
	}
	if a.Logger != nil {
		a.Logger.Printf("Mocked Request: %s\n", apiURL)
	}
	a.Logger.Printf("Calling GenerateDesiredOutput with: %#v\n", desiredOutput)
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
	if a.RespType == models.ResponseTypeXML.String() {
		return xml.Unmarshal(b, v)
	}
	return json.Unmarshal(b, v)
}

func FormatTime(t time.Time) string {
	return t.Format(models.TimeFormat)
}