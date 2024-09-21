package constants

// PaymentStatus represents the various statuses a payment can have.
type PaymentStatus string

const (
	PaymentStatusPending      PaymentStatus = "PENDING"
	PaymentStatusCancelled    PaymentStatus = "CANCELLED"
	PaymentStatusFailed       PaymentStatus = "FAILED"
	PaymentStatusSuccess      PaymentStatus = "SUCCESS"
	PaymentStatusNotAttempted PaymentStatus = "NOT_ATTEMPTED"
	PaymentStatusUserDropped  PaymentStatus = "USER_DROPPED"
	PaymentStatusVoid         PaymentStatus = "VOID"
)
