package models

/*
	-- date:   October 21, 23:27 (UTC-5)
	-- author: Fernando Vazquez
*/

type Serie struct {
	Id              int    `bson:"id" json:"id"`
	Name            string `bson:"name" json:"name"`
	BackgroundURL   string `bson:"background_url" json:"background_url"`
	CreatedAt       string `bson:"created_at" json:"created_at"`
	UpdatedAt       string `bson:"created_at" json:"updated_at"`
	DeletedAt       string `bson:"deleted_at" json:"deleted_at"`
	YearReleased    uint16 `bson:"year_released" json:"year_released"`
	NumberOfSeasons uint8  `bson:"number_of_seasons" json:"number_of_seasons"`
	Description     string `bson:"description" json:"description"`
}

type Series []Serie
