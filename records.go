package main

import "time"

/*
	       A0-File description rec      first record in file, contains
					    data on the file itself and
					    identifies the vendor software.
		 B1-Meet rec                one per file, contains data on
					    type of meet, location, dates
					    of competition.
		 B2-Meet host rec           one or more per meet, contains
					    meet host contact information
		    C1-Team ID rec          one per team, contains data on
					    team name, code, address.
		    C2-Team entry rec       one per team, contains coach
					    info plus stats on other
					    records to follow for that team.
		       D0-Ind. event rec    one per splash, contains data
					    on athlete, event, stroke and
					    distance, times, places
					    and lanes.
		       D3-Ind. Info rec     one per swimmer, a swimmer with
					    multiple D0 records will have
					    one D3 record following his/her
					    first D0 record, contains new
					    USS# and prefered first name.
			  G0-Split rec      one or more per "D0", contains
					    data on athlete name, ID, event
					    ID, split time and interval.
		       E0-Relay event rec   one per relay entry, contains
					    data on relay event, team,
					    times, places and lanes.
		       F0-Relay name rec    four or more per "E0", contains
					    data on the athlete name, ID,
					    time, split, and session.
		       D3-Ind. Info rec     one per "F0".
			  G0-Split rec      one or more per "E0"
	       Z0-File terminator rec       one per file, contains
					    statistics and text information.

*/

type MeetHostRecord struct {
	OrgCode     Code   `sdif:"start=3,len=1,code=001"`
	HostName    string `sdif:"start=12,len=30"`
	AddrLineOne string `sdif:"start=42,len=22"`
	AddrLineTwo string `sdif:"start=64,len=22"`
	City        string `sdif:"start=86,len=20"`
	State       USPS   `sdif:"start=106,len=2"`
	PostalCode  string `sdif:"start=108,len=10"`
	Country     Code   `sdif:"start=118,len=3,code=004"`
	Phone       string `sdif:"start=121,len=12"`
}

func (r MeetHostRecord) ConstTag() string { return "B2" }

type TeamIdRecord struct {
	OrgCode         Code   `sdif:"start=3,len=1,code=001"`
	TeamCode        Code   `sdif:"start=12,len=6,code=006"`
	FullTeamName    string `sdif:"start=18,len=30"`
	AbbrvTeamName   string `sdif:"start=48,len=16"`
	AddrLineOne     string `sdif:"start=64,len=22"`
	AddrLineTwo     string `sdif:"start=86,len=22"`
	City            string `sdif:"start=108,len=20"`
	State           USPS   `sdif:"start=128,len=2"`
	PostalCode      string `sdif:"start=130,len=10"`
	Country         Code   `sdif:"start=140,len=3,code=004"`
	REGION          Code   `sdif:"start=143,len=1,code=007"`
	OptTeamCodeChar string `sdif:"start=150,len=1"`
}

func (r TeamIdRecord) ConstTag() string { return "C1" }

type TeamEntryRecord struct {
	OrgCode          Code   `sdif:"start=3,len=1,code=001"`
	TeamCode         Code   `sdif:"start=12,len=6,code=006"`
	CoachName        string `sdif:"start=18,len=30"`
	CoachPhone       string `sdif:"start=48,len=12"`
	NumEntries       int    `sdif:"start=60,len=6"`
	NumAthletes      int    `sdif:"start=66,len=6"`
	NumRelays        int    `sdif:"start=72,len=5"`
	NumRelaySwimmers int    `sdif:"start=77,len=6"`
	NumRelaySplits   int    `sdif:"start=83,len=6"`
	ShortTeamName    string `sdif:"start=89,len=16"`
	OptTeamCodeChar  string `sdif:"start=150,len=1"`
}

func (r TeamEntryRecord) ConstTag() string { return "C2" }

