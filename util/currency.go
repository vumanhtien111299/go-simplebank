package util

// Constants for supported currencies
const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
)

// IsSupportedCurrency returns true if the currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD:
		return true
	case EUR:
		return true
	case CAD:
		return true
	}
	return false
}
