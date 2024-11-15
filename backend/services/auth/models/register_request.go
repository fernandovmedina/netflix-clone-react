package models

type RegisterRequest struct {
	Email          string      `bson:"email" json:"email"`
	HashedPassword string      `bson:"hashed_password" json:"hashed_password"`
	Plan           PlanDetails `bson:"plan_details" json:"plan_details"`
	Users          []User      `bson:"users" json:"users"`
	Payments       []Payment   `bson:"payment" json:"payment"`
	InitialDate    string      `bson:"initial_date" json:"initial_date"`
	ExpiresDate    string      `bson:"expires_date" json:"expires_date"`
	Amount         float32     `bson:"amount" json:"amount"`
	Status         bool        `bson:"status" json:"status"`
	AutoRenew      bool        `bson:"auto_renew" json:"auto_renew"`
}

type PlanDetails struct {
	Type        string `bson:"type" json:"type"`
	DeviceLimit uint8  `bson:"device_limit" json:"device_limit"`
	Resolution  string `bson:"resolution" json:"resolution"`
}

type User struct {
	Name     string `bson:"name" json:"name"`
	UserType string `bson:"user_type" json:"user_type"`
	Icon     string `bson:"icon" json:"icon"`
}

type Payment struct {
	Card *CardPaymentRequest     `bson:"card,omitempty" json:"card,omitempty"`
	Oxxo *OxxoPaymentRequest     `bson:"oxxo,omitempty" json:"oxxo,omitempty"`
	Gift *GiftCodePaymentRequest `bson:"gift,omitempty" json:"gift,omitempty"`
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
