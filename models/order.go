package models

import "time"

type Order struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	BookID    int       `json:"book_id"`
	Book      Book      `json:"book"`
	UserID    int       `json:"user_id"`
	User      User      `json:"user"`
}
