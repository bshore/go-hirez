package hirezapi

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/bshore/go-hirez/models"
)

// GetMatchDetails returns the statistics for a particular completed match.
func (a *APIClient) GetMatchDetails(matchID string) ([]models.MatchPlayer, error) {
	resp, err := a.makeRequest("getmatchdetails", matchID)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []models.MatchPlayer
	err = a.unmarshalResponse(body, &output)
	return output, err
}

// GetMatchDetailsBatch returns the statistics for a particular set of completed matches. (limit batch query to 5-10 matchIDs)
func (a *APIClient) GetMatchDetailsBatch(matchIDs []string) ([]models.MatchPlayer, error) {
	if len(matchIDs) > 10 {
		return nil, fmt.Errorf("per API docs, the list of matchIDs should contain no more than 10")
	}
	resp, err := a.makeRequest("getmatchdetailsbatch", strings.Join(matchIDs, ","))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []models.MatchPlayer
	err = a.unmarshalResponse(body, &output)
	return output, err
}

// GetOrganizedMatchDetailsBatch is the same as GetMatchDetailsBatch(), except it groups the players by match. (limit batch query to 5-10 matchIDs)
func (a *APIClient) GetOrganizedMatchDetailsBatch(matchIDs []string) ([][]models.MatchPlayer, error) {
	players, err := a.GetMatchDetailsBatch(matchIDs)
	if err != nil {
		return nil, err
	}
	var output [][]models.MatchPlayer
	for _, matchID := range matchIDs {
		id, err := strconv.ParseInt(matchID, 10, 64)
		if err != nil {
			return nil, err
		}
		var playersInMatch []models.MatchPlayer
		for _, player := range players {
			if player.Match == id {
				playersInMatch = append(playersInMatch, player)
			}
		}
		output = append(output, playersInMatch)
	}
	return output, err
}

// GetMatchPlayerDetails returns player information for a live match.
func (a *APIClient) GetMatchPlayerDetails(matchID string) ([]models.LiveMatchPlayer, error) {
	resp, err := a.makeRequest("getmatchplayerdetails", matchID)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []models.LiveMatchPlayer
	err = a.unmarshalResponse(body, &output)
	return output, err
}

/*GetMatchIDsByQueue lists all MatchIDs for a particular match queue.
	- queueID can be referened by constants defined in this package (eg, hirezapi.ConquestRanked).
	- date must be formatted/formattable by hirezapi.DateFormat (yyyyMMdd).
	- hour may be "0" - "23" and optionally may contain a ten minute window separated by a comma (eg, "6,30").
	- hour may also be "-1" to fetch the whole day, but may stall/fail due to the amount of data.
*/
func (a *APIClient) GetMatchIDsByQueue(queueID, date, hour string) ([]models.Match, error) {
	path := fmt.Sprintf("%s/%s/%s", queueID, date, hour)
	resp, err := a.makeRequest("getmatchidsbyqueue", path)
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

// GetQueueStats returns match summary stats for a player + queue, grouped by Gods played.
func (a *APIClient) GetQueueStats(player, queueID string) ([]models.PlayerGodQueueStat, error) {
	path := fmt.Sprintf("%s/%s", player, queueID)
	resp, err := a.makeRequest("getqueuestats", path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var output []models.PlayerGodQueueStat
	err = a.unmarshalResponse(body, &output)
	return output, err
}
