package models

type ApiResponse[T any] struct {
	Status  int    `bson:"status" json:"status,omitempty"`
	Message string `bson:"message" json:"message,omitempty"`
	Data    T      `bson:"data" json:"data,omitempty"`
}
