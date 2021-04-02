package hirezapi

import (
	"io/ioutil"
	"os"
)

// GetESportsProLeagueDetails returns matchup info for each matchup of the current season. match_status = 1 - scheduled, 2 - in progress, 3 - complete
func (a *APIClient) GetESportsProLeagueDetails() error {
	resp, err := a.makeRequest("testsession", "")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = os.WriteFile("output.json", body, 0644)
	return err
}

// GetGodLeaderboard returns the current season's leaderboard for a god/queue. [SmiteAPI: only queues 440, 450, 451 apply]
func (a *APIClient) GetGodLeaderboard() error {
	resp, err := a.makeRequest("testsession", "")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = os.WriteFile("output.json", body, 0644)
	return err
}

// GetLeagueLeaderboard returns the top players for a particular league.
func (a *APIClient) GetLeagueLeaderboard() error {
	resp, err := a.makeRequest("testsession", "")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = os.WriteFile("output.json", body, 0644)
	return err
}

// GetLeageSeasons returns a list of seasons for a match queue.
func (a *APIClient) GetLeagueSeasons() error {
	resp, err := a.makeRequest("testsession", "")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = os.WriteFile("output.json", body, 0644)
	return err
}

// GetMOTD returns information about the 20 most recent Match-of-the-Days.
func (a *APIClient) GetMOTD() error {
	resp, err := a.makeRequest("testsession", "")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = os.WriteFile("output.json", body, 0644)
	return err
}

// GetTopMatches returns the 50 most watched / recorded matches.
func (a *APIClient) GetTopMatches() error {
	resp, err := a.makeRequest("testsession", "")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = os.WriteFile("output.json", body, 0644)
	return err
}

// GetPatchInfo returns information about the currently deployed patch.
func (a *APIClient) GetPatchInfo() error {
	resp, err := a.makeRequest("testsession", "")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = os.WriteFile("output.json", body, 0644)
	return err
}