type IndividualEventRecord struct {
	OrgCode             Code          `sdif:"start=3,len=1,code=001"`
	SwimmerName         string        `sdif:"start=12,len=28"`
	USS                 string        `sdif:"start=40,len=12"`
	AttachCode          Code          `sdif:"start=52,len=1,code=016"`
	CitizenCode         Code          `sdif:"start=53,len=3,code=009"`
	SwimmerDOB          time.Time     `sdif:"start=56,len=8"`
	SwimmerAge          string        `sdif:"start=64,len=2"`
	SexCode             Code          `sdif:"start=66,len=1,code=010"`
	EventCode           Code          `sdif:"start=67,len=1,code=011"`
	EventDistance       int           `sdif:"start=68,len=4"`
	StrokeCode          Code          `sdif:"start=72,len=1,code=012"`
	EventAgeCode        Code          `sdif:"start=77,len=4,code=025"`
	DateOfSwim          time.Time     `sdif:"start=81,len=8"`
	SeedTime            time.Duration `sdif:"start=89,len=8"`
	SeedCourseCode      Code          `sdif:"start=97,len=1,code=013"`
	PrelimTime          time.Duration `sdif:"start=98,len=8"`
	PrelimCourseCode    Code          `sdif:"start=106,len=1,code=013"`
	SwimOffTime         time.Duration `sdif:"start=107,len=8"`
	SwimOffCourseCode   Code          `sdif:"start=115,len=1,code=013"`
	FinalsTime          time.Duration `sdif:"start=116,len=8"`
	FinalsCourseCode    Code          `sdif:"start=124,len=1,code=013"`
	PrelimHeat          int           `sdif:"start=125,len=2"`
	PrelimLane          int           `sdif:"start=127,len=2"`
	FinalsHeat          int           `sdif:"start=129,len=2"`
	FinalsLane          int           `sdif:"start=131,len=2"`
	PrelimRanking       int           `sdif:"start=133,len=3"`
	FinalsRanking       int           `sdif:"start=136,len=3"`
	FinalsPointsScored  float32       `sdif:"start=139,len=4"`
	EventTimeClassCode  Code          `sdif:"start=143,len=2,code=014"`
	SwimmerFlightStatus string        `sdif:"start=145,len=1"`
}

func (r IndividualEventRecord) ConstTag() string { return "D0" }

type IndividualContactRecord struct {
	OrgCode         Code   `sdif:"start=3,len=1,code=001"`
	TeamCode        Code   `sdif:"start=12,len=6,code=006"`
	OptTeamCodeChar string `sdif:"start=18,len=1"`
	SwimmerName     string `sdif:"start=19,len=28"`
	AltMailingName  string `sdif:"start=47,len=30"`
	MailingAddrSt   string `sdif:"start=77,len=30"`
	MailingCity     string `sdif:"start=107,len=20"`
	MailingState    USPS   `sdif:"start=127,len=2"`
	MailingCountry  string `sdif:"start=129,len=12"`
	PostalCode      string `sdif:"start=141,len=10"`
	CountryCode     Code   `sdif:"start=151,len=3,code=004"`
	ReigonCode      Code   `sdif:"start=154,len=1,code=007"`
	AnswerCode      Code   `sdif:"start=155,len=1,code=023"`
	SeasonCode      Code   `sdif:"start=156,len=1,code=022"`
}

func (r IndividualContactRecord) ConstTag() string { return "D2" }

type IndividualInfoRecord struct {
	USS                string `sdif:"start=3,len=14"`
	PrefFirstName      string `sdif:"start=17,len=15"`
	EthnicityCode      Code   `sdif:"start=32,len=2,code=026"`
	JuniorHighSchool   bool   `sdif:"start=34,len=1"`
	SeniorHighSchool   bool   `sdif:"start=35,len=1"`
	YMCAOrYWCA         bool   `SDIF:"START=36,LEN=1"`
	College            bool   `sdif:"start=37,len=1"`
	ParkandRec         bool   `sdif:"start=38,len=1"`
	SummerLeague       bool   `sdif:"start=39,len=1"`
	CountryClub        bool   `sdif:"start=40,len=1"`
	Masters            bool   `sdif:"start=41,len=1"`
	DisabledSportsOrgs bool   `sdif:"start=42,len=1"`
	WaterPolo          bool   `sdif:"start=43,len=1"`
}

func (r IndividualInfoRecord) ConstTag() string { return "D3" }

type RelayEventRecord struct {
	OrgCode            Code          `sdif:"start=3,len=1,code=001"`
	RelayTeamName      string        `sdif:"start=12,len=1"`
	TeamCode           Code          `sdif:"start=13,len=6,code=006"`
	NumRelayNameRecs   int           `sdif:"start=19,len=2"`
	SexCode            Code          `sdif:"start=21,len=1,code=011"`
	RelayDistance      int           `sdif:"start=22,len=4"`
	StrokeCode         Code          `sdif:"start=26,len=1,code=012"`
	EventAgeCode       Code          `sdif:"start=31,len=4,code=025"`
	AthletesTotalAge   int           `sdif:"start=35,len=3"`
	DateOfSwim         time.Time     `sdif:"start=38,len=8"`
	SeedTime           time.Duration `sdif:"start=46,len=8"`
	SeedCourseCode     Code          `sdif:"start=54,len=1,code=013"`
	PrelimTime         time.Duration `sdif:"start=55,len=8"`
	PrelimCourseCode   Code          `sdif:"start=63,len=1,code=013"`
	SwimOffTime        time.Duration `sdif:"start=64,len=8"`
	SwimOffCourseCode  Code          `sdif:"start=72,len=1,code=013"`
	FinalsTime         time.Duration `sdif:"start=73,len=8"`
	FinalsCourseCode   Code          `sdif:"start=81,len=1,code=013"`
	PrelimHeat         int           `sdif:"start=82,len=2"`
	PrelimLane         int           `sdif:"start=84,len=2"`
	FinalsHeat         int           `sdif:"start=86,len=2"`
	FinalsLane         int           `sdif:"start=88,len=2"`
	PrelimRanking      int           `sdif:"start=90,len=3"`
	FinalsRanking      int           `sdif:"start=93,len=3"`
	FinalsPointsScored float32       `sdif:"start=96,len=4"`
	EventTimeClassCode Code          `sdif:"start=100,len=2,code=014"`
}

