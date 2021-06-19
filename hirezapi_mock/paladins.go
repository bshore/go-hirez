package mock

import (
	"fmt"
	"strings"

	"github.com/bshore/go-hirez/models"
	"github.com/bshore/go-hirez/utils"
)

// GetPlayerBatch returns league and other high level data for a particular list of players. [20 max]
func (a *APIClient) GetPlayerBatch(playerIDs []string) ([]models.Player, error) {
	if len(playerIDs) > 20 {
		return nil, fmt.Errorf("per API docs, the list of playerIDs should contain no more than 20")
	}
	if !utils.IsPaladinsPath(a.BasePath) {
		return nil, fmt.Errorf("GetPlayerBatch() %s", utils.NotPaladinsErrMsg)
	}
	resp, err := a.makeRequest("getplayerbatch", strings.Join(playerIDs, ","), []models.Player{})
	if err != nil {
		return nil, err
	}
	var output []models.Player
	err = a.unmarshalResponse(resp, &output)
	return output, err
}

// GetChampionRanks returns the rank and worshipper values for each Champion a player has played.
func (a *APIClient) GetChampionRanks(player string) ([]models.ChampionRank, error) {
	if !utils.IsPaladinsPath(a.BasePath) {
		return nil, fmt.Errorf("GetChampionRanks() %s", utils.NotPaladinsErrMsg)
	}
	resp, err := a.makeRequest("getchampionranks", player, []models.ChampionRank{})
	if err != nil {
		return nil, err
	}
	var output []models.ChampionRank
	err = a.unmarshalResponse(resp, &output)
	return output, err
}

// GetChampions returns all Champions and their various attributes.
func (a *APIClient) GetChampions(langCode string) ([]models.Champion, error) {
	if !utils.IsPaladinsPath(a.BasePath) {
		return nil, fmt.Errorf("GetChampions() %s", utils.NotPaladinsErrMsg)
	}
	if langCode == "" {
		langCode = models.English
	}
	resp, err := a.makeRequest("getchampions", langCode, []models.Champion{})
	if err != nil {
		return nil, err
	}
	var output []models.Champion
	err = a.unmarshalResponse(resp, &output)
	return output, err
}

// GetChampionLeaderboard returns the current season's leaderboard for a champion/queue. [ Only queue 428]
func (a *APIClient) GetChampionLeaderboard(champID string) ([]models.ChampionLeaderboardEntry, error) {
	if !utils.IsPaladinsPath(a.BasePath) {
		return nil, fmt.Errorf("GetChampionLeaderboard() %s", utils.NotPaladinsErrMsg)
	}
	path := fmt.Sprintf("%s/428", champID)
	resp, err := a.makeRequest("getchampionleaderboard", path, []models.ChampionLeaderboardEntry{})
	if err != nil {
		return nil, err
	}
	var output []models.ChampionLeaderboardEntry
	err = a.unmarshalResponse(resp, &output)
	return output, err
}

// GetChampionSkins returns a list of skins for a particular champion.
func (a *APIClient) GetChampionSkins(champID, langCode string) ([]models.ChampionSkin, error) {
	if !utils.IsPaladinsPath(a.BasePath) {
		return nil, fmt.Errorf("GetChampionSkins() %s", utils.NotPaladinsErrMsg)
	}
	if langCode == "" {
		langCode = models.English
	}
	path := fmt.Sprintf("%s/%s", champID, langCode)
	resp, err := a.makeRequest("getchampionskins", path, []models.ChampionSkin{})
	if err != nil {
		return nil, err
	}
	var output []models.ChampionSkin
	err = a.unmarshalResponse(resp, &output)
	return output, err
}

// GetPlayerIDInfoForXBOXAndSwitch returns all PlayerID data associated with the playerName
func (a *APIClient) GetPlayerIDInfoForXBOXAndSwitch(player string) ([]models.PlayerIDInfoForXBOXAndSwitch, error) {
	if !utils.IsPaladinsPath(a.BasePath) {
		return nil, fmt.Errorf("GetPlayerIDInfoForXBOXAndSwitch() %s", utils.NotPaladinsErrMsg)
	}
	resp, err := a.makeRequest("getplayeridinfoforxboxandswitch", player, []models.PlayerIDInfoForXBOXAndSwitch{})
	if err != nil {
		return nil, err
	}
	var output []models.PlayerIDInfoForXBOXAndSwitch
	err = a.unmarshalResponse(resp, &output)
	return output, err
}

// GetPlayerLoadouts returns deck loadouts per Champion
func (a *APIClient) GetPlayerLoadouts(player, langCode string) ([]models.PlayerLoadout, error) {
	if !utils.IsPaladinsPath(a.BasePath) {
		return nil, fmt.Errorf("GetPlayerLoadouts() %s", utils.NotPaladinsErrMsg)
	}
	if langCode == "" {
		langCode = models.English
	}
	path := fmt.Sprintf("%s/%s", player, langCode)
	resp, err := a.makeRequest("getplayerloadouts", path, []models.PlayerLoadout{})
	if err != nil {
		return nil, err
	}
	var output []models.PlayerLoadout
	err = a.unmarshalResponse(resp, &output)
	return output, err
}


// GetChampionCards returns all Champion cards.
func (a *APIClient) GetChampionCards(champID, langCode string) ([]models.ChampionCard, error) {
	if !utils.IsPaladinsPath(a.BasePath) {
		return nil, fmt.Errorf("GetChampionCards() %s", utils.NotPaladinsErrMsg)
	}
	path := fmt.Sprintf("%s/%s", champID, langCode)
	resp, err := a.makeRequest("getchampioncards", path, []models.ChampionCard{})
	if err != nil {
		return nil, err
	}
	var output []models.ChampionCard
	err = a.unmarshalResponse(resp, &output)
	return output, err
}

// GetBountyItems returns daily Bounty Item history for the past 6 months.
func (a *APIClient) GetBountyItems() ([]models.BountyItem, error) {
	if !utils.IsPaladinsPath(a.BasePath) {
		return nil, fmt.Errorf("GetBountyItems() %s", utils.NotPaladinsErrMsg)
	}
	resp, err := a.makeRequest("getbountyitems", "", []models.BountyItem{})
	if err != nil {
		return nil, err
	}
	var output []models.BountyItem
	err = a.unmarshalResponse(resp, &output)
	return output, err
}