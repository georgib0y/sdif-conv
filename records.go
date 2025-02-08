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

type MeetHostRecord struct{}

func (r MeetHostRecord) ConstTag() string { return "B2" }

type TeamIdRecord struct{}

func (r TeamIdRecord) ConstTag() string { return "C1" }

type TeamEntryRecord struct{}

func (r TeamEntryRecord) ConstTag() string { return "C2" }

type IndividualEventRecord struct{}

func (r IndividualEventRecord) ConstTag() string { return "D0" }

type IndividualContactRecord struct{}

func (r IndividualContactRecord) ConstTag() string { return "D2" }

type IndividualInfoRecord struct{}

func (r IndividualInfoRecord) ConstTag() string { return "D3" }

type RelayEventRecord struct{}

func (r RelayEventRecord) ConstTag() string { return "E0" }

type RelayNameRecord struct{}

func (r RelayNameRecord) ConstTag() string { return "F0" }

type SplitsRecord struct{}

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

type MeetPyramid struct {
	fileDesc  FileDescriptionRecord
	meetRec   MeetRecord
	meetHosts []MeetHostRecord

	fileTerm FileTerminatorRecord
}
