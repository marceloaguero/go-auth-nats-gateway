package mysql_orm

import (
	"fmt"

	"github.com/BetuelSA/go-helpers/errors"
	"github.com/marceloaguero/go-auth-nats-gateway/users/pkg/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ormRepo struct {
	db *gorm.DB
}

// NewRepo creates a repository implemented in ORM (MySQL)
func NewRepo(dsName, dbName string) (user.Repository, error) {
	db, err := dbConnect(dsName, dbName)
	if err != nil {
		return nil, errors.Wrap(err, "can't connect to DB")
	}

	db.AutoMigrate(&user.User{})

	return &ormRepo{
		db: db,
	}, nil
}

func dbConnect(dsName, dbName string) (*gorm.DB, error) {
	conn := fmt.Sprintf("%s/%s?charset=utf8&parseTime=True&loc=Local", dsName, dbName)

	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (r *ormRepo) Create(user *user.User) (*user.User, error) {
	result := r.db.Create(&user)
	return user, result.Error
}

func (r *ormRepo) GetByID(id uint) (*user.User, error) {
	var user user.User
	result := r.db.Take(&user, id)
	return &user, result.Error
}

func (r *ormRepo) GetByEmail(email string) (*user.User, error) {
	var user user.User
	result := r.db.Take(&user, "email = ?", email)
	return &user, result.Error
}

func (r *ormRepo) GetAll() ([]*user.User, error) {
	users := []*user.User{}
	result := r.db.Find(&users)
	return users, result.Error
}

func (r *ormRepo) Update(user *user.User) (*user.User, error) {
	result := r.db.Model(&user).Updates(user)
	if result.Error == nil {
		r.db.Model(&user).Updates(map[string]interface{}{"is_active": user.IsActive})
	}
	return user, result.Error
}

func (r *ormRepo) Delete(user *user.User) error {
	result := r.db.Delete(&user)
	return result.Error
}
