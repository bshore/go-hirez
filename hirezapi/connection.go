package hirezapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

	sess := &sessionResponse{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, sess)
	if err != nil {
		return err
	}

	a.SessionID = sess.SessionID
	a.SessionStamp = sess.Timestamp
	return nil
}

// TestSession is a means of validating that a session is established.
func (a *APIClient) TestSession() {
	
}

// GetHirezServerStatus returns UP/DOWN status for the primary game/platform environments. Data is cached once a minute.
func (a *APIClient) GetHiRezServerStatus() {
	
}

// GetDataUsed returns API Developer daily usage limits and the current status against those limits.
func (a *APIClient) GetDataUsed() {
	
}

// ===== Types =====
type sessionResponse struct {
	RetMsg    string `json:"ret_msg"`
	SessionID string `json:"session_id"`
	Timestamp string `json:"timestmap"`
}