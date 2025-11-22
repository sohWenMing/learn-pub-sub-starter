package gamelogic

// player just has a username, and a map of units which maps the id of the unit
// to the Unit
type Player struct {
	Username string
	Units    map[int]Unit
}

// UnitRank is a string that is associated with the Unit struct
type UnitRank string

const (
	RankInfantry  = "infantry"
	RankCavalry   = "cavalry"
	RankArtillery = "artillery"
)

type Unit struct {
	ID       int
	Rank     UnitRank
	Location Location
}

// ArmyMove is a struct that houses all the units that a player is moving to a location
type ArmyMove struct {
	Player     Player
	Units      []Unit
	ToLocation Location
}

// RecognitionOfWar will record which two players are engaged in a "fight"
type RecognitionOfWar struct {
	Attacker Player
	Defender Player
}

type Location string

func getAllRanks() map[UnitRank]struct{} {
	return map[UnitRank]struct{}{
		RankInfantry:  {},
		RankCavalry:   {},
		RankArtillery: {},
	}
}

func getAllLocations() map[Location]struct{} {
	return map[Location]struct{}{
		"americas":   {},
		"europe":     {},
		"africa":     {},
		"asia":       {},
		"australia":  {},
		"antarctica": {},
	}
}
