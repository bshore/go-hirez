package utils

import (
	"math/rand"
	"strings"
	"time"
)

var trollIcon = "https://static.wikia.nocookie.net/smite_gamepedia/images/7/7f/Icon_Player_WhatANoob.png/revision/latest/scale-to-width-down/120"

var numbers = []string{"0","1","2","3","4","5","6","7","8","9"}

var names = []string{
	"Achilles", "Agni", "Ah Muzen Cab", "Ah Puch", "Amaterasu", "Anhur", "Anubis", "Ao Kuang",
	"Aphrodite", "Apollo", "Arachne", "Ares", "Artemis", "Artio", "Athena", "Awilix",
	"Baba Yaga", "Bacchus", "Bakasura", "Baron Samedi", "Bastet", "Bellona",
	"Cabrakan", "Camazotz", "Cerberus", "Cernunnos", "Chaac", "Chang'e", "Chernobog", "Chiron",
	"Chronos", "Cthulhu", "Cu Chulainn", "Cupid",
	"Da Ji", "Danzaburou", "Discordia",
	"Erlang Shen", "Eset",
	"Fafnir", "Fenrir", "Freya",
	"Ganesha", "Geb", "Gilgamesh", "Guan Yu",
	"Hachiman", "Hades", "He Bo", "Heimdallr", "Hel", "Hera", "Hercules", "Horus", "Hou Yi", "Hun Batz",
	"Izanami",
	"Janus", "Jing Wei", "Jormungandr",
	"Kali", "Khepri", "King Arthur", "Kukulkan", "Kumbhakarna", "Kuzenbo",
	"Loki",
	"Medusa", "Mercury", "Merlin", "Morgan Le Fay", "Mulan",
	"Ne Zha", "Neith", "Nemesis", "Nike", "Nox", "Nu Wa",
	"Odin", "Olorun", "Osiris",
	"Pele", "Persephone", "Poseidon",
	"Ra", "Raijin", "Rama",
	"Ratatoskr", "Ravana",
	"Scylla", "Serqet", "Set", "Skadi", "Sobek", "Sol", "Sun Wukong", "Susano", "Sylvanus",
	"Terra", "Thanatos", "The Morrigan", "Thor", "Thoth", "Tiamat", "Tsukuyomi", "Tyr",
	"Ullr",
	"Vamana", "Vulcan",
	"Xbalanque", "Xing Tian",
	"Yemoja", "Ymir",
	"Zeus", "Zhong Kui",
}

var actives = []string{
	"Aegis Amulet",	"Belt of Frenzy",	"Blink Rune",	"Bracer of Undoing", "Cursed Ankh",
	"Heavenly Wings",	"Horrific Emblem", "Magic Shell", "Meditation Cloak",	"Phantom Veil",
	"Purification Beads",	"Shield of Thorns",	"Sundering Spear", "Teleport Glyph",
}

var items = []string{
	"Ninja Tabi", "Reinforced Greaves", "Reinforced Shoes", "Shoes of the Magi", "Shoes of Focus", 
	"Talaria Boots", "Traveler's Shoes", "Warrior Tabi", "Relic Dagger", "Winged Blade", 
	"Emperor's Armor", "Golden Blade", "Shield of Regrowth", "Bristlebush Acorn", "Contagion", 
	"Evergreen Acorn", "Lotus Crown", "Thickbark Acorn", "Thistlethorn Acorn", "Witchblade", 
	"Heartward Amulet", "Odysseus' Bow", "Oni Hunter's Garb", "Soul Eater", "Sovereignty", 
	"Spectral Armor", "Celestial Legion Helm", "Hydra's Lament", "Jade Emperor's Crown", "Magi's Cloak",
	"Runic Shield", "Talisman of Energy", "Atalanta's Bow", "Blackthorn Hammer", "Genji's Guard",
	"Hide of the Nemean Lion", "Shogun's Kusari", "Silverbranch Bow", "The Executioner", "Toxic Blade",
	"Ancile", "Pestilence", "Stone of Gaia", "Void Stone", "Breastplate of Valor", "Bulwark of Hope",
	"Demonic Grip", "Devourer's Gauntlet", "Divine Ruin", "Frostbound Hammer", "Lono's Mask",
	"Midgardian Mail", "Poisoned Star", "Polynomicon", "Pythagorem's Piece", "Soul Gem", "The Sledge",
	"Brawler's Beat Stick", "Gauntlet of Thebes", "Jotunn's Wrath", "Mystical Mail", "Shifter's Shield",
	"Stone Cutting Sword", "Stone of Fal", "Charon's Coin", "Gladiator's Shield", "Mail of Renewal",
	"Pridwen", "Runeforged Hammer", "Shadowsteel Shuriken", "The Crusher", "Hide of the Urchin", "Ichaival",
	"Transcendence", "Bancroft's Talon", "Book of Thoth", "Caduceus Shield", "Dominance", "Hastened Katana",
	"Rage", "Rangda's Mask", "Ring of Hecate", "Spirit Robe", "Warlock's Staff", "Wind Demon", "Asi",
	"Obsidian Shard", "Spear of the Magus", "Titan's Bane", "Arondight", "Berserker's Shield",
	"Book of the Dead", "Ethereal Staff", "Rod of Asclepius", "Spear of Desolation", "Tyrannical Plate Helm",
	"Void Shield", "Malice", "Serrated Edge", "Soul Reaver", "Chronos' Pendant", "Gem of Isolation", 
	"Hastened Ring", "Qin's Sais", "Telkhines Ring", "Bloodforge", "Fail-not", "Staff of Myrddin",
	"Typhon's Fang", "Doom Orb", "Heartseeker", "Mantle of Discord", "Deathbringer", "Rod of Tahuti",
}

func generateString(fieldName string) string {
	f := strings.ToLower(fieldName)
	rand.Seed(time.Now().Add(time.Duration(rand.Int63()) * time.Millisecond).Unix())
	if f == "" {
		return "some string"
	}
	if strings.Contains(f, "avatar") || strings.Contains(f, "url") {
		return trollIcon
	}
	if strings.Contains(f, "datetime") {
		return FormatTime(time.Now())
	}
	if strings.Contains(f, "id") {
		return numbers[rand.Intn(len(numbers))]
	}
	if strings.Contains(f, "name") || strings.Contains(f, "god") || strings.Contains(f, "ban") {
		return names[rand.Intn(len(names))]
	}
	if strings.Contains(f, "active") {
		return actives[rand.Intn(len(actives))]
	}
	if strings.Contains(f, "item") {
		return items[rand.Intn(len(items))]
	}
	if strings.Contains(f, "summary") || strings.Contains(f, "description"){
		return `many many words many many words many many words many many words many many words
		many many words many many words many many words many many words many many words many many words
		many many words many many words many many words many many words many many words many many words
		many many words many many words many many words many many words many many words many many words
		many many words many many words many many words many many words many many words many many words
		many many words many many words many many words many many words many many words many many words
		`
	}
	return "some string"
}

