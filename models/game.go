package models

// God stores data related to an in game God
type God struct {
	Ability1                   string                 `json:"Ability1"`
	Ability2                   string                 `json:"Ability2"`
	Ability3                   string                 `json:"Ability3"`
	Ability4                   string                 `json:"Ability4"`
	Ability5                   string                 `json:"Ability5"`
	AbilityId1                 int64                  `json:"AbilityId1"`
	AbilityId2                 int64                  `json:"AbilityId2"`
	AbilityId3                 int64                  `json:"AbilityId3"`
	AbilityId4                 int64                  `json:"AbilityId4"`
	AbilityId5                 int64                  `json:"AbilityId5"`
	AbilityOne                 GodAbility             `json:"Ability_1"`
	AbilityTwo                 GodAbility             `json:"Ability_2"`
	AbilityThree               GodAbility             `json:"Ability_3"`
	AbilityFour                GodAbility             `json:"Ability_4"`
	AbilityFive                GodAbility             `json:"Ability_5"`
	AttackSpeed                float32                `json:"AttackSpeed"`
	AttackSpeedPerLevel        float32                `json:"AttackSpeedPerLevel"`
	AutoBanned                 string                 `json:"AutoBanned"`
	Cons                       string                 `json:"Cons"`
	HP5PerLevel                float32                `json:"HP5PerLevel"`
	Health                     int64                  `json:"Health"`
	HealthPerFive              int64                  `json:"HealthPerFive"`
	HealthPerLevel             int64                  `json:"HealthPerLevel"`
	Lore                       string                 `json:"Lore"`
	MP5PerLevel                float32                `json:"MP5PerLevel"`
	MagicProtection            float32                `json:"MagicProtection"`
	MagicProtectionPerLevel    float32                `json:"MagicProtectionPerLevel"`
	MagicalPower               float32                `json:"MagicalPower"`
	MagicalPowerPerLevel       float32                `json:"MagicalPowerPerLevel"`
	Mana                       int64                  `json:"Mana"`
	ManaPerFive                float32                `json:"ManaPerFive"`
	ManaPerLevel               int64                  `json:"ManaPerLevel"`
	Name                       string                 `json:"Name"`
	OnFreeRotation             string                 `json:"OnFreeRotation"`
	Pantheon                   string                 `json:"Pantheon"`
	PhysicalPower              float32                `json:"PhysicalPower"`
	PhysicalPowerPerLevel      float32                `json:"PhysicalPowerPerLevel"`
	PhysicalProtection         float32                `json:"PhysicalProtection"`
	PhysicalProtectionPerLevel float32                `json:"PhysicalProtectionPerLevel"`
	Pros                       string                 `json:"Pros"`
	Roles                      string                 `json:"Roles"`
	Speed                      int64                  `json:"Speed"`
	Title                      string                 `json:"Title"`
	Type                       string                 `json:"Type"`
	AbilityDescription1        AbilityItemDescription `json:"abilityDescription1"`
	AbilityDescription2        AbilityItemDescription `json:"abilityDescription2"`
	AbilityDescription3        AbilityItemDescription `json:"abilityDescription3"`
	AbilityDescription4        AbilityItemDescription `json:"abilityDescription4"`
	AbilityDescription5        AbilityItemDescription `json:"abilityDescription5"`
	BasicAttack                GodBasicAttack         `json:"basicAttack"`
	GodAbility1URL             string                 `json:"godAbility1_URL"`
	GodAbility2URL             string                 `json:"godAbility2_URL"`
	GodAbility3URL             string                 `json:"godAbility3_URL"`
	GodAbility4URL             string                 `json:"godAbility4_URL"`
	GodAbility5URL             string                 `json:"godAbility5_URL"`
	GodCardURL                 string                 `json:"godCard_URL"`
	GodIconURL                 string                 `json:"godIcon_URL"`
	ID                         int64                  `json:"id"`
	LatestGod                  string                 `json:"latestGod"`
	RetMsg                     string                 `json:"ret_msg"`
}

