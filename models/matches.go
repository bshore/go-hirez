package models

// Match stores data related to an in game Match
type Match struct {
	ActiveFlag string `json:"Active_Flag"`
	Match      int64  `json:"Match"`
	RetMsg     string `json:"ret_msg"`
}

// MatchPlayer stores data related to a Player present in a Match
type MatchPlayer struct {
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
	PlayerID            string `json:"playerId"`
	PlayerName          string `json:"playerName"`
	RetMsg              string `json:"ret_msg"`
}

// LiveMatchPlayer stores data related to a Player in a currently Live Match
type LiveMatchPlayer struct {
	AccountGodsPlayed int64  `json:"Account_Gods_Played"`
	AccountLevel      int64  `json:"Account_Level"`
	GodID             int64  `json:"GodId"`
	GodLevel          int64  `json:"GodLevel"`
	GodName           string `json:"GodName"`
	MasteryLevel      int64  `json:"Mastery_Level"`
	Match             int64  `json:"Match"`
	Queue             string `json:"Queue"`
	RankStat          int64  `json:"Rank_Stat"`
	SkinID            int64  `json:"SkinId"`
	Tier              int64  `json:"Tier"`
	MapGame           string `json:"mapGame"`
	PlayerCreated     string `json:"playerCreated"`
	PlayerID          string `json:"playerId"`
	PlayerName        string `json:"playerName"`
	PlayerRegion      string `json:"playerRegion"`
	RetMsg            string `json:"ret_msg"`
	TaskForce         int64  `json:"taskForce"`
	TierLosses        int64  `json:"tierLosses"`
	TierWins          int64  `json:"tierWins"`
}

// PlayerGodQueueStat stores data related to a Player's stats for a given Queue type
type PlayerGodQueueStat struct {
	Assists    int64  `json:"Assists"`
	Deaths     int64  `json:"Deaths"`
	God        string `json:"God"`
	GodID      int64  `json:"GodId"`
	Gold       int64  `json:"Gold"`
	Kills      int64  `json:"Kills"`
	LastPlayed string `json:"LastPlayed"`
	Losses     int64  `json:"Losses"`
	Matches    int64  `json:"Matches"`
	Minutes    int64  `json:"Minutes"`
	Queue      string `json:"Queue"`
	Wins       int64  `json:"Wins"`
	PlayerID   string `json:"player_id"`
	RetMsg     string `json:"ret_msg"`
}
