// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package db

import (
	"database/sql"
	"time"
)

type Country struct {
	Code          int32  `json:"code"`
	CountryName   string `json:"countryName"`
	ContinentName string `json:"continentName"`
}

type Director struct {
	ID        int32  `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type Movie struct {
	ID                    int32         `json:"id"`
	Title                 string        `json:"title"`
	DirectorID            sql.NullInt32 `json:"directorID"`
	AdminID               sql.NullInt32 `json:"adminID"`
	ReleaseYear           time.Time     `json:"releaseYear"`
	ProductionCountryCode int32         `json:"productionCountryCode"`
	CreatedAt             time.Time     `json:"createdAt"`
}

type User struct {
	ID          int32         `json:"id"`
	FullName    string        `json:"fullName"`
	CreatedAt   time.Time     `json:"createdAt"`
	CountryCode sql.NullInt32 `json:"countryCode"`
}
