package models

type Champion struct {
	Ability1               string          `json:"Ability1"`
	Ability2               string          `json:"Ability2"`
	Ability3               string          `json:"Ability3"`
	Ability4               string          `json:"Ability4"`
	Ability5               string          `json:"Ability5"`
	AbilityID1             int64           `json:"AbilityId1"`
	AbilityID2             int64           `json:"AbilityId2"`
	AbilityID3             int64           `json:"AbilityId3"`
	AbilityID4             int64           `json:"AbilityId4"`
	AbilityID5             int64           `json:"AbilityId5"`
	AbilityOne             ChampionAbility `json:"Ability_1"`
	AbilityTwo             ChampionAbility `json:"Ability_2"`
	AbilityThree           ChampionAbility `json:"Ability_3"`
	AbilityFour            ChampionAbility `json:"Ability_4"`
	AbilityFive            ChampionAbility `json:"Ability_5"`
	ChampionAbility1URL    string          `json:"ChampionAbility1_URL"`
	ChampionAbility2URL    string          `json:"ChampionAbility2_URL"`
	ChampionAbility3URL    string          `json:"ChampionAbility3_URL"`
	ChampionAbility4URL    string          `json:"ChampionAbility4_URL"`
	ChampionAbility5URL    string          `json:"ChampionAbility5_URL"`
	ChampionAbilityCardURL string          `json:"ChampionCard_URL"`
	ChampionAbilityIconURL string          `json:"ChampionIcon_URL"`
	Cons                   string          `json:"Cons"`
	Health                 int64           `json:"Health"`
	Lore                   string          `json:"Lore"`
	Name                   string          `json:"Name"`
	NameEnglish            string          `json:"Name_English"`
	OnFreeRotation         string          `json:"OnFreeRotation"`
	OnFreeWeeklyRotation   string          `json:"OnFreeWeeklyRotation"`
	Pantheon               string          `json:"Pantheon"`
	Pros                   string          `json:"Pros"`
	Roles                  string          `json:"Roles"`
	Speed                  int64           `json:"Speed"`
	Title                  string          `json:"Title"`
	Type                   string          `json:"Type"`
	AbilityDescription1    string          `json:"abilityDescription1"`
	AbilityDescription2    string          `json:"abilityDescription2"`
	AbilityDescription3    string          `json:"abilityDescription3"`
	AbilityDescription4    string          `json:"abilityDescription4"`
	AbilityDescription5    string          `json:"abilityDescription5"`
	ID                     int64           `json:"id"`
	LatestChamption        string          `json:"latestChampion"`
	RetMsg                 string          `json:"ret_msg"`
}

type ChampionAbility struct {
	Description     string `json:"Description"`
	ID              int64  `json:"Id"`
	Summary         string `json:"Summary"`
	URL             string `json:"URL"`
	DamageType      string `json:"damageType"`
	RechargeSeconds int64  `json:"rechargeSeconds"`
}

type ChampionLeaderboardEntry struct {
	ChampionID    string `json:"champion_id"`
	Losses        string `json:"losses"`
	PlayerID      string `json:"player_id"`
	PlayerName    string `json:"player_name"`
	PlayerRanking string `json:"player_ranking"`
	Rank          string `json:"rank"`
	RetMsg        string `json:"ret_msg"`
	Wins          string `json:"wins"`
}

type ChampionSkin struct {
	ChampionID      int64  `json:"champion_id"`
	ChampionName    string `json:"champion_name"`
	Rarity          string `json:"rarity"`
	RetMsg          string `json:"ret_msg"`
	SkinID1         int64  `json:"skin_id1"`
	SkinID2         int64  `json:"skin_id2"`
	SkinName        string `json:"skin_name"`
	SkinNameEnglish string `json:"skin_name_english"`
}

type ChampionRank struct {
	Assists     int64  `json:"Assists"`
	Deaths      int64  `json:"Deaths"`
	Gold        int64  `json:"Gold"`
	Kills       int64  `json:"Kills"`
	LastPlayed  string `json:"LastPlayed"`
	Losses      int64  `json:"Losses"`
	MinionKills int64  `json:"MinionKills"`
	Minutes     int64  `json:"Minutes"`
	Rank        int64  `json:"Rank"`
	Wins        int64  `json:"Wins"`
	Worshippers int64  `json:"Worshippers"`
	Champion    string `json:"champion"`
	ChampionID  string `json:"champion_id"`
	PlayerID    string `json:"player_id"`
	RetMsg      string `json:"ret_msg"`
}

type PlayerLoadout struct {
	ChampionID   int64         `json:"ChampionId"`
	ChampionName string        `json:"ChampionName"`
	DeckID       int64         `json:"DeckId"`
	DeckName     string        `json:"DeckName"`
	LoadoutItems []LoadoutItem `json:"LoadoutItems"`
	PlayerID     int64         `json:"playerId"`
	PlayerName   string        `json:"playerName"`
	RetMsg       string        `json:"ret_msg"`
}

type LoadoutItem struct {
	ItemID   int64  `json:"ItemId"`
	ItemName string `json:"ItemName"`
	Points   int64  `json:"Points"`
}

type PlayerIDInfoForXBOXAndSwitch struct {
	Name         string `json:"Name"`
	Gamertag     string `json:"gamer_tag"`
	Platform     string `json:"platform"`
	PlayerID     string `json:"player_id"`
	PortalUserID int64  `json:"portal_userid"`
	RetMsg       string `json:"ret_msg"`
}

type ChampionCard struct {
	ActiveFlagActivationSchedule string `json:"active_flag_activation_schedule"`
	ActiveFlagLTI                string `json:"active_flag_lti"`
	CardDescription              string `json:"card_description"`
	CardID1                      int64  `json:"card_id1"`
	CardID2                      int64  `json:"card_id2"`
	CardName                     string `json:"card_name"`
	CardNameEnglish              string `json:"card_name_english"`
	ChampionCardURL              string `json:"championCard_URL"`
	ChampionCardIconURL          string `json:"championIcon_URL"`
	ChampionTalentURL            string `json:"championTalent_URL"`
	ChampionID                   int64  `json:"champion_id"`
	ChampionName                 string `json:"champion_name"`
	Exclusive                    string `json:"exclusive"`
	Rank                         int64  `json:"rank"`
	Rarity                       string `json:"rarity"`
	RechargeSeconds              int64  `json:"recharge_seconds"`
	RetMsg                       string `json:"ret_msg"`
}

type BountyItem struct {
	Active          string `json:"active"`
	BountyItemID1   int64  `json:"bounty_item_id1"`
	BountyItemID2   int64  `json:"bounty_item_id2"`
	BountyItemName  string `json:"bounty_item_name"`
	ChampionID      int64  `json:"champion_id"`
	ChampionName    string `json:"champion_name"`
	FinalPrice      string `json:"final_price"`
	InitialPrice    string `json:"initial_price"`
	RetMsg          string `json:"ret_msg"`
	SaleEndDatetime string `json:"sale_end_datetime"`
	SaleType        string `json:"sale_type"`
}