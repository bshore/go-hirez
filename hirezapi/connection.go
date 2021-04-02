package hirezapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/bshore/go-hirez/models"
)

// Ping is a quick way of validating access to the Hi-Rez API.
func (a *APIClient) Ping() error {
	url := fmt.Sprintf("%s/%s%s", a.BasePath, "ping", a.RespType)
	resp, err := http.Get(url)
	if resp.StatusCode != http.StatusOK {
		return err
	}
	return nil
}

// CreateSession is a required step to Authenticate the developerId/signature for further API use.
func (a *APIClient) CreateSession() error {
	sig, stamp := a.generateSignature("createsession")
	url := fmt.Sprintf(
		"%s/%s%s/%s/%s/%s",
		URLSmitePC.String(), // Smite PC is the session management endpoint.
		"createsession",
		ResponseTypeJSON.String(), // Create session using json
		a.DeveloperID,
		sig,
		stamp,
	)
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error creating session: %v", err)
	}
	defer resp.Body.Close()
	sess := &models.Session{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %v", err)
	}
	err = json.Unmarshal(body, sess)
	if err != nil {
		log.Print(string(body))
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
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// GetHirezServerStatus returns UP/DOWN status for the primary game/platform environments. Data is cached once a minute.
func (a *APIClient) GetHiRezServerStatus() ([]models.HiRezServerStatus, error) {
	resp, err := a.makeRequest("gethirezserverstatus", "")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
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
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []models.DataUsed
	err = a.unmarshalResponse(body, &output)
	return output, err
}

// ChangeBasePath modifies the base path if you want to query a different platform
func (a *APIClient) ChangeBasePath(url MustBeURL) {
	a.BasePath = url.String()
}

// ChangeResponseType modifies the response type if you want to switch between JSON and XML
func (a *APIClient) ChangeResponseType(respType MustBeResponseType) {
	a.RespType = respType.String()
}