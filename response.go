package shopeepay

type ResponseCharge struct {
	RequestID       string `json:"request_id"`
	ErrCode         int32  `json:"errcode"`
	DebugMsg        string `json:"debug_msg"`
	StoreName       string `json:"store_name"`
	QRContent       string `json:"qr_content,omitempty"`
	QRURL           string `json:"qr_url,omitempty"`
	RedirectURLApp  string `json:"redirect_url_app,omitempty"`
	RedirectURLHTTP string `json:"redirect_url_http,omitempty"`
}

type ResponseCancelQRCode struct {
	RequestID string `json:"request_id"`
	ErrCode   int32  `json:"errcode"`
	DebugMsg  string `json:"debug_message"`
}

type ResponseCheckPaymentStatus struct {
	RequestID       string          `json:"request_id"`
	ErrCode         int32           `json:"errcode"`
	DebugMsg        string          `json:"debug_message"`
	PaymentStatus   uint32          `json:"payment_status"`
	TransactionList TransactionList `json:"transaction_list,omitempty"`
}

type TransactionList struct {
	ReferenceID     string `json:"reference_id"`
	Amount          int64  `json:"amount"`
	TransactionSN   string `json:"transaction_sn"`
	Status          uint32 `json:"status"`
	TransactionType uint32 `json:"transaction_type"`
	CreateTime      uint32 `json:"create_time"`
	UpdateTime      uint32 `json:"update_time"`
	UserIDHash      string `json:"user_id_hash"`
	MerchantExtID   string `json:"merchant_ext_id"`
	StoreExtID      string `json:"store_ext_id"`
	TerminalID      string `json:"terminal_id"`
}
