package hirezapi

import (
	"fmt"
	"io/ioutil"

	"github.com/bshore/go-hirez/models"
)

// GetPlayer returns league and other high level data for a particular player.
func (a *APIClient) GetPlayer(player string) ([]models.Player, error) {
	resp, err := a.makeRequest("getplayer", player)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []models.Player
	err = a.unmarshalResponse(body, &output)
	return output, err
}

// GetFriends returns Smite Usernames of each of the player's friends. [PC Only]
func (a *APIClient) GetFriends(player string) ([]models.Friend, error) {
	if IsConsolePath(a.BasePath) {
		return nil, fmt.Errorf("GetFriends() %s", isConsoleErrMsg)
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
	var output []models.Friend
	err = a.unmarshalResponse(body, &output)
	return output, err
}

// GetGodRanks returns the rank and worshipper values for each God a player has played.
func (a *APIClient) GetGodRanks(player string) ([]models.GodRank, error) {
	resp, err := a.makeRequest("getgodranks", player)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []models.GodRank
	err = a.unmarshalResponse(body, &output)
	return output, err
}

// GetMatchHistory returns a list of recent matches and high level match statistics for a particular player
func (a *APIClient) GetMatchHistory(player string) ([]models.Match, error) {
	resp, err := a.makeRequest("getmatchhistory", player)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []models.Match
	err = a.unmarshalResponse(body, &output)
	return output, err
}

// GetPlayerStatus returns a player status.
func (a *APIClient) GetPlayerStatus(player string) ([]models.PlayerStatus, error) {
	resp, err := a.makeRequest("getplayerstatus", player)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []models.PlayerStatus
	err = a.unmarshalResponse(body, &output)
	return output, err
}

// GetPlayerAchievements returns select achievement totals.
func (a *APIClient) GetPlayerAchievements(playerID string) (*models.PlayerAchievements, error) {
	if IsNotSmitePath(a.BasePath) {
		return nil, fmt.Errorf("GetPlayerAchievements() %s", notSmiteErrMsg)
	}
	resp, err := a.makeRequest("getplayerachievements", playerID)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output *models.PlayerAchievements
	err = a.unmarshalResponse(body, &output)
	return output, err
}

// GetTeamDetails returns the number of players and other high level details for a particular Clan.
func (a *APIClient) GetTeamDetails(clanID string) ([]models.TeamDetail, error) {
	resp, err := a.makeRequest("getteamdetails", clanID)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []models.TeamDetail
	err = a.unmarshalResponse(body, &output)
	return output, err
}

// GetTeamPlayers returns a list of players for a particular Clan.
func (a *APIClient) GetTeamPlayers(clanID string) ([]models.TeamPlayer, error) {
	resp, err := a.makeRequest("getteamplayers", clanID)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []models.TeamPlayer
	err = a.unmarshalResponse(body, &output)
	return output, err
}

// SearchTeams returns high level info for Clan names containing search term.
func (a *APIClient) SearchTeams(searchTeam string) ([]models.TeamDetail, error) {
	if IsNotSmitePath(a.BasePath) {
		return nil, fmt.Errorf("SearchTeams() %s", notSmiteErrMsg)
	}
	resp, err := a.makeRequest("searchteams", searchTeam)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []models.TeamDetail
	err = a.unmarshalResponse(body, &output)
	return output, err
}