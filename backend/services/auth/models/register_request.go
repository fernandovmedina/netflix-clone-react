package models

type RegisterRequest struct {
	Email          string      `bson:"email" json:"email"`
	HashedPassword string      `bson:"hashed_password" json:"hashed_password"`
	Plan           PlanDetails `bson:"plan" json:"plan"`
	Users          []User      `bson:"users" json:"users"`
	Payment        Payment     `bson:"payment" json:"payment"`
	InitialDate    string      `bson:"initial_date" json:"initial_date"`
	ExpiresDate    string      `bson:"expires_date" json:"expires_date"`
	Amount         float32     `bson:"amount" json:"amount"`
	Status         bool        `bson:"status" json:"status"`
	AutoRenew      bool        `bson:"auto_renew" json:"auto_renew"`
}

type PlanDetailsRequest struct {
	Type        string `bson:"type" json:"type"`
	DeviceLimit uint8  `bson:"device_limit" json:"device_limit"`
	Resolution  string `bson:"resolution" json:"resolution"`
}

type UserRequest struct {
	Name     string `bson:"name" json:"name"`
	UserType string `bson:"user_type" json:"user_type"`
	Icon     string `bson:"icon" json:"icon"`
}

type PaymentRequest struct {
	CARD CardPaymentRequest
	OXXO OxxoPaymentRequest
	GIFT GiftCodePaymentRequest
}

type CardPaymentRequest struct {
	Number  string `bson:"card_number" json:"card_number"`
	DueDate string `bson:"due_date" json:"due_date"`
	CVV     uint16 `bson:"cvv" json:"cvv"`
	Name    string `bson:"card_name" json:"card_name"`
}

type OxxoPaymentRequest struct {
	Name string `bson:"oxxo_name" json:"oxxo_name"`
	Code string `bson:"oxxo_code" json:"oxxo_code"`
}

type GiftCodePaymentRequest struct {
	Code string `bson:"gift_code" json:"gift_code"`
}
