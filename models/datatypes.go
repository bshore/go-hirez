package models

// Excessive '// nolint' directives are an attempt to keep the Go Report Card happy

const (
	// URLSmitePC is the base URL for making Smite requests
	URLSmitePC string = "https://api.smitegame.com/smiteapi.svc"

	// URLPaladinsPC is the base URL for making Paladins requests
	URLPaladinsPC string = "https://api.paladins.com/paladinsapi.svc"

	// ResponseTypeJSON is the json response type
	ResponseTypeJSON string = "json"

	// ResponseTypeXML is the xml response type
	ResponseTypeXML string = "xml"

	// DateFormat yyyyMMdd
	DateFormat string = "20060102"

	// TimeFormat yyyyMMddHHmmss
	TimeFormat string = "20060102150405"

	// PortalIDs (Gaming Platforms)

	HiRez, Steam, PS4, XBOX, Switch, Discord, Epic string = "1", "5", "9", "10", "22", "25", "28" // nolint

	// Language Codes

	English    string = "1"  // nolint
	German     string = "2"  // nolint
	French     string = "3"  // nolint
	Chinese    string = "5"  // nolint
	Spanish    string = "7"  // nolint
	SpanishLAM string = "9"  // nolint
	Portuguese string = "10" // nolint
	Russian    string = "11" // nolint
	Polish     string = "12" // nolint
	Turkish    string = "13" // nolint

	// League Tiers

	Bronze5, Bronze4, Bronze3, Bronze2, Bronze1           string = "1", "2", "3", "4", "5"      // nolint
	Silver5, Silver4, Silver3, Silver2, Silver1           string = "6", "7", "8", "9", "10"     // nolint
	Gold5, Gold4, Gold3, Gold2, Gold1                     string = "11", "12", "13", "14", "15" // nolint
	Platinum5, Platinum4, Platinum3, Platinum2, Platinum1 string = "16", "17", "18", "19", "20" // nolint
	Diamond5, Diamond4, Diamond3, Diamond2, Diamond1      string = "21", "22", "23", "24", "25" // nolint
	Masters, Grandmaster                                  string = "26", "27"                   // nolint

	// Smite Match Queue IDs

	Adventures             string = "495" // nolint
	Arena                  string = "435" // nolint
	ArenaAIEasy            string = "457" // nolint
	ArenaAIMedium          string = "468" // nolint
	ArenaChallenge         string = "438" // nolint
	ArenaPracticeEasy      string = "443" // nolint
	ArenaPracticeMedium    string = "472" // nolint
	ArenaTraining          string = "483" // nolint
	ArenaTutorial          string = "462" // nolint
	Assault                string = "445" // nolint
	AssaultAIEasy          string = "481" // nolint
	AssaultAIMedium        string = "454" // nolint
	AssaultChallenge       string = "446" // nolint
	AssaultPracticeEasy    string = "479" // nolint
	AssaultPracticeMedium  string = "480" // nolint
	BasicTutorial          string = "436" // nolint
	Clash                  string = "466" // nolint
	ClashAIEasy            string = "478" // nolint
	ClashAIMedium          string = "469" // nolint
	ClashChallenge         string = "467" // nolint
	ClashPracticeEasy      string = "470" // nolint
	ClashPracticeMedium    string = "477" // nolint
	ClashTutorial          string = "471" // nolint
	ConqeustChallenge      string = "429" // nolint
	Conquest               string = "426" // nolint
	ConquestAIEasy         string = "476" // nolint
	ConquestAIMedium       string = "461" // nolint
	ConquestPracticeEasy   string = "458" // nolint
	ConquestPracticeMedium string = "475" // nolint
	ConquestRanked         string = "451" // nolint
	ConquestTutorial       string = "463" // nolint
	Joust                  string = "448" // nolint
	Joust1v1Ranked         string = "440" // nolint
	Joust3v3Ranked         string = "450" // nolint
	Joust3v3Training       string = "482" // nolint
	JoustAIEasy            string = "474" // nolint
	JoustAIMedium          string = "456" // nolint
	JoustChallenge         string = "441" // nolint
	JoustPracticeEasy      string = "464" // nolint
	JunglePracicePreselect string = "496" // nolint
	JunglePracitceMedium   string = "473" // nolint
	JunglePractice         string = "444" // nolint
	MatchOfTheDay          string = "434" // nolint
	Siege                  string = "459" // nolint
	SiegeChallenge         string = "460" // nolint

	// Paladins Match Queue IDs

	ClassicSiege                       string = "465" // nolint
	CustomOnslaughtForemansRise        string = "462" // nolint
	CustomOnslaughtMagistratesArchives string = "464" // nolint
	CustomOnslaughtPrimalCourt         string = "455" // nolint
	CustomOnslaughtSnowfallJunction    string = "454" // nolint
	CustomSiegeBrightmarsh             string = "458" // nolint
	CustomSiegeFishMarket              string = "431" // nolint
	CustomSiegeFrogIsle                string = "433" // nolint
	CustomSiegeFrozenGuard             string = "432" // nolint
	CustomSiegeIceMines                string = "439" // nolint
	CustomSiegeJaguarFalls             string = "438" // nolint
	CustomSiegeSerpentBeach            string = "440" // nolint
	CustomSiegeSplitstoneQuarry        string = "459" // nolint
	CustomSiegeStoneKeep               string = "423" // nolint
	CustomSiegeTimberMill              string = "430" // nolint
	CustomTDMForemansRise              string = "471" // nolint
	CustomTDMMagistrateArchives        string = "472" // nolint
	CustomTDMTradeDistrict             string = "468" // nolint
	LiveBattleGroundsDuo               string = "475" // nolint
	LiveBattleGroundsQuad              string = "476" // nolint
	LiveBattlegroundsSolo              string = "474" // nolint
	LiveCasual                         string = "424" // nolint
	LiveCompetitive                    string = "428" // nolint
	LiveOnslaught                      string = "452" // nolint
	LiveOnslaughtPractice              string = "453" // nolint
	LiveSiegePractice                  string = "425" // nolint
	LiveTeamDeathmatch                 string = "469" // nolint
	LiveTeamDeathmatchPractice         string = "470" // nolint
	LiveTestMaps                       string = "445" // nolint
)
