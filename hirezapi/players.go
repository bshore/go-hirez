package hirezapi

import (
	"fmt"
	"io/ioutil"
)

// GetFriends returns Smite Usernames of each of the player's friends. [PC Only]
func (a *APIClient) GetFriends() error {
	resp, err := a.makeRequest("testsession", "")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("\n%s\n", body)
	return nil

}

// GetGodRanks returns the rank and worshipper values for each God a player has played.
func (a *APIClient) GetGodRanks() error {
	resp, err := a.makeRequest("testsession", "")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("\n%s\n", body)
	return nil

}

// GetMatchHistory returns a list of recent matches and high level match statistics for a particular player
func (a *APIClient) GetMatchHistory() error {
	resp, err := a.makeRequest("testsession", "")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("\n%s\n", body)
	return nil

}

// GetPlayer returns league and other high level data for a particular player.
func (a *APIClient) GetPlayer() error {
	resp, err := a.makeRequest("testsession", "")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("\n%s\n", body)
	return nil

}

// GetPlayerStatus returns a player status. 0 - offline, 1 - lobby, 2 - god select, 3 - in game, 4 - online, 5 - unknown
func (a *APIClient) GetPlayerStatus() error {
	resp, err := a.makeRequest("testsession", "")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("\n%s\n", body)
	return nil

}

// GetTeamDetails returns the number of players and other high level details for a particular Clan.
func (a *APIClient) GetTeamDetails() error {
	resp, err := a.makeRequest("testsession", "")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("\n%s\n", body)
	return nil

}

// GetTeamPlayers returns a list of players for a particular Clan.
func (a *APIClient) GetTeamPlayers() error {
	resp, err := a.makeRequest("testsession", "")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("\n%s\n", body)
	return nil

}

// SearchTeams returns high level info for Clan names containing search term.
func (a *APIClient) SearchTeams() error {
	resp, err := a.makeRequest("testsession", "")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("\n%s\n", body)
	return nil

}

// GetPlayerAchievements returns select achievement totals.
func (a *APIClient) GetPlayerAchievements() error {
	resp, err := a.makeRequest("testsession", "")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("\n%s\n", body)
	return nil

}
