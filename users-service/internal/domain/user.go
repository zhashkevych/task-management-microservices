package domain

type (
	User struct {
		Id        int    `json:"id" db:"id"`
		FirstName string `json:"first_name" db:"first_name" binding:"required,min=2,max=20"`
		LastName  string `json:"last_name" db:"last_name" binding:"required,min=2,max=20"`
		Username  string `json:"username" db:"username" binding:"required,min=5,max=20"`
		Password  string `json:"password,omitempty" binding:"required,min=8,max=30"`
	}
)