// GodAbility stores data related to a God Ability
type GodAbility struct {
	Description AbilityItemDescription `json:"Description"`
	ID          int64                  `json:"Id"`
	Summary     string                 `json:"Summary"`
	URL         string                 `json:"URL"`
}

// GodAltAbility stores data related to a God's Alternate Ability
type GodAltAbility struct {
	AltName     string `json:"alt_name"`
	AltPosition string `json:"alt_position"`
	God         string `json:"God"`
	GodID       int64  `json:"god_id"`
	ItemID      int64  `json:"item_id"`
	RetMsg      string `json:"ret_msg"`
}

// GodBasicAttack stores data related to a God's Basic Attack Ability
type GodBasicAttack struct {
	ItemDescription AbilityItemDescription `json:"itemDescription"`
}

// AbilityItemDescription stores data related to a God's Ability
type AbilityItemDescription struct {
	ItemDescription AbilityDescription `json:"itemDescription"`
}

// AbilityDescription stores data related to a God's Ability
type AbilityDescription struct {
	Cooldown    string                        `json:"cooldown"`
	Cost        string                        `json:"cost"`
	Description string                        `json:"description"`
	MenuItems   []AbilityDescriptionValueItem `json:"menuitems"`
	RankItems   []AbilityDescriptionValueItem `json:"rankitems"`
}

// AbilityDescriptionValueItem stores data related to the values of a God's Ability
type AbilityDescriptionValueItem struct {
	Description string `json:"description"`
	Value       string `json:"value"`
}

// GodSkin stores data related to a Skin for a God
type GodSkin struct {
	GodIconURL    string `json:"godIcon_URL"`
	GodSkinURL    string `json:"godSkin_URL"`
	GodID         int64  `json:"god_id"`
	GodName       string `json:"god_name"`
	Obtainability string `json:"obtainability"`
	PriceFavor    int64  `json:"price_favor"`
	PriceGems     int64  `json:"price_gems"`
	RetMsg        string `json:"ret_msg"`
	SkinID1       int64  `json:"skin_id1"`
	SkinID2       int64  `json:"skin_id2"`
	SkinName      string `json:"skin_name"`
}

// GodRecommendedItem stores data related to a Recommended Item to build for a God
type GodRecommendedItem struct {
	Category        string `json:"Category"`
	Item            string `json:"Item"`
	Role            string `json:"Role"`
	CategoryValueID int64  `json:"category_value_id"`
	GodID           int64  `json:"god_id"`
	GodName         string `json:"god_name"`
	IconID          int64  `json:"icon_id"`
	ItemID          int64  `json:"item_id"`
	RetMsg          string `json:"ret_msg"`
	RoleValueID     int64  `json:"role_value_id"`
}

// Item stores data related to an in game Item
type Item struct {
	ActiveFlag      string          `json:"ActiveFlag"`
	ChildItemID     int64           `json:"ChildItemId"`
	DeviceName      string          `json:"DeviceName"`
	IconID          int64           `json:"IconId"`
	ItemDescription ItemDescription `json:"ItemDescription"`
	ItemID          int64           `json:"ItemId"`
	ItemTier        int             `json:"ItemTier"`
	Price           int64           `json:"Price"`
	RestrictedRoles string          `json:"RestrictedRoles"`
	RootItemID      int64           `json:"RootItemId"`
	ShortDesc       string          `json:"ShortDesc"`
	StartingItem    bool            `json:"StartingItem"`
	Type            string          `json:"Type"`
	ItemIconURL     string          `json:"itemIcon_URL"`
	RetMsg          string          `json:"ret_msg"`
}

// ItemDescription stores data related to an Item's Description
type ItemDescription struct {
	Description          string                     `json:"Description"`
	MenuItems            []ItemDescriptionValueItem `json:"Menuitems"`
	SecondaryDescription string                     `json:"SecondaryDescription"`
}

// ItemDescriptionValueItem stores data related to the Value of an Item stat
type ItemDescriptionValueItem struct {
	Description string `json:"Description"`
	Value       string `json:"value"`
}
