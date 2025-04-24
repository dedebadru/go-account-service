package utils

import (
	"github.com/shopspring/decimal"
)

func FormatDecimal(value decimal.Decimal, precision int32) decimal.Decimal {
	return value.Round(precision)
}

func AddDecimal(value, enhancer decimal.Decimal) decimal.Decimal {
	return value.Add(enhancer)
}

func SubtractDecimal(value, reducer decimal.Decimal) decimal.Decimal {
	return value.Sub(reducer)
}
