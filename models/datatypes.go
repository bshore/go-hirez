package models

// URL determinmes the baseURL to use when calling Hi-Rez API endpoints
type URL int

const (
	URLSmitePC URL = iota
	URLPaladinsPC
)

func (u URL) String() string {
	return [...]string{
		"http://api.smitegame.com/smiteapi.svc",
		"http://api.paladins.com/paladinsapi.svc",
	}[u]
}

// IsURL exists to satisfy the MustBeURL interface
func (u URL) IsURL() bool {
	return true
}

// MustBeURL exists to force the constructors to use our URL type
type MustBeURL interface {
	String() string
	IsURL() bool
}

// ResponseType determines the type of the Response body: JSON or XML
type ResponseType int

const (
	ResponseTypeJSON ResponseType = iota
	ResponseTypeXML
)

func (r ResponseType) String() string {
	return [...]string{"json", "xml"}[r]
}

// IsResponseType exists to satisfy the MustBeResponseType interface
func (r ResponseType) IsResponseType() bool {
	return true
}

// MustBeResponseType exists to force the constructors to use our ResponseType type
type MustBeResponseType interface {
	String() string
	IsResponseType() bool
}

const (
	// yyyyMMdd
	DateFormat string = "20060102"

	// yyyyMMddHHmmss
	TimeFormat string = "20060102150405"

	// PortalIDs (Gaming Platforms)
	HiRez, Steam, PS4, XBOX, Switch, Discord, Epic string = "1", "5", "9", "10", "22", "25", "28"

	// Language Codes
	English string = "1"
	German string = "2"
	French string = "3"
	Chinese string = "5"
	Spanish string = "7"
	SpanishLAM string = "9"
	Portuguese string = "10"
	Russian string = "11"
	Polish string = "12"
	Turkish string = "13"

	// League Tiers
	Bronze5, Bronze4, Bronze3, Bronze2, Bronze1 string = "1", "2", "3", "4", "5"
	Silver5, Silver4, Silver3, Silver2, Silver1 string = "6", "7", "8", "9", "10"
	Gold5, Gold4, Gold3, Gold2, Gold1 string = "11", "12", "13", "14", "15"
	Platinum5, Platinum4, Platinum3, Platinum2, Platinum1 string = "16", "17", "18", "19", "20"
	Diamond5, Diamond4, Diamond3, Diamond2, Diamond1 string = "21", "22", "23", "24", "25"
	Masters, Grandmaster string = "26", "27"
	

	// Smite Match Queue IDs
	Adventures string = "495"
	Arena string = "435"
	ArenaAIEasy string = "457"
	ArenaAIMedium string = "468"
	ArenaChallenge string = "438"
	ArenaPracticeEasy string = "443"
	ArenaPracticeMedium string = "472"
	ArenaTraining string = "483"
	ArenaTutorial string = "462"
	Assault string = "445"
	AssaultAIEasy string = "481"
	AssaultAIMedium string = "454"
	AssaultChallenge string = "446"
	AssaultPracticeEasy string = "479"
	AssaultPracticeMedium string = "480"
	BasicTutorial string = "436"
	Clash string = "466"
	ClashAIEasy string = "478"
	ClashAIMedium string = "469"
	ClashChallenge string = "467"
	ClashPracticeEasy string = "470"
	ClashPracticeMedium string = "477"
	ClashTutorial string = "471"
	ConqeustChallenge string = "429"
	Conquest string = "426"
	ConquestAIEasy string = "476"
	ConquestAIMedium string = "461"
	ConquestPracticeEasy string = "458"
	ConquestPracticeMedium string = "475"
	ConquestRanked string = "451"
	ConquestTutorial string = "463"
	Joust string = "448"
	Joust1v1Ranked  string = "440"
	Joust3v3Ranked string = "450"
	Joust3v3Training string = "482"
	JoustAIEasy string = "474"
	JoustAIMedium string = "456"
	JoustChallenge string = "441"
	JoustPracticeEasy string = "464"
	JunglePracicePreselect string = "496"
	JunglePracitceMedium string = "473"
	JunglePractice string = "444"
	MatchOfTheDay string = "434"
	Siege string = "459"
	SiegeChallenge string = "460"

	// Paladins Match Queue IDs
	ClassicSiege string = "465"
	CustomOnslaughtForemansRise string = "462"
	CustomOnslaughtMagistratesArchives string = "464"
	CustomOnslaughtPrimalCourt string = "455"
	CustomOnslaughtSnowfallJunction string = "454"
	CustomSiegeBrightmarsh string = "458"
	CustomSiegeFishMarket string = "431"
	CustomSiegeFrogIsle string = "433"
	CustomSiegeFrozenGuard string = "432"
	CustomSiegeIceMines string = "439"
	CustomSiegeJaguarFalls string = "438"
	CustomSiegeSerpentBeach string = "440"
	CustomSiegeSplitstoneQuarry string = "459"
	CustomSiegeStoneKeep string = "423"
	CustomSiegeTimberMill string = "430"
	CustomTDMForemansRise string = "471"
	CustomTDMMagistrateArchives string = "472"
	CustomTDMTradeDistrict string = "468"
	LiveBattleGroundsDuo string = "475"
	LiveBattleGroundsQuad string = "476"
	LiveBattlegroundsSolo string = "474"
	LiveCasual string = "424"
	LiveCompetitive string = "428"
	LiveOnslaught string = "452"
	LiveOnslaughtPractice string = "453"
	LiveSiegePractice string = "425"
	LiveTeamDeathmatch string = "469"
	LiveTeamDeathmatchPractice string = "470"
	LiveTestMaps string = "445"
)