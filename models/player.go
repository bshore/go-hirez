package models

type Player struct {
	ActivePlayerID             int64            `json:"ActivePlayerId"`
	AvatarURL                  string           `json:"Avatar_URL"`
	CreatedDatetime            string           `json:"Created_Datetime"`
	HoursPlayed                int64            `json:"HoursPlayed"`
	ID                         int64            `json:"Id"`
	LastLoginDatetime          string           `json:"Last_Login_Datetime"`
	Leaves                     int64            `json:"Leaves"`
	Level                      int64            `json:"Level"`
	Losses                     int64            `json:"Losses"`
	MasteryLevel               int              `json:"MasteryLevel"`
	MergedPlayers              []MergedPlayer   `json:"MergedPlayers"`
	MinutesPlayed              int64            `json:"MinutesPlayed"`
	Name                       string           `json:"Name"`
	PersonalStatusMessage      string           `json:"Personal_Status_Message"`
	Platform                   string           `json:"Platform"`
	RankStatConquest           int              `json:"Rank_Stat_Conquest"`
	RankStatConquestController int              `json:"Rank_Stat_Conquest_Controller"`
	RankStatDuel               int              `json:"Rank_Stat_Duel"`
	RankStatDuelController     int              `json:"Rank_Stat_Duel_Controller"`
	RankStatJoust              int              `json:"Rank_Stat_Joust"`
	RankStatJoustController    int              `json:"Rank_Stat_Joust_Controller"`
	RankedConquest             PlayerRankedInfo `json:"RankedConquest"`
	RankConquestController     PlayerRankedInfo `json:"RankedConquestController"`
	RankDuel                   PlayerRankedInfo `json:"RankedDuel"`
	RankDuelController         PlayerRankedInfo `json:"RankedDuelController"`
	RankJoust                  PlayerRankedInfo `json:"RankedJoust"`
	RankJoustController        PlayerRankedInfo `json:"RankedJoustController"`
	Region                     string           `json:"Region"`
	TeamID                     int64            `json:"TeamId"`
	TeamName                   string           `json:"Team_Name"`
	TierConquest               int              `json:"Tier_Conquest"`
	TierDuel                   int              `json:"Tier_Duel"`
	TierJoust                  int              `json:"Tier_Joust"`
	TotalAchievements          int64            `json:"Total_Achievements"`
	TotalWorshippers           int64            `json:"Total_Worshippers"`
	Wins                       int64            `json:"Wins"`
	HZGamerTag                 string           `json:"hz_gamer_tag"`
	HZPlayerName               string           `json:"hz_player_name"`
	RetMsg                     string           `json:"ret_msg"`
}

type MergedPlayer struct {
	MergeDatetime string `json:"merge_datetime"`
	PlayerID      string `json:"playerId"`
	PortalID      string `json:"portalId"`
}

type PlayerRankedInfo struct {
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

type Friend struct {
	AccountID   string `json:"account_id"`
	AvatarURL   string `json:"avatar_url"`
	FriendFlags string `json:"friend_flags"`
	Name        string `json:"name"`
	PlayerID    string `json:"player_id"`
	PortalID    string `json:"portal_id"`
	RetMsg      string `json:"ret_msg"`
	Status      string `json:"status"`
}

type GodRank struct {
	Assists     int64  `json:"Assists"`
	Deaths      int64  `json:"Deaths"`
	Kills       int64  `json:"Kills"`
	Losses      int64  `json:"Losses"`
	MinionKills int64  `json:"MinionKills"`
	Rank        int    `json:"Rank"`
	Wins        int64  `json:"Wins"`
	Worshippers int64  `json:"Worshippers"`
	God         string `json:"god"`
	GodID       string `json:"god_id"`
	PlayerID    string `json:"player_id"`
	RetMsg      string `json:"ret_msg"`
}

type PlayerStatus struct {
	Match                 int64  `json:"Match"`
	MatchQueueID          int64  `json:"match_queue_id"`
	PersonalStatusMessage string `json:"personal_status_message"`
	RetMsg                string `json:"ret_msg"`
	Status                int    `json:"status"`
	StatusString          string `json:"status_string"`
}

type PlayerAchievements struct {
	AssistedKills        int64  `json:"AssistedKills"`
	CampsCleared         int64  `json:"CampsCleared"`
	Deaths               int64  `json:"Deaths"`
	DivineSpree          int64  `json:"DivineSpree"`
	DoubleKills          int64  `json:"DoubleKills"`
	FireGiantKills       int64  `json:"FireGiantKills"`
	FirstBloods          int64  `json:"FirstBloods"`
	GodLikeSpree         int64  `json:"GodLikeSpree"`
	GoldFuryKills        int64  `json:"GoldFuryKills"`
	ID                   int64  `json:"Id"`
	ImmortalSpree        int64  `json:"ImmortalSpree"`
	KillingSpree         int64  `json:"KillingSpree"`
	MinionKills          int64  `json:"MinionKills"`
	Name                 string `json:"Name"`
	PentaKills           int64  `json:"PentaKills"`
	PhoenixKills         int64  `json:"PhoenixKills"`
	PlayerKills          int64  `json:"PlayerKills"`
	QuadraKills          int64  `json:"QuadraKills"`
	RampageSpree         int64  `json:"RampageSpree"`
	ShutdownSpree        int64  `json:"ShutdownSpree"`
	SiegeJuggernautKills int64  `json:"SiegeJuggernautKills"`
	TowerKills           int64  `json:"TowerKills"`
	TripleKills          int64  `json:"TripleKills"`
	UnstoppableSpree     int64  `json:"UnstoppableSpree"`
	WildJuggernautKills  int64  `json:"WildJuggernautKills"`
	RetMsg               string `json:"ret_msg"`
}

type TeamDetail struct {
	Founder   string `json:"Founder"`
	FounderID string `json:"FounderId,omitempty"`
	Losses    int64  `json:"Losses,omitempty"`
	Name      string `json:"Name"`
	Players   int64  `json:"Players"`
	Rating    int64  `json:"Rating,omitempty"`
	Tag       string `json:"Tag"`
	TeamID    int64  `json:"TeamId"`
	Wins      int64  `json:"Wins,omitempty"`
	RetMsg    string `json:"ret_msg"`
}

type TeamPlayer struct {
	AccountLevel      int    `json:"AccountLevel"`
	JoinedDatetime    string `json:"JoinedDatetime"`
	LastLoginDatetime string `json:"LastLoginDatetime"`
	Name              string `json:"Name"`
	RetMsg            string `json:"ret_msg"`
}