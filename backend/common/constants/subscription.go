package constants

// PaymentStatus represents the various statuses a payment can have.
type PaymentStatus string

const (
	PaymentStatusCreated           PaymentStatus = "CREATED"
	PaymentStatusAuthorized        PaymentStatus = "AUTHORIZED"
	PaymentStatusCaptured          PaymentStatus = "CAPTURED"
	PaymentStatusFailed            PaymentStatus = "FAILED"
	PaymentStatusRefunded          PaymentStatus = "REFUNDED"
	PaymentStatusPartiallyRefunded PaymentStatus = "PARTIALLY_REFUNDED"
	PaymentStatusPending           PaymentStatus = "PENDING"
	PaymentStatusProcessing        PaymentStatus = "PROCESSING"
	PaymentStatusCancelled         PaymentStatus = "CANCELLED"
	PaymentStatusDisputed          PaymentStatus = "DISPUTED"
)
