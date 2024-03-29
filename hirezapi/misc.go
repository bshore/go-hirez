package hirezapi

import (
	"fmt"
	"io"

	"github.com/bshore/go-hirez/models"
	"github.com/bshore/go-hirez/utils"
)

// GetESportsProLeagueDetails returns matchup info for each matchup of the current season. match_status = 1 - scheduled, 2 - in progress, 3 - complete
func (a *APIClient) GetESportsProLeagueDetails() ([]models.ESportsProLeagueDetail, error) {
	resp, err := a.makeRequest("getesportsproleaguedetails", "")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []models.ESportsProLeagueDetail
	err = a.unmarshalResponse(body, &output)
	return output, err
}

// GetGodLeaderboard returns the current season's leaderboard for a god/queue. [SmiteAPI: only queues 440, 450, 451]
func (a *APIClient) GetGodLeaderboard(godID, queueID string) ([]models.GodLeaderboardEntry, error) {
	if !utils.IsSmitePath(a.BasePath) {
		return nil, fmt.Errorf("GetGodLeaderboard() %s", utils.NotSmiteErrMsg)
	}
	if utils.IsNotRanked(queueID) {
		return nil, fmt.Errorf("GetGodLeaderboard() %s", utils.NotRankedErrMsg)
	}
	path := fmt.Sprintf("%s/%s", godID, queueID)
	resp, err := a.makeRequest("getgodleaderboard", path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []models.GodLeaderboardEntry
	err = a.unmarshalResponse(body, &output)
	return output, err
}

// GetLeagueLeaderboard returns the top players for a particular league.
func (a *APIClient) GetLeagueLeaderboard(queueID, tier, round string) ([]models.LeagueLeaderboardEntry, error) {
	if !utils.IsSmitePath(a.BasePath) {
		return nil, fmt.Errorf("GetLeageLeaderboard() %s", utils.NotSmiteErrMsg)
	}
	if utils.IsNotRanked(queueID) {
		return nil, fmt.Errorf("GetLeagueLeaderboard() %s", utils.NotRankedErrMsg)
	}
	path := fmt.Sprintf("%s/%s/%s", queueID, tier, round)
	resp, err := a.makeRequest("getleagueleaderboard", path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []models.LeagueLeaderboardEntry
	err = a.unmarshalResponse(body, &output)
	return output, err
}

// GetLeagueSeasons returns a list of seasons for a match queue.
func (a *APIClient) GetLeagueSeasons(queueID string) ([]models.Season, error) {
	if !utils.IsSmitePath(a.BasePath) {
		return nil, fmt.Errorf("GetLeageSeasons() %s", utils.NotSmiteErrMsg)
	}
	if utils.IsNotRanked(queueID) {
		return nil, fmt.Errorf("GetLeagueSeasons() %s", utils.NotRankedErrMsg)
	}
	resp, err := a.makeRequest("getleagueseasons", queueID)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []models.Season
	err = a.unmarshalResponse(body, &output)
	return output, err
}

// GetMOTD returns information about the 20 most recent Match-of-the-Days.
func (a *APIClient) GetMOTD() ([]models.MOTD, error) {
	resp, err := a.makeRequest("getmotd", "")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []models.MOTD
	err = a.unmarshalResponse(body, &output)
	return output, err
}

// GetTopMatches returns the 50 most watched / recorded matches.
func (a *APIClient) GetTopMatches() ([]models.TopMatch, error) {
	resp, err := a.makeRequest("gettopmatches", "")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []models.TopMatch
	err = a.unmarshalResponse(body, &output)
	return output, err
}

// GetPatchInfo returns information about the currently deployed patch.
func (a *APIClient) GetPatchInfo() (*models.VersionInfo, error) {
	resp, err := a.makeRequest("getpatchinfo", "")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output *models.VersionInfo
	err = a.unmarshalResponse(body, &output)
	return output, err
}
