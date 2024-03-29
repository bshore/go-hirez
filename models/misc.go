package models

// ESportsProLeagueDetail stores data related to ESports Matches
type ESportsProLeagueDetail struct {
	AwayTeamClanID    int64  `json:"away_team_clan_id"`
	AwayTeamName      string `json:"away_team_name"`
	AwayTeamTagName   string `json:"away_team_tagname"`
	HomeTeamClanID    int64  `json:"home_team_clan_id"`
	HomeTeamName      string `json:"home_team_name"`
	HomeTeamTagName   string `json:"home_team_tagname"`
	MapInstanceID     string `json:"map_instance_id"`
	MatchDate         string `json:"match_date"`
	MatchNumber       string `json:"match_number"`
	MatchStatus       string `json:"match_status"`
	MatchupID         string `json:"matchup_id"`
	Region            string `json:"region"`
	RetMsg            string `json:"ret_msg"`
	TournamentName    string `json:"tournament_name"`
	WinningTeamClanID int64  `json:"winning_team_clan_id"`
}

// GodLeaderboardEntry stores data related to the God Leaderboard
type GodLeaderboardEntry struct {
	GodID         string `json:"god_id"`
	Losses        string `json:"losses"`
	PlayerID      string `json:"player_id"`
	PlayerName    string `json:"player_name"`
	PlayerRanking string `json:"player_ranking"`
	Rank          string `json:"rank"`
	RetMsg        string `json:"ret_msg"`
	Wins          string `json:"wins"`
}

// LeagueLeaderboardEntry stores data related to the League Leaderboard
type LeagueLeaderboardEntry struct {
	Leaves       int64  `json:"Leaves"`
	Losses       int64  `json:"Losses"`
	Name         string `json:"Name"`
	Points       int64  `json:"Points"`
	PrevRank     int64  `json:"PrevRank"`
	Rank         int64  `json:"Rank"`
	RankStat     int64  `json:"Rank_Stat"`
	RankVariance int64  `json:"Rank_Variance"`
	Season       int64  `json:"Season"`
	Tier         int64  `json:"Tier"`
	Trend        int64  `json:"Trend"`
	Wins         int64  `json:"Wins"`
	PlayerID     string `json:"player_id"`
	RetMsg       string `json:"ret_msg"`
}

// Season stores data related to a game Season
type Season struct {
	Complete          bool   `json:"complete"`
	ID                int64  `json:"id"`
	LeagueDescription string `json:"league_description"`
	Name              string `json:"name"`
	RetMsg            string `json:"ret_msg"`
	Round             int64  `json:"round"`
	Season            int64  `json:"season"`
}

// MOTD stores data related to Match Of The Day modes
type MOTD struct {
	Description   string `json:"description"`
	GameMode      string `json:"gameMode"`
	MaxPlayers    string `json:"maxPlayers"`
	Name          string `json:"name"`
	RetMsg        string `json:"ret_msg"`
	StartDateTime string `json:"startDateTime"`
	Team1GodsCSV  string `json:"team1GodsCSV"`
	Team2GodsCSV  string `json:"team2GodsCSV"`
	Title         string `json:"title"`
}

// TopMatch stores data related to a recent Top Match
type TopMatch struct {
	Ban1              string `json:"Ban1"`
	Ban1ID            int64  `json:"Ban1Id"`
	Ban2              string `json:"Ban2"`
	Ban2ID            int64  `json:"Ban2Id"`
	EntryDatetime     string `json:"Entry_Datetime"`
	LiveSpectators    int64  `json:"LiveSpectators"`
	Match             int64  `json:"Match"`
	MatchTime         int64  `json:"Match_Time"`
	OfflineSpectators int64  `json:"OfflineSpectators"`
	Queue             string `json:"Queue"`
	RecordingFinished string `json:"RecordingFinished"`
	RecordingStarted  string `json:"RecordingStarted"`
	Team1AvgLevel     int64  `json:"Team1_AvgLevel"`
	Team1Gold         int64  `json:"Team1_Gold"`
	Team1Kills        int64  `json:"Team1_Kills"`
	Team1Score        int64  `json:"Team1_Score"`
	Team2AvgLevel     int64  `json:"Team2_AvgLevel"`
	Team2Gold         int64  `json:"Team2_Gold"`
	Team2Kills        int64  `json:"Team2_Kills"`
	Team2Score        int64  `json:"Team2_Score"`
	WinningTeam       int64  `json:"WinningTeam"`
	RetMsg            string `json:"ret_msg"`
}

// VerstionInfo stores data related to the version of a game patch
type VersionInfo struct {
	RetMsg        string `json:"ret_msg"`
	VersionString string `json:"version_string"`
}
