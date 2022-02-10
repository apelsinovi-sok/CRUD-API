package entity

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID
	Firstname string `json:"firstname"`
	Lastname  string `form:"lastname" json:"lastname" binding:"required"`
	Email     string `json:"email"`
	Age       uint   `json:"age"`
	Created   time.Time
}
