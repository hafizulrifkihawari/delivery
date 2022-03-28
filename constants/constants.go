package constants

const (
	// time format reference
	MilitaryTime = "15:04:05"
	Kitchen      = "3:04 pm"
	Kitchen2     = "3 pm"
	DateTime     = "01/02/2006 15:04 PM"
)

const (
	// regex format
	Alphanumeric      = "[^a-zA-Z0-9 ]+"
	AlphanumericSpace = "[^a-zA-Z0-9 ]+"
)

var (
	DayMapping = map[int]string{
		0: "Sun",
		1: "Mon",
		2: "Tues",
		3: "Weds",
		4: "Thurs",
		5: "Fri",
		6: "Sat",
	}
)
