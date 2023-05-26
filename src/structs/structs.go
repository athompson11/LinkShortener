package structs

import "time"

type Link struct {
	LinkID    string    `db:"linkid" json:"linkid"`
	Link      string    `db:"link" json:"link"`
	CreatedAt time.Time `db:"creationtime" json:"createdat"`
}
