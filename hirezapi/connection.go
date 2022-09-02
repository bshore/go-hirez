package hirezapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/bshore/go-hirez/models"
)

// Ping is a quick way of validating access to the Hi-Rez API.
func (a *APIClient) Ping() error {
	url := fmt.Sprintf("%s/%s%s", a.BasePath, "ping", a.RespType)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error: Ping() did not return status 200: %d, Response: %s", resp.StatusCode, resp.Body)
	}
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
	resp, err := http.Get(path)
	if err != nil {
		return fmt.Errorf("error creating session: %v", err)
	}
	defer resp.Body.Close()
	sess := &models.Session{}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %v", err)
	}
	err = json.Unmarshal(body, sess)
	if err != nil {
		return fmt.Errorf("error unmarshaling response: %v", err)
	}
	if sess.RetMsg != "Approved" {
		return fmt.Errorf("error creating session: %v", sess.RetMsg)
	}
	a.SessionID = sess.SessionID
	a.SessionStamp = sess.Timestamp
	return nil
}

// TestSession is a means of validating that a session is established.
func (a *APIClient) TestSession() (string, error) {
	resp, err := a.makeRequest("testsession", "")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// GetHiRezServerStatus returns UP/DOWN status for the primary game/platform environments. Data is cached once a minute.
func (a *APIClient) GetHiRezServerStatus() ([]models.HiRezServerStatus, error) {
	resp, err := a.makeRequest("gethirezserverstatus", "")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []models.HiRezServerStatus
	err = a.unmarshalResponse(body, &output)
	return output, err
}

// GetDataUsed returns API Developer daily usage limits and the current status against those limits.
func (a *APIClient) GetDataUsed() ([]models.DataUsed, error) {
	resp, err := a.makeRequest("getdataused", "")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []models.DataUsed
	err = a.unmarshalResponse(body, &output)
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
