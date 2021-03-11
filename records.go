package godestone

import (
	"time"

	"github.com/karashiiro/godestone/data/baseparam"
	"github.com/karashiiro/godestone/data/gcrank"
	"github.com/karashiiro/godestone/data/gender"
	"github.com/karashiiro/godestone/data/role"
)

// NamedEntity represents an instance of an object with a name.
type NamedEntity struct {
	ID   uint32
	Name string
	Icon string

	NameEN string
	NameJA string
	NameDE string
	NameFR string
}

// GenderedEntity represents an instance of an object with masculine and feminine names.
type GenderedEntity struct {
	ID   uint32
	Name string

	NameMasculineEN string
	NameMasculineJA string
	NameMasculineDE string
	NameMasculineFR string
	NameFeminineEN  string
	NameFeminineJA  string
	NameFeminineDE  string
	NameFeminineFR  string
}

// PageInfo represents pagination information in a search.
type PageInfo struct {
	CurrentPage int
	TotalPages  int
}

// CrestLayers represents the layers of a community crest.
type CrestLayers struct {
	Bottom string
	Middle string
	Top    string
}

// AllAchievementInfo represents information about a character's achievements in aggregate.
type AllAchievementInfo struct {
	Private                bool
	TotalAchievements      uint32
	TotalAchievementPoints uint32
}

// AchievementInfo represents information about a character's achievements.
type AchievementInfo struct {
	Name string
	ID   uint32
	Date time.Time

	NameEN string
	NameJA string
	NameDE string
	NameFR string
}

// Title represents a character title.
type Title struct {
	*GenderedEntity

	Prefix bool
}

// Character represents the information available about a character on The Lodestone.
type Character struct {
	ActiveClassJob    *ClassJob
	Avatar            string
	Bio               string
	ClassJobs         []*ClassJob
	ClassJobBozjan    *ClassJobBozja
	ClassJobElemental *ClassJobEureka
	DC                string
	FreeCompanyID     string
	FreeCompanyName   string
	GearSet           *GearSet
	Gender            gender.Gender
	GrandCompanyInfo  *GrandCompanyInfo
	GuardianDeity     *NamedEntity
	ID                uint32
	Name              string
	Nameday           string
	ParseDate         time.Time
	Portrait          string
	PvPTeamID         string
	Race              *GenderedEntity
	Title             *Title
	Town              *NamedEntity
	Tribe             *GenderedEntity
	World             string
}

// CharacterSearchResult contains data from the character search page about a character.
type CharacterSearchResult struct {
	Error error `json:"-"`

	Avatar   string
	ID       uint32
	Lang     string
	Name     string
	Rank     gcrank.GCRank
	RankIcon string
	World    string
	DC       string
}

// ClassJob represents class and job information.
type ClassJob struct {
	ClassID       uint8
	ExpLevel      uint32
	ExpLevelMax   uint32
	ExpLevelTogo  uint32
	IsSpecialized bool
	JobID         uint8
	Level         uint8
	Name          string
	UnlockedState ClassJobUnlockedState
}

// ClassJobBozja represents character progression data in the Bozjan Southern Front.
type ClassJobBozja struct {
	Level  uint8
	Mettle uint32
	Name   string
}

// ClassJobEureka represents character progression data in Eureka.
type ClassJobEureka struct {
	ExpLevel     uint32
	ExpLevelMax  uint32
	ExpLevelTogo uint32
	Level        uint8
	Name         string
}

// ClassJobUnlockedState represents the unlock state of a ClassJob
type ClassJobUnlockedState struct {
	ID   uint8
	Name string
}

// CWLS represents basic CWLS information.
type CWLS struct {
	Name      string
	DC        string
	ID        string
	ParseDate time.Time
	Members   []*CWLSMember
}

// CWLSMember represents information about a CWLS member.
type CWLSMember struct {
	Avatar            string
	ID                uint32
	Name              string
	LinkshellRank     string
	LinkshellRankIcon string
	Rank              gcrank.GCRank
	RankIcon          string
	World             string
	DC                string
}

// CWLSSearchResult represents basic CWLS information returned from a search.
type CWLSSearchResult struct {
	Error error `json:"-"`

	Name          string
	ID            string
	DC            string
	ActiveMembers uint32
}

// Estate represents a housing area estate.
type Estate struct {
	Greeting string
	Name     string
	Plot     string
}

// GearItem represents information about a single gear item on a character.
type GearItem struct {
	Name    string
	Creator string
	Dye     uint32
	ID      uint32
	HQ      bool
	Materia []uint32
	Mirage  uint32

	NameEN string
	NameJA string
	NameDE string
	NameFR string
}

// GearItemBuild represents a full gearset on a character. All gear items can be nil.
type GearItemBuild struct {
	Body        *GearItem
	Bracelets   *GearItem
	Earrings    *GearItem
	Feet        *GearItem
	Hands       *GearItem
	Head        *GearItem
	Legs        *GearItem
	MainHand    *GearItem
	Necklace    *GearItem
	OffHand     *GearItem
	Ring1       *GearItem
	Ring2       *GearItem
	SoulCrystal *GearItem
	Waist       *GearItem
}

// GearSet represents the current gear information of a character.
type GearSet struct {
	Attributes map[baseparam.BaseParam]uint32
	ClassID    uint8
	Gear       *GearItemBuild
	GearKey    string
	JobID      uint8
	Level      uint8
}

// GrandCompanyInfo represents Grand Company information about a character.
type GrandCompanyInfo struct {
	GrandCompany *NamedEntity
	RankID       gcrank.GCRank
}

