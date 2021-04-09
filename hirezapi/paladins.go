package hirezapi

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/bshore/go-hirez/models"
)

// GetPlayerBatch returns league and other high level data for a particular list of players. [20 max]
func (a *APIClient) GetPlayerBatch(playerIDs []string) ([]models.Player, error) {
	if !IsPaladinsPath(a.BasePath) {
		return nil, fmt.Errorf("GetPlayerBatch() %s", notPaladinsErrMsg)
	}
	resp, err := a.makeRequest("getplayerbatch", strings.Join(playerIDs, ","))
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

// GetChampionRanks returns the rank and worshipper values for each Champion a player has played.
func (a *APIClient) GetChampionRanks(player string) ([]models.ChampionRank, error) {
	if !IsPaladinsPath(a.BasePath) {
		return nil, fmt.Errorf("GetChampionRanks() %s", notPaladinsErrMsg)
	}
	resp, err := a.makeRequest("getchampionranks", player)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []models.ChampionRank
	err = a.unmarshalResponse(body, &output)
	return output, err
}

// GetChampions returns all Champions and their various attributes.
func (a *APIClient) GetChampions(langCode string) ([]models.Champion, error) {
	if !IsPaladinsPath(a.BasePath) {
		return nil, fmt.Errorf("GetChampions() %s", notPaladinsErrMsg)
	}
	if langCode == "" {
		langCode = English
	}
	resp, err := a.makeRequest("getchampions", langCode)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []models.Champion
	err = a.unmarshalResponse(body, &output)
	return output, err
}

// GetChampionLeaderboard returns the current season's leaderboard for a champion/queue. [ Only queue 428]
func (a *APIClient) GetChampionLeaderboard(champID string) ([]models.ChampionLeaderboardEntry, error) {
	if !IsPaladinsPath(a.BasePath) {
		return nil, fmt.Errorf("GetChampionLeaderboard() %s", notPaladinsErrMsg)
	}
	path := fmt.Sprintf("%s/428", champID)
	resp, err := a.makeRequest("getchampionleaderboard", path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []models.ChampionLeaderboardEntry
	err = a.unmarshalResponse(body, &output)
	return output, err
}

// GetChampionSkins returns a list of skins for a particular champion.
func (a *APIClient) GetChampionSkins(champID, langCode string) ([]models.ChampionSkin, error) {
	if !IsPaladinsPath(a.BasePath) {
		return nil, fmt.Errorf("GetChampionSkins() %s", notPaladinsErrMsg)
	}
	if langCode == "" {
		langCode = English
	}
	path := fmt.Sprintf("%s/%s", champID, langCode)
	resp, err := a.makeRequest("getchampionskins", path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []models.ChampionSkin
	err = a.unmarshalResponse(body, &output)
	return output, err
}

// GetPlayerIDInfoForXBOXAndSwitch returns all PlayerID data associated with the playerName
func (a *APIClient) GetPlayerIDInfoForXBOXAndSwitch(player string) ([]models.PlayerIDInfoForXBOXAndSwitch, error) {
	if !IsPaladinsPath(a.BasePath) {
		return nil, fmt.Errorf("GetPlayerIDInfoForXBOXAndSwitch() %s", notPaladinsErrMsg)
	}
	resp, err := a.makeRequest("getplayeridinfoforxboxandswitch", player)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []models.PlayerIDInfoForXBOXAndSwitch
	err = a.unmarshalResponse(body, &output)
	return output, err
}

// GetPlayerLoadouts returns deck loadouts per Champion
func (a *APIClient) GetPlayerLoadouts(player, langCode string) ([]models.PlayerLoadout, error) {
	if !IsPaladinsPath(a.BasePath) {
		return nil, fmt.Errorf("GetPlayerLoadouts() %s", notPaladinsErrMsg)
	}
	if langCode == "" {
		langCode = English
	}
	path := fmt.Sprintf("%s/%s", player, langCode)
	resp, err := a.makeRequest("getplayerloadouts", path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []models.PlayerLoadout
	err = a.unmarshalResponse(body, &output)
	return output, err
}


// GetChampionCards returns all Champion cards.
func (a *APIClient) GetChampionCards(champID, langCode string) ([]models.ChampionCard, error) {
	if !IsPaladinsPath(a.BasePath) {
		return nil, fmt.Errorf("GetChampionCards() %s", notPaladinsErrMsg)
	}
	path := fmt.Sprintf("%s/%s", champID, langCode)
	resp, err := a.makeRequest("getchampioncards", path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []models.ChampionCard
	err = a.unmarshalResponse(body, &output)
	return output, err
}

// GetBountyItems returns daily Bounty Item history for the past 6 months.
func (a *APIClient) GetBountyItems() ([]models.BountyItem, error) {
	if !IsPaladinsPath(a.BasePath) {
		return nil, fmt.Errorf("GetBountyItems() %s", notPaladinsErrMsg)
	}
	resp, err := a.makeRequest("getbountyitems", "")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []models.BountyItem
	err = a.unmarshalResponse(body, &output)
	return output, err
}