package model

type PaymentStatus string

const (
	PaymentPending   PaymentStatus = "PENDING"
	PaymentPaid      PaymentStatus = "PAID"
	PaymentExpired   PaymentStatus = "EXPIRED"
	PaymentCancelled PaymentStatus = "CANCELLED"
	PaymentFailed    PaymentStatus = "FAILED"
	PaymentRefunded  PaymentStatus = "REFUNDED"
)

type Payment struct {
	OrderID uint `json:"order_id,omitempty"`

	Method string `json:"method,omitempty"`

	Token       string `json:"token,omitempty"`
	RedirectURL string `json:"redirect_url,omitempty"`

	TransactionID string        `json:"transaction_id,omitempty"`
	Status        PaymentStatus `json:"status,omitempty"`
}
