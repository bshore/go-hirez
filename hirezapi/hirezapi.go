package hirezapi

import "github.com/bshore/go-hirez/models"

// HiRezAPI is a collection of endpoint methods for interacting with the Hi-Rez API
// Definitions here are taken from the publicly available "Smite API Developer Guide" PDF
// For the purpose of this project, only Smite-related API endpoints are defined
type HiRezAPI interface {
	// ===== Connectivity =====
	// Ping is a quick way of validating access to the Hi-Rez API.
	Ping() error
	// CreateSession is a required step to Authenticate the developerId/signature for further API use.
	CreateSession() error
	// TestSession is a means of validating that a session is established.
	TestSession() (string, error)
	// GetHirezServerStatus returns UP/DOWN status for the primary game/platform environments. Data is cached once a minute.
	GetHiRezServerStatus() ([]models.HiRezServerStatus, error)
	// GetDataUsed returns API Developer daily usage limits and the current status against those limits.
	GetDataUsed() ([]models.DataUsed, error)
	// ChangeBasePath modifies the base path if you want to query a different platform
	ChangeBasePath(url MustBeURL)
	// ChangeResponseType modifies the response type if you want to switch between JSON and XML
	ChangeResponseType(respType MustBeResponseType)

	// ===== Player Related =====
	// GetPlayer returns league and other high level data for a particular player.
	GetPlayer(player string) ([]models.Player, error)
	// GetFriends returns Smite Usernames of each of the player's friends. [PC Only]
	GetFriends(player string) ([]models.Friend, error)
	// GetGodRanks returns the rank and worshipper values for each God a player has played.
	GetGodRanks(player string) ([]models.GodRank, error)
	// GetMatchHistory returns a list of recent matches and high level match statistics for a particular player
	GetMatchHistory(player string) ([]models.Match, error)
	// GetPlayerStatus returns a player status. 0 - offline, 1 - lobby, 2 - god select, 3 - in game, 4 - online, 5 - unknown
	GetPlayerStatus(player string) ([]models.PlayerStatus, error)
	// GetPlayerAchievements returns select achievement totals.
	GetPlayerAchievements(playerID string) (*models.PlayerAchievements, error)
	// GetTeamDetails returns the number of players and other high level details for a particular Clan.
	GetTeamDetails(clanID string) ([]models.TeamDetail, error)
	// GetTeamPlayers returns a list of players for a particular Clan.
	GetTeamPlayers(clanID string) ([]models.TeamPlayer, error)
	// SearchTeams returns high level info for Clan names containing search term.
	SearchTeams(searchTeam string) ([]models.TeamDetail, error)

	// ===== Game Entity Related =====
	// GetGods returns all Gods and their various attributes.
	GetGods(langCode string) ([]models.God, error)
	// GetGodSkins returns all available skins for a particular God.
	GetGodSkins(godID int64, langCode string) ([]models.GodSkin, error)
	// GetGodRecommendedItems returns the recommended items for a particular God.
	GetGodRecommendedItems(godID int64, langCode string) ([]models.GodRecommendedItem, error)
	// GetItems returns all items and their various attributes.
	GetItems(langCode string) ([]models.Item, error)

	// ===== Match Related =====
	// GetMatchDetails returns the statistics for a particular completed match.
	GetMatchDetails(matchID string) ([]models.MatchPlayer, error)
	/*GetMatchDetailsBatch returns the statistics for a particular set of completed matches. (limit batch query to 5-10 matchIDs)

	- The actual Hi-Rez API does not automatically group the response by matchID.
  - If you want the results to be grouped by matchID, use GetOrganizedMatchDetailsBatch() instead.
	*/
	GetMatchDetailsBatch(matchIDs []string) ([]models.MatchPlayer, error)
	// GetOrganizedMatchDetailsBatch is the same as GetMatchDetailsBatch(), except it groups the players by match.
	GetOrganizedMatchDetailsBatch(matchIDs []string) ([][]models.MatchPlayer, error)
	// GetMatchPlayerDetails returns player information for a live match.
	GetMatchPlayerDetails(matchID string) ([]models.LiveMatchPlayer, error)
	/*GetMatchIDsByQueue lists all MatchIDs for a particular match queue.
	- queueID can be referened by constants defined in this package (eg, hirezapi.ConquestRanked).
	- date must be formatted/formattable by hirezapi.DateFormat (yyyyMMdd).
	- hour may be "0" - "23" and optionally may contain a ten minute window separated by a comma (eg, "6,30").
	- hour may also be "-1" to fetch the whole day, but may stall/fail due to the amount of data.
	*/
	GetMatchIDsByQueue(queueID, date, hour string) ([]models.Match, error)
	// GetQueueStats returns match summary stats for a player + queue, grouped by Gods played.
	GetQueueStats(player, queueID string) ([]models.PlayerGodQueueStat, error)

	// ===== Miscellaneous =====
	// GetESportsProLeagueDetails returns matchup info for each matchup of the current season. match_status = 1 - scheduled, 2 - in progress, 3 - complete
	GetESportsProLeagueDetails() ([]models.ESportsProLeagueDetail, error)
	// GetGodLeaderboard returns the current season's leaderboard for a god/queue. [SmiteAPI: only queues 440, 450, 451 apply]
	GetGodLeaderboard(godID, queueID string) ([]models.GodLeaderboardEntry, error)
	// GetLeagueLeaderboard returns the top players for a particular league.
	GetLeagueLeaderboard(queueID, tier, season string) error
	// GetLeagueSeasons returns a list of seasons for a match queue.
	GetLeagueSeasons(queueID string) ([]models.Season, error)
	// GetMOTD returns information about the 20 most recent Match-of-the-Days.
	GetMOTD() ([]models.MOTD, error)
	// GetTopMatches returns the 50 most watched / recorded matches.
	GetTopMatches() ([]models.TopMatch, error)
	// GetPatchInfo returns information about the currently deployed patch.
	GetPatchInfo() (*models.VersionInfo, error)

	// ===== Paladins =====
	// GetChampionRanks returns the rank and worshipper values for each Champion a player has played.
	GetChampionRanks(player string) error
	// GetChampions returns all Champions and their various attributes.
	GetChampions(langCode string) ([]models.Champion, error)
	// GetChampionLeaderboard returns the current season's leaderboard for a champion/queue. [ Only queue 428]
	GetChampionLeaderboard(champID string) ([]models.ChampionLeaderboardEntry, error)
	// GetChampionSkins
	GetChampionSkins(champID, langCode string) ([]models.ChampionSkin, error)

	// GetPlayerIDInfoForXBOXAndSwitch returns all PlayerID data associated with the playerName
	// GetPlayerIDInfoForXBOXAndSwitch(player string) error

	// GetPlayerLoadouts returns deck loadouts per Champion
	GetPlayerLoadouts(player, langCode string) ([]models.PlayerLoadout, error)
}