package payrex

// PaymentMethod enumerates the valid payment methods for PayRex payments.
type PaymentMethod string

const (
	// Pay through bank card.
	PaymentMethodCard PaymentMethod = "card"
	// Pay through GCash.
	PaymentMethodGCash PaymentMethod = "gcash"
	// Pay through Maya.
	PaymentMethodMaya PaymentMethod = "maya"
	// Pay through QR Ph.
	PaymentMethodQRPh PaymentMethod = "qrph"
)

type PaymentMethodOptions struct {
	Card Card `json:"card" form:"card"`
}

type Card struct {
	CaptureType    CaptureType       `json:"capture_type" form:"capture_type"`
	AllowedBins    *[]string         `json:"allowed_bins" form:"allowed_bins"`
	AllowedFunding *[]AllowedFunding `json:"allowed_funding" form:"allowed_funding"`
}

type CaptureType string

const (
	CaptureTypeAutomatic CaptureType = "automatic"
	CaptureTypeManual    CaptureType = "manual"
)

type AllowedFunding string

const (
	AllowedFundingCredit AllowedFunding = "credit"
	AllowedFundingDebit  AllowedFunding = "debit"
)
