package models

/*
	-- date:   October 21, 23:26 (UTC-5)
	-- author: Fernando Vazquez
*/

type Movie struct {
	Id            int        `bson:"id" json:"id"`
	Name          string     `bson:"name" json:"name"`
	CreatedAt     string     `bson:"created_at" json:"created_at"`
	UpdatedAt     string     `bson:"updated_at" json:"updated_at"`
	DeletedAt     string     `bson:"deleted_at" json:"deleted_at"`
	BackgroundURL string     `bson:"background_url" json:"background_url"`
	Duration      string     `bson:"duration" json:"duration"`
	YearReleased  uint16     `bson:"year_released" json:"year_released"`
	Description   string     `bson:"description" json:"description"`
	VideoURL      string     `bson:"video_url" json:"video_url"`
	Actors        Actors     `bson:"actors" json:"actors"`
	Directors     Directors  `bson:"directors" json:"directors"`
	Categories    Categories `bson:"categories" json:"categories"`
}

type Movies []Movie
