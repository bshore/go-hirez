package hirezapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/bshore/go-hirez/models"
)

// GetPlayer returns league and other high level data for a particular player.
func (a *APIClient) GetPlayer(player string) ([]models.PlayerResponse, error) {
	resp, err := a.makeRequest("getplayer", player)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []models.PlayerResponse
	err = json.Unmarshal(body, &output)
	return output, err
}

// GetFriends returns Smite Usernames of each of the player's friends. [PC Only]
func (a *APIClient) GetFriends(player string) ([]models.FriendsResponse, error) {
	if !IsPCOnly(a.BasePath) {
		return nil, fmt.Errorf("GetFriends() is only supported for PC requests")
	}
	resp, err := a.makeRequest("getfriends", player)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []models.FriendsResponse
	err = json.Unmarshal(body, &output)
	return output, err
}

// GetGodRanks returns the rank and worshipper values for each God a player has played.
func (a *APIClient) GetGodRanks(player string) ([]models.GodRanksResponse, error) {
	resp, err := a.makeRequest("getgodranks", player)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []models.GodRanksResponse
	err = json.Unmarshal(body, &output)
	return output, err
}

// GetMatchHistory returns a list of recent matches and high level match statistics for a particular player
func (a *APIClient) GetMatchHistory(player string) ([]models.MatchHistoryResponse, error) {
	resp, err := a.makeRequest("getmatchhistory", player)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []models.MatchHistoryResponse
	err = json.Unmarshal(body, &output)
	return output, err
}

// GetPlayerStatus returns a player status.
func (a *APIClient) GetPlayerStatus(player string) ([]models.PlayerStatusResponse, error) {
	resp, err := a.makeRequest("getplayerstatus", player)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []models.PlayerStatusResponse
	err = json.Unmarshal(body, &output)
	return output, err
}

// GetPlayerAchievements returns select achievement totals.
func (a *APIClient) GetPlayerAchievements(playerID string) error {
	resp, err := a.makeRequest("getplayerachievements", playerID)
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
func (a *APIClient) GetTeamDetails(clanID string) error {
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
func (a *APIClient) GetTeamPlayers(clanID string) error {
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
func (a *APIClient) SearchTeams(searchTeam string) error {
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