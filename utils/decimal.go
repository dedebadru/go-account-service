package utils

import (
    "github.com/shopspring/decimal"
)

func FormatDecimal(value decimal.Decimal, precision int32) decimal.Decimal {
    return value.Round(precision)
}

func AddDecimal(a, b decimal.Decimal) decimal.Decimal {
    return a.Add(b)
}

func SubtractDecimal(a, b decimal.Decimal) decimal.Decimal {
    return a.Sub(b)
}
