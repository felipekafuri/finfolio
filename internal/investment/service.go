package investment

import (
	"errors"
	"strconv"
	"time"
)

// ParseFormData converts raw form strings into an investment
func ParseFormData(data map[string]string) (*Investment, error) {
	appDate, err := time.Parse("2006/01/02", data["Application Date"])
	if err != nil {
		return nil, errors.New("invalid application date format (use YYYY/MM/DD)")
	}

	value, err := strconv.ParseFloat(data["Value"], 64)
	if err != nil {
		return nil, errors.New("invalid value format")
	}

	redDate, err := time.Parse("2006/01/02", data["Redemption Date"])
	if err != nil {
		return nil, errors.New("invalid redemption date format (use YYYY/MM/DD)")
	}
	if redDate.Before(appDate) {
		return nil, errors.New("redemption date cannot be before application date")
	}

	// Calculate period in days
	periodDays := int(redDate.Sub(appDate).Hours() / 24)

	investment := &Investment{
		ApplicationDate: appDate,
		Value:           value,
		RedemptionDate:  redDate,
		Bank:            data["Bank"],
		Title:           data["Title"],
		PeriodDays:      periodDays,
		// TODO: Add this part
		// GrossReturn, Tax, etc. will be added later when user updates the investment
	}

	return investment, nil
}

// CalculateReturns updates the calculated fields of an investment
func CalculateReturns(inv *Investment) {
	// Net return = Gross - Tax
	inv.NetReturn = inv.GrossReturn - inv.Tax

	// Return percentage for the period
	if inv.Value > 0 {
		inv.ReturnPercent = inv.NetReturn / inv.Value
	}

	// Annualized return: ((1 + ReturnPercent)^(365/period)) - 1
	if inv.PeriodDays > 0 {
		inv.AnnualizedRate = (pow(1+inv.ReturnPercent, 365.0/float64(inv.PeriodDays))) - 1
	}
}

// Simple power function
func pow(base, exp float64) float64 {
	result := 1.0
	for i := 0; i < int(exp); i++ {
		result *= base
	}
	return result
}
