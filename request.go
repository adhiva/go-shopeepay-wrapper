package shopeepay

type RequestCharge struct {
	RequestID                string `json:"request_id"`
	PaymentReferenceID       string `json:"payment_reference_id"`
	MerchantExtID            string `json:"merchant_ext_id"`
	StoreExtID               string `json:"store_ext_id"`
	Amount                   int64  `json:"amount"`
	Currency                 string `json:"currency"`
	ReturnURL                string `json:"return_url,omitempty"`
	PointOfInitiation        string `json:"point_of_initiation,omitempty"`
	ValidityPeriod           uint32 `json:"validity_period,omitempty"`
	AdditionalInfo           string `json:"additional_info,omitempty"`
	TerminalID               string `json:"terminal_id,omitempty"`
	QRValidityPeriod         uint32 `json:"qr_validity_period,omitempty"`
	ConvenienceFeeIndicator  string `json:"convenience_fee_indicator,omitempty"`
	ConvenienceFeeFixed      string `json:"convenience_fee_fixed,omitempty"`
	ConvenienceFeePercentage string `json:"convenience_fee_percentage,omitempty"`
}

// This Request can use for Invalidate QR Code
// Or Checking Payment Status
type RequestGeneral struct {
	RequestID          string `json:"request_id"`
	MerchantExtID      string `json:"merchant_ext_id"`
	StoreExtID         string `json:"store_ext_id"`
	PaymentReferenceID string `json:"payment_reference_id"`
}

type RequestLog struct {
	Method       string `json:"method"`
	Host         string `json:"host"`
	Path         string `json:"path"`
	RequestBody  string `json:"request_body"`
	ResponseBody string `json:"response_body"`
	TimeLength   int    `json:"time_length"`
}

type RequestNotification struct {
	PaymentReferenceID string `json:"payment_reference_id"`
	Amount             int64  `json:"amount"`
	TransactionSN      string `json:"transaction_sn"`
	PaymentStatus      uint32 `json:"payment_status"`
	PaymentMethod      uint32 `json:"payment_method"`
	UserIDHash         string `json:"user_id_hash"`
	MerchantExtID      string `json:"merchant_ext_id"`
	StoreExtID         string `json:"store_ext_id"`
	TerminalID         string `json:"terminal_id"`
}
