package models

import "github.com/shopspring/decimal"

// ParseDecimal wraps decimal.NewFromString, returns zero on error.
func ParseDecimal(val string) (decimal.Decimal, error) {
	return decimal.NewFromString(val)
}

// MustParseDecimal returns decimal zero when parse fails (non-panicking).
func MustParseDecimal(val string) decimal.Decimal {
	d, err := decimal.NewFromString(val)
	if err != nil {
		return decimal.Zero
	}
	return d
}
