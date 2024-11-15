package models

/*
type Account struct {
	Email          string      `bson:"email" json:"email"`
	HashedPassword string      `bson:"hashed_password" json:"hashed_password"`
	Plan           PlanDetails `bson:"plan_details" json:"plan_details"`
	Users          []User      `bson:"users" json:"users"`
	Payment        Payment     `bson:"payment" json:"payment"`
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

type History struct {
	MoviesHistory []MovieHistory `bson:"movies_history" json:"movies_history"`
	SeriesHistory []SerieHistory `bson:"series_history" json:"series_history"`
}

type MovieHistory struct {
	Id                uint   `bson:"id_movie" json:"id_movie"`
	LastMinuteWatched string `bson:"last_minute_watched" json:"last_minute_watched"`
}

type SerieHistory struct {
	Id                 uint   `bson:"id_serie" json:"id_serie"`
	LastSeasonWatched  uint8  `bson:"last_season_watched" json:"last_season_watched"`
	LastChapterWatched uint8  `bson:"last_chapter_watched" json:"last_chapter_watched"`
	LastMinuteWatched  string `bson:"last_minute_watched" json:"last_minute_watched"`
}

type User struct {
	Name     string  `bson:"name" json:"name"`
	UserType string  `bson:"user_type" json:"user_type"`
	Icon     string  `bson:"icon" json:"icon"`
	History  History `bson:"history" json:"history"`
}

type Payment struct {
	CARD CardPayment     `bson:"card" json:"card"`
	OXXO OxxoPayment     `bson:"oxxo" json:"oxxo"`
	GIFT GiftCodePayment `bson:"gift" json:"gift"`
}

type CardPayment struct {
	Number  string `bson:"card_number" json:"card_number"`
	DueDate string `bson:"due_date" json:"due_date"`
	CVV     uint16 `bson:"cvv" json:"cvv"`
	Name    string `bson:"card_name" json:"card_name"`
}

type OxxoPayment struct {
	Name string `bson:"oxxo_name" json:"oxxo_name"`
	Code string `bson:"oxxo_code" json:"oxxo_code"`
}

type GiftCodePayment struct {
	Code string `bson:"gift_code" json:"gift_code"`
}
*/
