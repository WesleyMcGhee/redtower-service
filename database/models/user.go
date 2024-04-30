package models

// Eventually we will want to do timestampping on this, and also probably expand it some more
type User struct {
	ID           uint   `json:"id" gorm:"primary_key"`
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash"`
	Email        string `json:"email"`
}