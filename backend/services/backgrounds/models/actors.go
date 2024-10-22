package models

/*
	-- date:   October 21, 23:29 (UTC-5)
	-- author: Fernando Vazquez
*/

type Actor struct {
	Id   int    `bson:"id" json:"id"`
	Name string `bson:"name" json:"name"`
}

type Actors []Actor
