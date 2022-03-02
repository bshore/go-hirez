package mock

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/bshore/go-hirez/models"
	"github.com/bshore/go-hirez/utils"
)

// Ping is a quick way of validating access to the Hi-Rez API.
func (a *APIClient) Ping() error {
	log.Println("Ping() has been called")
	return nil
}

// CreateSession is a required step to Authenticate the developerId/signature for further API use.
func (a *APIClient) CreateSession() error {
	sig, stamp := a.generateSignature("createsession")
	path := fmt.Sprintf(
		"%s/%s%s/%s/%s/%s",
		a.BasePath,
		"createsession",
		models.ResponseTypeJSON,
		a.DeveloperID,
		sig,
		stamp,
	)
	log.Printf("Mocked Request: %s\n", path)
	var sess *models.Session
	body, err := utils.GenerateDesiredOutput(models.Session{})
	if err != nil {
		return fmt.Errorf("error generating session: %v", err)
	}
	err = json.Unmarshal(body, &sess)
	if err != nil {
		return fmt.Errorf("error unmarshaling response: %v", err)
	}
	a.SessionID = sess.SessionID
	a.SessionStamp = sess.Timestamp
	return nil
}

// TestSession is a means of validating that a session is established.
func (a *APIClient) TestSession() (string, error) {
	if a.SessionID != "" && a.SessionStamp != "" {
		return "session ok", nil
	}
	return "", fmt.Errorf("session not initialized")
}

// GetHirezServerStatus returns UP/DOWN status for the primary game/platform environments. Data is cached once a minute.
func (a *APIClient) GetHiRezServerStatus() ([]models.HiRezServerStatus, error) {
	resp, err := a.makeRequest("gethirezserverstatus", "", []models.HiRezServerStatus{})
	if err != nil {
		return nil, err
	}
	var output []models.HiRezServerStatus
	err = a.unmarshalResponse(resp, &output)
	return output, err
}

// GetDataUsed returns API Developer daily usage limits and the current status against those limits.
func (a *APIClient) GetDataUsed() ([]models.DataUsed, error) {
	resp, err := a.makeRequest("getdataused", "", []models.DataUsed{})
	if err != nil {
		return nil, err
	}
	var output []models.DataUsed
	err = a.unmarshalResponse(resp, &output)
	return output, err
}

// ChangeBasePath modifies the base path if you want to query a different platform.
func (a *APIClient) ChangeBasePath(url string) {
	a.BasePath = url
	a.CreateSession()
}

// ChangeResponseType modifies the response type if you want to switch between JSON and XML
func (a *APIClient) ChangeResponseType(respType string) {
	a.RespType = respType
}
