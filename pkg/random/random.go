package random

import (
	"errors"
	"fmt"
	"math/big"
	"math/rand"
	"strings"
	"time"
)

func RandomStringFromGiven(options []string) string {
	if len(options) == 0 {
		return ""
	}

	randomIndex := rand.Intn(len(options))

	return options[randomIndex]
}

func RandomIntFromRange(min, max int) int {
	if min > max {
		min, max = max, min
	}

	rangeSize := max - min + 1

	return min + rand.Intn(rangeSize)
}

func RandomTimestamp(start, end time.Time) time.Time {
	if start.After(end) {
		start, end = end, start
	}

	delta := end.Unix() - start.Unix()

	if delta == 0 {
		return start
	}

	randomDelta := rand.Int63n(delta)

	return start.Add(time.Duration(randomDelta) * time.Second)
}

func RandomNickname(walletAddress string) (string, error) {
	if !strings.HasPrefix(walletAddress, "0x") || len(walletAddress) != 42 {
		return "", errors.New("invalid wallet address format")
	}

	address := strings.ToLower(walletAddress[2:])

	firstPart := address[4:12]
	lastPart := address[24:32]

	firstNum := new(big.Int)
	firstNum.SetString(firstPart, 16)

	lastNum := new(big.Int)
	lastNum.SetString(lastPart, 16)

	prefixes := []string{
		"Ace", "Arc", "Ash", "Aura", "Axe", "Blaze", "Bolt", "Boss", "Buzz", "Byte",
		"Cap", "Chief", "Claw", "Comet", "Core", "Crash", "Cryo", "Cyber", "Dash", "Data",
		"Doc", "Doom", "Dusk", "Echo", "Edge", "Fang", "Flux", "Frost", "Fury", "Gale",
		"Gem", "Glitch", "Glow", "Grit", "Guard", "Halo", "Havoc", "Hawk", "Hex", "Ice",
		"Ion", "Iris", "Jet", "Jinx", "Jolt", "Jugg", "Kal", "Khan", "Kite", "Knight",
		"Lark", "Lens", "Link", "Lode", "Lore", "Lynx", "Mace", "Max", "Maze", "Moon",
		"Myth", "Neo", "Nero", "Nexus", "Night", "Nova", "Nuke", "Null", "Oath", "Onyx",
		"Orb", "Pace", "Pax", "Peak", "Phoenix", "Pix", "Plasma", "Prime", "Prism", "Pulse",
		"Quake", "Quartz", "Quick", "Quill", "Raid", "Rain", "Raven", "Razor", "Rebel", "Red",
		"Relic", "Rex", "Riff", "Rift", "Riot", "Rook", "Root", "Rune", "Rush", "Rust",
		"Sage", "Scout", "Shade", "Shadow", "Sharp", "Shift", "Shine", "Shock", "Silk", "Silver",
		"Slate", "Sling", "Slick", "Spark", "Specter", "Spike", "Spin", "Splash", "Star", "Storm",
		"Streak", "Strike", "Surge", "Swift", "Sync", "Tac", "Tech", "Thorne", "Thunder", "Tide",
		"Tiger", "Titan", "Trek", "Tron", "Tundra", "Venom", "Verge", "Vex", "Vibe", "Void",
		"Volt", "Vortex", "Warp", "Wave", "Whirl", "Wise", "Wrath", "Xeno", "Xtreme", "Zap",
		"Zen", "Zero", "Zest", "Zinc", "Zion", "Zone", "Zoom", "Zulu", "Arc", "Atom",
	}

	suffixes := []string{
		"Alpha", "Apex", "Arrow", "Astro", "Aurora", "Blade", "Blazer", "Bot", "Breaker", "Burst",
		"Byte", "Champion", "Chaos", "Claw", "Coast", "Core", "Crypto", "Curator", "Dagger", "Dawn",
		"Dazzle", "Delta", "Drift", "Edge", "Ember", "Enigma", "Factor", "Fang", "Flash", "Flare",
		"Flux", "Force", "Fox", "Fragment", "Frost", "Fuse", "Fury", "Ghost", "Glider", "Glimmer",
		"Grind", "Guardian", "Hammer", "Havoc", "Hawk", "Heart", "Hex", "Hunter", "Impact", "Impulse",
		"Infinity", "Jester", "Judge", "Jumper", "Justice", "King", "Knight", "Legend", "Lightning", "Lore",
		"Lynx", "Machine", "Mage", "Major", "Mask", "Master", "Matrix", "Maven", "Mirage", "Mist",
		"Myth", "Nexus", "Night", "Ninja", "Noble", "Nova", "Omega", "Oracle", "Orbit", "Origin",
		"Phantom", "Phase", "Pilot", "Pioneer", "Pixel", "Plasma", "Prime", "Prophet", "Pulse", "Quasar",
		"Quest", "Racer", "Raider", "Ranger", "Raven", "Reaper", "Rebel", "Rex", "Rider", "Rift",
		"Rogue", "Runner", "Sage", "Seeker", "Sentinel", "Shadow", "Shatter", "Shield", "Shift", "Shock",
		"Slayer", "Sniper", "Specter", "Spike", "Spirit", "Stalker", "Star", "Stealth", "Storm", "Striker",
		"Surge", "Tamer", "Tech", "Tempest", "Thunder", "Titan", "Tracer", "Tracker", "Vanguard", "Vector",
		"Venom", "Vertex", "Vision", "Void", "Vortex", "Walker", "Warden", "Warrior", "Wave", "Whisper",
		"Wing", "Wink", "Winter", "Wolf", "Wonder", "Wraith", "Wrath", "X", "Z", "Zero",
		"Zodiac", "Zone", "Brain", "Byte", "Code", "Engine", "Key", "Mind", "Protocol", "Soul",
	}

	lenPrefixes := big.NewInt(int64(len(prefixes)))
	prefixIndex := new(big.Int).Mod(firstNum, lenPrefixes).Int64()
	prefix := prefixes[prefixIndex]

	lenSuffixes := big.NewInt(int64(len(suffixes)))
	suffixIndex := new(big.Int).Mod(lastNum, lenSuffixes).Int64()
	suffix := suffixes[suffixIndex]

	return fmt.Sprintf("%s%s", prefix, suffix), nil
}