func (r RelayEventRecord) ConstTag() string { return "E0" }

type RelayNameRecord struct {
	OrgCode             Code          `sdif:"start=3,len=1,code=001"`
	TeamCode            Code          `sdif:"start=16,len=6,code=006"`
	RelayTeamName       string        `sdif:"start=22,len=1"`
	SwimmerName         string        `sdif:"start=23,len=28"`
	USS                 string        `sdif:"start=51,len=12"`
	CitizenCode         Code          `sdif:"start=63,len=3,code=009"`
	SwimmerDOB          time.Time     `sdif:"start=66,len=8"`
	SwimmerAge          string        `sdif:"start=74,len=2"`
	SexCode             Code          `sdif:"start=76,len=1,code=010"`
	PrelimLegOrderCode  Code          `sdif:"start=77,len=1,code=024"`
	SwimOffLegOrderCode Code          `sdif:"start=78,len=1,code=024"`
	FinalsLegOrderCode  Code          `sdif:"start=79,len=1,code=024"`
	LegTime             time.Duration `sdif:"start=80,len=8"`
	LegCourseCode       Code          `sdif:"start=88,len=1,code=13"`
	AutoTakeOffTime     float32       `sdif:"start=89,len=4"`
}

func (r RelayNameRecord) ConstTag() string { return "F0" }

type SplitsRecord struct{
	
	OrgCode Code `sdif:"start=3,len=1,code=001"`
	  SwimmerName string `sdif:"start=16,len=28"`
	  USS string `sdif:"start=44,len=12"`
	  sequence number to order multiple splits INT `sdif:"start=56,len=1"`
				 records for one athlete and one event
	  total number of splits for this event, INT `sdif:"start=57,len=2"`
	  split distance INT `sdif:"start=59,len=4"`
	  SPLIT Code 015, table checked CODE `sdif:"start=63,len=1"`
	  split time TIME `sdif:"start=64,len=8"`
	  split time TIME `sdif:"start=72,len=8"`
	  split time TIME `sdif:"start=80,len=8"`
	  split time TIME `sdif:"start=88,len=8"`
	  split time TIME `sdif:"start=96,len=8"`
	  split time TIME `sdif:"start=104,len=8"`
	  split time TIME `sdif:"start=112,len=8"`
	  split time TIME `sdif:"start=120,len=8"`
	  split time TIME `sdif:"start=128,len=8"`
	  split time TIME `sdif:"start=136,len=8"`
	  PRELIMS/FINALS Code 019, table checked CODE `sdif:"start=144,len=1"`

}

func (r SplitsRecord) ConstTag() string { return "G0" }

type FileTerminatorRecord struct{}

func (r FileTerminatorRecord) ConstTag() string { return "Z0" }

type FileDescriptionRecord struct {
	OrgCode      Code      `sdif:"start=3,len=1,code=001"`
	VersionNum   string    `sdif:"start=4,len=8"`
	FileCode     Code      `sdif:"start=12,len=2"`
	SoftwareName string    `sdif:"start=44,len=20"`
	SoftwareVer  string    `sdif:"start=64,len=10"`
	ContactName  string    `sdif:"start=74,len=20"`
	PhoneNumber  string    `sdif:"start=94,len=12"`
	LastModified time.Time `sdif:"start=106,len=8"`
	LscSubmitted string    `sdif:"start=156,len=2"`
}

func (r FileDescriptionRecord) ConstTag() string {
	return "A0"
}

type FileTerminationRecord struct {
}

func (r FileTerminationRecord) ConstTag() string {
	return "Z0"
}

type MeetRecord struct {
	OrgCode     Code      `sdif:"start=3,len=1,code=001"`
	MeetName    string    `sdif:"start=12,len=30"`
	AddrLineOne string    `sdif:"start=42,len=22"`
	AddrLineTwo string    `sdif:"start=64,len=22"`
	City        string    `sdif:"start=86,len=20"`
	State       USPS      `sdif:"start=106,len=2"`
	PostCode    string    `sdif:"start=108,len=10"`
	Country     Code      `sdif:"start=118,len=3,code=004"`
	MeetCode    Code      `sdif:"start=121,len=1,code=005"`
	Start       time.Time `sdif:"start=122,len=8"`
	End         time.Time `sdif:"start=130,len=8"`
	Altitude    int       `sdif:"start=138,len=4"`
	Course      Code      `sdif:"start=150,len=1,code=013"`
}

func (r MeetRecord) ConstTag() string {
	return "B1"
}
