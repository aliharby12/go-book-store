package models

import "time"

type Book struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time `json:"created_at"`
	Name         string    `json:"name"`
	SerialNumber string    `json:"serial_number"`
	Price        float64   `json:"price"`
}
