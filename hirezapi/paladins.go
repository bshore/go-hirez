package hirezapi

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/bshore/go-hirez/models"
)

// GetChampionRanks returns the rank and worshipper values for each Champion a player has played.
func (a *APIClient) GetChampionRanks(player string) error {
	resp, err := a.makeRequest("getchampionranks", player)
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
func (a *APIClient) GetChampions(langCode string) ([]models.Champion, error) {
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

// NOTE: No idea what's going on with this endpoint...
// GetPlayerIDInfoForXBOXAndSwitch returns all PlayerID data associated with the playerName
// func (a *APIClient) GetPlayerIDInfoForXBOXAndSwitch(player string) error {
// 	if !IsNotSmitePath(a.BasePath) || !strings.Contains(a.BasePath, "xbox") {
// 		return fmt.Errorf("GetPlayerIDInfoForXBOXAndSwitch() is only supported for Paladins XBOX requests")
// 	}
// 	resp, err := a.makeRequest("getplayeridinfoforxboxandswitch", player)
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return err
// 	}
// 	err = os.WriteFile("output.json", body, 0644)
// 	return err
// }

// GetPlayerLoadouts returns deck loadouts per Champion
func (a *APIClient) GetPlayerLoadouts(player, langCode string) ([]models.PlayerLoadout, error) {
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
