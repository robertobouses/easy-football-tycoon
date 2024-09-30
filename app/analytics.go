package app

import "github.com/google/uuid"

type Analytics struct {
	AnalyticsId             uuid.UUID
	Training                int
	Finances                int
	Scouting                int
	Physiotherapy           int
	TotalSalaries           int
	TotalTraining           int
	TotalFinances           int
	TotalScouting           int
	TotalPhysiotherapy      int
	TrainingStaffCount      int
	FinanceStaffCount       int
	ScoutingStaffCount      int
	PhysiotherapyStaffCount int
	Trust                   int
	StadiumCapacity         int
}
