package staff

import "github.com/google/uuid"

type Staff struct {
	StaffId       uuid.UUID
	StaffName     string
	Job           string
	Age           int
	Fee           int
	Salary        int
	Training      int
	Finances      int
	Scouting      int
	Physiotherapy int
	Rarity        int
}
