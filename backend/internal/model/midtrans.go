package model

type MidtransNotificationRequest struct {
	TransactionTime   string     `json:"transaction_time"`
	TransactionStatus string     `json:"transaction_status"`
	TransactionID     string     `json:"transaction_id"`
	StatusMessage     string     `json:"status_message"`
	StatusCode        string     `json:"status_code"`
	SignatureKey      string     `json:"signature_key"`
	PaymentType       string     `json:"payment_type"`
	OrderID           string     `json:"order_id"`
	MerchantID        string     `json:"merchant_id"`
	GrossAmount       string     `json:"gross_amount"`
	FraudStatus       string     `json:"fraud_status"`
	Currency          string     `json:"currency"`
	VANumbers         []VANumber `json:"va_numbers,omitempty"`
}

type VANumber struct {
	VANumber string `json:"va_number"`
	Bank     string `json:"bank"`
}
