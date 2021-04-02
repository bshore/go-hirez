package hirezapi

import "github.com/bshore/go-hirez/models"

// HiRezAPI is a collection of endpoint methods for interacting with the Hi-Rez API
// Definitions here are taken from the publicly available "Smite API Developer Guide" PDF
// For the purpose of this project, only Smite-related API endpoints are defined
type HiRezAPI interface {
	// ===== Connectivity =====
	// Ping is a quick way of validating access to the Hi-Rez API.
	Ping() error
	// GenerateSignature creates the required signature for sending data requests to the Hi-Rez API
	GenerateSignature(methodName string) (string, string)
	// CreateSession is a required step to Authenticate the developerId/signature for further API use.
	CreateSession() error
	// TestSession is a means of validating that a session is established.
	TestSession() (string, error)
	// GetHirezServerStatus returns UP/DOWN status for the primary game/platform environments. Data is cached once a minute.
	GetHiRezServerStatus() ([]models.HiRezServerStatusResponse, error)
	// GetDataUsed returns API Developer daily usage limits and the current status against those limits.
	GetDataUsed() ([]models.DataUsedResponse, error)

	// ===== Player Related =====
	// GetPlayer returns league and other high level data for a particular player.
	GetPlayer(player string) ([]models.PlayerResponse, error)
	// GetFriends returns Smite Usernames of each of the player's friends. [PC Only]
	GetFriends(player string) ([]models.FriendsResponse, error)
	// GetGodRanks returns the rank and worshipper values for each God a player has played.
	GetGodRanks(player string) ([]models.GodRanksResponse, error)
	// GetMatchHistory returns a list of recent matches and high level match statistics for a particular player
	GetMatchHistory(player string) ([]models.MatchHistoryResponse, error)
	// GetPlayerStatus returns a player status. 0 - offline, 1 - lobby, 2 - god select, 3 - in game, 4 - online, 5 - unknown
	GetPlayerStatus(player string) ([]models.PlayerStatusResponse, error)
	// GetPlayerAchievements returns select achievement totals.
	GetPlayerAchievements(playerID string) error
	// GetTeamDetails returns the number of players and other high level details for a particular Clan.
	GetTeamDetails(clanID string) error
	// GetTeamPlayers returns a list of players for a particular Clan.
	GetTeamPlayers(clanID string) error
	// SearchTeams returns high level info for Clan names containing search term.
	SearchTeams(searchTeam string) error

	// ===== Game Entity Related =====
	// GetGods returns all Gods and their various attributes.
	GetGods() error
	// GetGodSkins returns all available skins for a particular God.
	GetGodSkins() error
	// GetGodRecommendedItems returns the recommended items for a particular God.
	GetGodRecommendedItems() error
	// GetItems returns all items and their various attributes.
	GetItems() error

	// ===== Match Related =====
	// GetMatchDetails returns the statistics for a particular completed match.
	GetMatchDetails() error
	// GetMatchDetailsBatch returns the statistics for a particular set of completed matches. [limit batch query to 5-10 matchIDs]
	GetMatchDetailsBatch() error
	// GetMatchPlayerDetails returns player information for a live match.
	GetMatchPlayerDetails() error
	// GetMatchIDsByQueue list all MatchIDs for a particular match queue
	GetMatchIDsByQueue() error
	// GetQueueStats returns match summary stats for a player + queue, grouped by Gods played.
	GetQueueStats() error

	// ===== Miscellaneous =====
	// GetESportsProLeagueDetails returns matchup info for each matchup of the current season. match_status = 1 - scheduled, 2 - in progress, 3 - complete
	GetESportsProLeagueDetails() error
	// GetGodLeaderboard returns the current season's leaderboard for a god/queue. [SmiteAPI: only queues 440, 450, 451 apply]
	GetGodLeaderboard() error
	// GetLeagueLeaderboard returns the top players for a particular league.
	GetLeagueLeaderboard() error
	// GetLeagueSeasons returns a list of seasons for a match queue.
	GetLeagueSeasons() error
	// GetMOTD returns information about the 20 most recent Match-of-the-Days.
	GetMOTD() error
	// GetTopMatches returns the 50 most watched / recorded matches.
	GetTopMatches() error
	// GetPatchInfo returns information about the currently deployed patch.
	GetPatchInfo() error

	// ===== Paladins =====
	// GetChampionRanks returns the rank and worshipper values for each Champion a player has played.
	GetChampionRanks() error
	// GetChampions returns all Champions and their various attributes.
	GetChampions() error
	// GetChampionLeaderboard returns the current season's leaderboard for a champion/queue. [ Only queue 428]
	GetChampionLeaderboard() error
	// GetChampionSkins
	GetChampionSkins() error
	// GetPlayerIDInfoForXBOXAndSwitch returns all PlayerID data associated with the playerName
	GetPlayerIDInfoForXBOXAndSwitch() error
	// GetPlayerLoadouts returns deck loadouts per Champion
	GetPlayerLoadouts() error
}