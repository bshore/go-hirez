package hirezapi

import (
	"io/ioutil"
	"os"
)

// GetChampionRanks returns the rank and worshipper values for each Champion a player has played.
func (a *APIClient) GetChampionRanks() error {
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

// GetChampions returns all Champions and their various attributes.
func (a *APIClient) GetChampions() error {
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

// GetChampionLeaderboard returns the current season's leaderboard for a champion/queue. [ Only queue 428]
func (a *APIClient) GetChampionLeaderboard() error {
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

// GetChampionSkins
func (a *APIClient) GetChampionSkins() error {
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

// GetPlayerIDInfoForXBOXAndSwitch returns all PlayerID data associated with the playerName
func (a *APIClient) GetPlayerIDInfoForXBOXAndSwitch() error {
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

// GetPlayerLoadouts returns deck loadouts per Champion
func (a *APIClient) GetPlayerLoadouts() error {
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
