package models

import (
	"time"
)

type GetProducts struct {
	Product string    `bson:"product" json:"product,omitempty"`
	Date    time.Time `bson:"date" json:"date,omitempty"`
}
