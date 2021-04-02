package hirezapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/bshore/go-hirez/models"
)

// Ping is a quick way of validating access to the Hi-Rez API.
func (a *APIClient) Ping() error {
	method := "ping"
	url := fmt.Sprintf("%s/%s%s", a.BasePath, method, a.RespType)
	resp, err := http.Get(url)
	if resp.StatusCode != http.StatusOK {
		return err
	}
	return nil
}

// CreateSession is a required step to Authenticate the developerId/signature for further API use.
func (a *APIClient) CreateSession() error {
	method := "createsession"
	sig, stamp := a.GenerateSignature(method)
	url := fmt.Sprintf("%s/%s%s/%s/%s/%s", a.BasePath, method, a.RespType, a.DeveloperID, sig, stamp)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	sess := &models.SessionResponse{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, sess)
	if err != nil {
		return err
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
func (a *APIClient) GetHiRezServerStatus() ([]models.HiRezServerStatusResponse, error) {
	resp, err := a.makeRequest("gethirezserverstatus", "")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []models.HiRezServerStatusResponse
	err = json.Unmarshal(body, &output)
	return output, err
}

// GetDataUsed returns API Developer daily usage limits and the current status against those limits.
func (a *APIClient) GetDataUsed() ([]models.DataUsedResponse, error) {
	resp, err := a.makeRequest("getdataused", "")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []models.DataUsedResponse
	err = json.Unmarshal(body, &output)
	return output, err
}