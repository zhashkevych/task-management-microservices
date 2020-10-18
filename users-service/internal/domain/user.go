package domain

type (
	User struct {
		FirstName string `json:"first_name" db:"first_name" binding:"required,min=2,max=20"`
		LastName  string `json:"last_name" db:"last_name" binding:"required,min=2,max=20"`
		Username  string `json:"username" db:"username" binding:"required,min=5,max=20"`
		Password  string `json:"password" db:"password" binding:"required,min=8,max=30"`
	}
)
