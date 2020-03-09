package model

type User struct {
	Model
	Email    string `gorm:"unique_index;not null" json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
