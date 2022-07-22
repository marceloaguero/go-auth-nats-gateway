package user

// User describes a user in the system
type User struct {
	ID       uint   `gorm:"primaryKey"`
	Surname  string `gorm:"size:60" validate:"required,gte=2,lte=60"`
	Name     string `gorm:"size:60" validate:"lte=60"`
	Email    string `gorm:"size=60,unique" validate:"email,required"`
	Password string `gorm:"-"`
	Hash     string `gorm:"size=60"`
	IsActive bool
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
