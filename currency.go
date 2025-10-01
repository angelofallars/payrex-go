package payrex

// Currency is a three-letter ISO currency code in uppercase.
// As of the moment, PayRex only supports PHP.
type Currency string

const (
	CurrencyPHP Currency = "PHP"
)
