package models

type PlayerResponse struct {
	ActivePlayerID             int64                    `json:"ActivePlayerId"`
	AvatarURL                  string                   `json:"Avatar_URL"`
	CreatedDatetime            string                   `json:"Created_Datetime"`
	HoursPlayed                string                   `json:"HoursPlayed"`
	ID                         int64                    `json:"Id"`
	LastLoginDatetime          string                   `json:"Last_Login_Datetime"`
	Leaves 										 int64                    `json:"Leaves"`
	Level 									 	 int64                    `json:"Level"`
	Losses                     int64                    `json:"Losses"`
	MasteryLevel               int                      `json:"MasteryLevel"`
	MergedPlayers              []MergedPlayerResponse   `json:"MergedPlayers"`
	MinutesPlayed              int64                    `json:"MinutesPlayed"`
	Name                       string                   `json:"Name"`
	PersonalStatusMessage      string                   `json:"Personal_Status_Message"`
	Platform                   string                   `json:"Platform"`
	RankStatConquest           int                      `json:"Rank_Stat_Conquest"`
	RankStatConquestController int                      `json:"Rank_Stat_Conquest_Controller"`
	RankStatDuel               int                      `json:"Rank_Stat_Duel"`
	RankStatDuelController     int                      `json:"Rank_Stat_Duel_Controller"`
	RankStatJoust              int                      `json:"Rank_Stat_Joust"`
	RankStatJoustController    int                      `json:"Rank_Stat_Joust_Controller"`
	RankedConquest             PlayerRankedInfoResponse `json:"RankedConquest"`
	RankConquestController     PlayerRankedInfoResponse `json:"RankedConquestController"`
	RankDuel                   PlayerRankedInfoResponse `json:"RankedDuel"`
	RankDuelController         PlayerRankedInfoResponse `json:"RankedDuelController"`
	RankJoust                  PlayerRankedInfoResponse `json:"RankedJoust"`
	RankJoustController        PlayerRankedInfoResponse `json:"RankedJoustController"`
	Region                     string                   `json:"Region"`
	TeamID                     int64                    `json:"TeamId"`
	TeamName                   string                   `json:"Team_Name"`
	TierConquest               int                      `json:"Tier_Conquest"`
	TierDuel                   int                      `json:"Tier_Duel"`
	TierJoust                  int                      `json:"Tier_Joust"`
	TotalAchievements          int64                    `json:"Total_Achievements"`
	TotalWorshippers           int64                    `json:"Total_Worshippers"`
	Wins                       int64                    `json:"Wins"`
	HZGamerTag                 string                   `json:"hz_gamer_tag"`
	HZPlayerName               string                   `json:"hz_player_name"`
	RetMsg                     string                   `json:"ret_msg"`
}

type MergedPlayerResponse struct {
	MergeDatetime string `json:"merge_datetime"`
	PlayerID      string `json:"playerId"`
	PortalID      string `json:"portalId"`
}