// FreeCompanyActiveState is the active state of an FC.
type FreeCompanyActiveState string

// Active state for an FC.
const (
	FCActiveNotSpecified FreeCompanyActiveState = "Not specified"
	FCActiveWeekdaysOnly FreeCompanyActiveState = "Weekdays Only"
	FCActiveWeekendsOnly FreeCompanyActiveState = "Weekends Only"
	FCActiveAlways       FreeCompanyActiveState = "Always"
)

// FreeCompanyRecruitingState is the recruiting state of an FC.
type FreeCompanyRecruitingState string

// Recruiting state for an FC.
const (
	FCRecruitmentClosed FreeCompanyRecruitingState = "Closed"
	FCRecruitmentOpen   FreeCompanyRecruitingState = "Open"
)

// FreeCompanyFocus is an FC's focus.
type FreeCompanyFocus string

// Free Company Focus.
const (
	FCFocusRolePlaying FreeCompanyFocus = "Role-playing"
	FCFocusLeveling    FreeCompanyFocus = "Leveling"
	FCFocusCasual      FreeCompanyFocus = "Casual"
	FCFocusHardcore    FreeCompanyFocus = "Hardcore"
	FCFocusDungeons    FreeCompanyFocus = "Dungeons"
	FCFocusGuildhests  FreeCompanyFocus = "Guildhests"
	FCFocusTrials      FreeCompanyFocus = "Trials"
	FCFocusRaids       FreeCompanyFocus = "Raids"
	FCFocusPvP         FreeCompanyFocus = "PvP"
)

// FreeCompanyFocusInfo represents a particular FC's intentions for a focus.
type FreeCompanyFocusInfo struct {
	Icon   string
	Kind   FreeCompanyFocus
	Status bool
}

// FreeCompanyRanking represents a particular FC's periodic rankings.
type FreeCompanyRanking struct {
	Monthly uint32
	Weekly  uint32
}

// FreeCompanyReputation represents an FC's alignment with each Grand Company.
type FreeCompanyReputation struct {
	GrandCompany *NamedEntity
	Progress     uint8
	Rank         *NamedEntity
}

// FreeCompanySeekingInfo represents a particular FC's intentions for a recruit roles.
type FreeCompanySeekingInfo struct {
	Icon   string
	Kind   role.Role
	Status bool
}

// FreeCompanyMember represents information about a Free Company member.
type FreeCompanyMember struct {
	Error error `json:"-"`

	Avatar   string
	ID       uint32
	Name     string
	Rank     gcrank.GCRank
	RankIcon string
	World    string
	DC       string
}

// FreeCompany represents all of the basic information about an FC.
type FreeCompany struct {
	Active            FreeCompanyActiveState
	ActiveMemberCount uint32
	CrestLayers       *CrestLayers
	DC                string
	Estate            *Estate
	Focus             []*FreeCompanyFocusInfo
	Formed            time.Time
	GrandCompany      *NamedEntity
	ID                string
	Name              string
	ParseDate         time.Time
	Rank              uint8
	Ranking           *FreeCompanyRanking
	Recruitment       FreeCompanyRecruitingState
	Reputation        []*FreeCompanyReputation
	Seeking           []*FreeCompanySeekingInfo
	Slogan            string
	Tag               string
	World             string
}

// FreeCompanySearchResult represents all of the searchable information about an FC.
type FreeCompanySearchResult struct {
	Error error `json:"-"`

	Active        FreeCompanyActiveState
	ActiveMembers uint32
	CrestLayers   *CrestLayers
	DC            string
	Estate        string
	Formed        time.Time
	GrandCompany  *NamedEntity
	ID            string
	Name          string
	Recruitment   FreeCompanyRecruitingState
	World         string
}

// Linkshell represents basic linkshell information.
type Linkshell struct {
	Name      string
	ID        string
	ParseDate time.Time
	Members   []*LinkshellMember
}

// LinkshellMember represents information about a linkshell member.
type LinkshellMember struct {
	Avatar            string
	ID                uint32
	Name              string
	LinkshellRank     string
	LinkshellRankIcon string
	Rank              gcrank.GCRank
	RankIcon          string
	World             string
	DC                string
}

// LinkshellSearchResult represents basic linkshell information returned from a search.
type LinkshellSearchResult struct {
	Error error `json:"-"`

	Name          string
	ID            string
	World         string
	DC            string
	ActiveMembers uint32
}

// Minion represents a minion.
type Minion struct {
	ID   uint32
	Name string
	Icon string

	NameEN string
	NameDE string
	NameFR string
	NameJA string
}

// Mount represents a mount.
type Mount struct {
	ID   uint32
	Name string
	Icon string

	NameEN string
	NameDE string
	NameFR string
	NameJA string
}

// PVPTeam represents information about a PVP team.
type PVPTeam struct {
	Name        string
	ID          string
	DC          string
	ParseDate   time.Time
	Formed      time.Time
	CrestLayers *CrestLayers
	Members     []*PVPTeamMember
}

// PVPTeamMember represents information about a PVP team member.
type PVPTeamMember struct {
	Avatar   string
	ID       uint32
	Name     string
	Matches  uint32
	Rank     gcrank.GCRank
	RankIcon string
	World    string
	DC       string
}

// PVPTeamSearchResult represents basic PVP team information returned from a search.
type PVPTeamSearchResult struct {
	Error error `json:"-"`

	Name        string
	ID          string
	DC          string
	CrestLayers *CrestLayers
}