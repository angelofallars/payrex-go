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
	Card PaymentMethodOptionsCard `json:"card" query:"card"`
}

type PaymentMethodOptionsCard struct {
	CaptureType    CaptureType       `json:"capture_type" query:"capture_type"`
	AllowedBins    *[]string         `json:"allowed_bins" query:"allowed_bins"`
	AllowedFunding *[]AllowedFunding `json:"allowed_funding" query:"allowed_funding"`
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
