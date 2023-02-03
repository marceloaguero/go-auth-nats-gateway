package user

// User describes a user in the system
type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Surname  string `json:"surname" gorm:"size:60" validate:"required,gte=2,lte=60"`
	Name     string `json:"name,omitempty" gorm:"size:60" validate:"lte=60"`
	Email    string `json:"email" gorm:"size=60,unique" validate:"email,required"`
	Password string `json:"password" gorm:"-"`
	Hash     string `json:"-" gorm:"size=60"`
	IsActive bool   `json:"is_active"`
}

// Repository represents the user permanent repo
type Repository interface {
	Create(user *User) (*User, error)
	GetByID(id uint) (*User, error)
	GetByEmail(email string) (*User, error)
	GetAll() ([]*User, error)
	Update(user *User) (*User, error)
	Delete(user *User) error
}