func RandomTeamName(walletAddress string) (string, error) {
	if !strings.HasPrefix(walletAddress, "0x") || len(walletAddress) != 42 {
		return "", errors.New("invalid wallet address format")
	}

	address := strings.ToLower(walletAddress[2:])

	firstPart := address[0:8]
	middlePart := address[16:24]
	lastPart := address[32:40]

	firstNum := new(big.Int)
	firstNum.SetString(firstPart, 16)

	middleNum := new(big.Int)
	middleNum.SetString(middlePart, 16)

	lastNum := new(big.Int)
	lastNum.SetString(lastPart, 16)

	adjectives := []string{
		"Blazing", "Mighty", "Phantom", "Mystic", "Quantum", "Savage", "Titan", "Cosmic",
		"Valiant", "Shadow", "Crystal", "Thunder", "Golden", "Emerald", "Carbon", "Lunar",
		"Solar", "Steel", "Diamond", "Royal", "Crimson", "Noble", "Astral", "Fierce",
		"Raging", "Ancient", "Arcane", "Frozen", "Burning", "Electric", "Toxic", "Radiant",
		"Shining", "Cyber", "Hyper", "Alpha", "Omega", "Supreme", "Ultimate", "Phoenix",
		"Eternal", "Prime", "Nebula", "Azure", "Cobalt", "Jade", "Obsidian", "Onyx",
		"Ruby", "Sapphire", "Scarlet", "Silver", "Void", "Wild", "Iron", "Bronze",
		"Brass", "Copper", "Platinum", "Neon", "Plasma", "Core", "Primal", "Chaos",
		"Celestial", "Spectral", "Cursed", "Blessed", "Divine", "Heroic", "Legendary",
		"Mythic", "Epic", "Hidden", "Veiled", "Covert", "Secret", "Silent", "Howling",
		"Roaring", "Screaming", "Whispering", "Unstoppable", "Invincible", "Immortal",
		"Ethereal", "Elemental", "Temporal", "Spatial", "Dimensional", "Gravitational",
		"Atomic", "Molecular", "Kinetic", "Potential", "Thermal", "Magnetic", "Digital",
		"Analog", "Viral", "Bionic", "Mecha", "Techno", "Psychic", "Psionic", "Neural",
		"Virtual", "Augmented", "Enhanced", "Modified", "Upgraded", "Evolved", "Ascended",
		"Transcendent", "Enlightened", "Awakened", "Revenant", "Undying", "Deathless",
		"Imperishable", "Indomitable", "Resolute", "Steadfast", "Tenacious", "Unyielding",
		"Vorpal", "Zealous", "Aberrant", "Abyssal", "Bizarre", "Cryptic", "Dreadful",
		"Eldritch", "Forbidden", "Grotesque", "Hallowed", "Infernal", "Jovial", "Killer",
	}

	nouns := []string{
		"Dragons", "Warriors", "Wolves", "Eagles", "Knights", "Guardians", "Titans", "Legends",
		"Phoenixes", "Panthers", "Vipers", "Hawks", "Ninjas", "Raiders", "Voyagers", "Spartans",
		"Pythons", "Falcons", "Hunters", "Giants", "Demons", "Ghosts", "Lions", "Oracles",
		"Hammers", "Shields", "Swords", "Arrows", "Axes", "Spears", "Maces", "Daggers",
		"Claws", "Fangs", "Talons", "Horns", "Wings", "Scales", "Shells", "Tentacles",
		"Minds", "Souls", "Spirits", "Wraiths", "Specters", "Banshees", "Ghouls", "Zombies",
		"Vampires", "Werewolves", "Cyclops", "Chimeras", "Griffins", "Hydras", "Krakens", "Leviathans",
		"Minotaurs", "Pegasi", "Unicorns", "Wyverns", "Basilisks", "Cockatrices", "Manticores", "Sphinx",
		"Trolls", "Goblins", "Orcs", "Ogres", "Golems", "Gargoyles", "Elementals", "Djinn",
		"Wizards", "Warlocks", "Mages", "Sorcerers", "Witches", "Shamans", "Druids", "Clerics",
		"Paladins", "Crusaders", "Templars", "Berserkers", "Assassins", "Rogues", "Rangers", "Snipers",
		"Commandos", "Marines", "Troopers", "Soldiers", "Pilots", "Captains", "Admirals", "Generals",
		"Overlords", "Warlords", "Emperors", "Monarchs", "Tyrants", "Dictators", "Sovereigns", "Regents",
		"Barons", "Counts", "Dukes", "Earls", "Lords", "Nobles", "Samurai", "Ronin",
		"Monks", "Acolytes", "Disciples", "Zealots", "Cultists", "Fanatics", "Prophets", "Seers",
		"Mystics", "Shamans", "Sages", "Scholars", "Alchemists", "Scientists", "Engineers", "Builders",
		"Creators", "Destroyers", "Defenders", "Protectors", "Watchers", "Sentinels", "Wardens", "Keepers",
		"Seekers", "Wanderers", "Nomads", "Pilgrims", "Outcasts", "Exiles", "Outlaws", "Bandits",
		"Pirates", "Corsairs", "Marauders", "Plunderers", "Reavers", "Scavengers", "Vultures", "Ravens",
		"Crows", "Owls", "Bats", "Sharks", "Tigers", "Bears", "Rhinos", "Scorpions",
	}

	lenAdjectives := big.NewInt(int64(len(adjectives)))
	adjIndex := new(big.Int).Mod(firstNum, lenAdjectives).Int64()
	adjective := adjectives[adjIndex]

	lenNouns := big.NewInt(int64(len(nouns)))
	nounIndex := new(big.Int).Mod(lastNum, lenNouns).Int64()
	noun := nouns[nounIndex]

	suffixMod := big.NewInt(1000)
	suffix := new(big.Int).Mod(middleNum, suffixMod).Int64()

	return fmt.Sprintf("%s %s %d", adjective, noun, suffix), nil
}
