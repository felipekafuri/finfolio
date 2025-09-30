package investment

import "time"

type Investment struct {
	ID              int
	ApplicationDate time.Time
	Value           float64
	RedemptionDate  time.Time
	Bank            string
	Title           string
	GrossReturn     float64
	Tax             float64
	NetReturn       float64
	PeriodDays      int
	ReturnPercent   float64
	AnnualizedRate  float64
}
