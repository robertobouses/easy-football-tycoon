package staff

import "github.com/google/uuid"

type Staff struct {
	StaffId       uuid.UUID
	FirstName     string
	LastName      string
	Nationality   string
	Job           string
	Age           int
	Fee           int
	Salary        int
	Training      int
	Finances      int
	Scouting      int
	Physiotherapy int
	Knowledge     int
	Intelligence  int
	Rarity        int
}