type PlayerRankedInfoResponse struct {
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

type FriendsResponse struct {
	AccountID   string `json:"account_id"`
	AvatarURL   string `json:"avatar_url"`
	FriendFlags string `json:"friend_flags"`
	Name        string `json:"name"`
	PlayerID    string `json:"player_id"`
	PortalID    string `json:"portal_id"`
	RetMsg      string `json:"ret_msg"`
	Status      string `json:"status"`
}

type GodRanksResponse struct {
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

type MatchHistoryResponse struct {
	ActiveID1           int64  `json:"ActiveId1"`
	ActiveID2           int64  `json:"ActiveId2"`
	Active1             string `json:"Active_1"`
	Active2             string `json:"Active_2"`
	Active3             string `json:"Active_3"`
	Assists             int64  `json:"Assists"`
	Ban1                string `json:"Ban1"`
	Ban10               string `json:"Ban10"`
	Ban10ID             int64  `json:"Ban10Id"`
	Ban1ID              int64  `json:"Ban1Id"`
	Ban2                string `json:"Ban2"`
	Ban2ID              int64  `json:"Ban2Id"`
	Ban3                string `json:"Ban3"`
	Ban3ID              int64  `json:"Ban3Id"`
	Ban4                string `json:"Ban4"`
	Ban4ID              int64  `json:"Ban4Id"`
	Ban5                string `json:"Ban5"`
	Ban5ID              int64  `json:"Ban5Id"`
	Ban6                string `json:"Ban6"`
	Ban6ID              int64  `json:"Ban6Id"`
	Ban7                string `json:"Ban7"`
	Ban7ID              int64  `json:"Ban7Id"`
	Ban8                string `json:"Ban8"`
	Ban8ID              int64  `json:"Ban8Id"`
	Ban9                string `json:"Ban9"`
	Ban9ID              int64  `json:"Ban9Id"`
	Creeps              int64  `json:"Creeps"`
	Damage              int64  `json:"Damage"`
	DamageBot           int64  `json:"Damage_Bot"`
	DamageDoneInHand    int64  `json:"Damage_Done_In_Hand"`
	DamageMitigated     int64  `json:"Damage_Mitigated"`
	DamageStructure     int64  `json:"Damage_Structure"`
	DamageTaken         int64  `json:"Damage_Taken"`
	DamageTakenMagical  int64  `json:"Damage_Taken_Magical"`
	DamageTakenPhysical int64  `json:"Damage_Taken_Physical"`
	Deaths              int64  `json:"Deaths"`
	DistanceTraveled    int64  `json:"Distance_Traveled"`
	FirstBanSide        string `json:"First_Ban_Side"`
	God                 string `json:"God"`
	GodID               int64  `json:"GodId"`
	Gold                int64  `json:"Gold"`
	Healing             int64  `json:"Healing"`
	HealingBot          int64  `json:"Healing_Bot"`
	HealingPlayerSelf   int64  `json:"Healing_Player_Self"`
	ItemID1             int64  `json:"ItemId1"`
	ItemID2             int64  `json:"ItemId2"`
	ItemID3             int64  `json:"ItemId3"`
	ItemID4             int64  `json:"ItemId4"`
	ItemID5             int64  `json:"ItemId5"`
	ItemID6             int64  `json:"ItemId6"`
	Item1               string `json:"Item_1"`
	Item2               string `json:"Item_2"`
	Item3               string `json:"Item_3"`
	Item4               string `json:"Item_4"`
	Item5               string `json:"Item_5"`
	Item6               string `json:"Item_6"`
	KillingSpree        int64  `json:"Killing_Spree"`
	Kills               int64  `json:"Kills"`
	Level               int    `json:"Level"`
	MapGame             string `json:"Map_Game"`
	Match               int64  `json:"Match"`
	MatchQueueID        int    `json:"Match_Queue_Id"`
	MatchTime           string `json:"Match_Time"`
	Minutes             int64  `json:"Minutes"`
	MultiKillMax        int64  `json:"Multi_kill_Max"`
	ObjectiveAssists    int64  `json:"Objective_Assists"`
	Queue               string `json:"Queue"`
	Region              string `json:"Region"`
	Skin                string `json:"Skin"`
	SkinID              int64  `json:"SkinId"`
	Surrendered         int64  `json:"Surrendered"`
	TaskForce           int64  `json:"TaskForce"`
	Team1Score          int64  `json:"Team1Score"`
	Team2Score          int64  `json:"Team2Score"`
	TimeInMatchSeconds  int64  `json:"Time_In_Match_Seconds"`
	WardsPlaced         int64  `json:"Wards_Placed"`
	WinStatus           string `json:"Win_Status"`
	WinningTaksForce    int64  `json:"Winning_TaskForce"`
	PlayerID            int64  `json:"playerId"`
	PlayerName          string `json:"playerName"`
	RetMsg              string `json:"ret_msg"`
}


type PlayerStatusResponse struct {
	Match                 int64  `json:"Match"`
	MatchQueueID          int64  `json:"match_queue_id"`
	PersonalStatusMessage string `json:"personal_status_message"`
	RetMsg                string `json:"ret_msg"`
	Status                int    `json:"status"`
	StatusString          string `json:"status_string"`
}

type PlayerAchievementsOutput struct {
	AssistedKills int64 `json:"AssistedKills"`
	CampsCleared int64 `json:"CampsCleared"`
	Deaths int64 `json:"Deaths"`
	DivineSpree int64 `json:"DivineSpree"`
	DoubleKills int64 `json:"DoubleKills"`
	FireGiantKills int64 `json:"FireGiantKills"`
	FirstBloods int64 `json:"FirstBloods"`
	GodLikeSpree int64 `json:"GodLikeSpree"`
	GoldFuryKills int64 `json:"GoldFuryKills"`
	ID int64 `json:"Id"`
	ImmortalSpree int64 `json:"ImmortalSpree"`
	KillingSpree int64 `json:"KillingSpree"`
	MinionKills int64 `json:"MinionKills"`
	Name string `json:"Name"`
	PentaKills int64 `json:"PentaKills"`
	PhoenixKills int64 `json:"PhoenixKills"`
	PlayerKills int64 `json:"PlayerKills"`
	QuadraKills int64 `json:"QuadraKills"`
	RampageSpree int64 `json:"RampageSpree"`
	ShutdownSpree int64 `json:"ShutdownSpree"`
	SiegeJuggernautKills int64 `json:"SiegeJuggernautKills"`
	TowerKills int64 `json:"TowerKills"`
	TripleKills int64 `json:"TripleKills"`
	UnstoppableSpree int64 `json:"UnstoppableSpree"`
	WildJuggernautKills int64 `json:"WildJuggernautKills"`
	RetMsg string `json:"ret_msg"`
}