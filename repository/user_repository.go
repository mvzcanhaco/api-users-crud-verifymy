package repository

import (
	"github.com/mvzcanhaco/api-users-crud-verifymy/domain/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *entity.User) error
	FindByID(id uint64) (*entity.User, error)
	FindAll(page, pageSize int) ([]*entity.User, error)
	Update(user *entity.User) error
	Delete(id uint64) error
	FindByEmail(email string) (*entity.User, error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (r *UserRepositoryImpl) Create(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepositoryImpl) FindByID(id uint64) (*entity.User, error) {
	var user entity.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) FindAll(page, pageSize int) ([]*entity.User, error) {
	var users []*entity.User
	offset := (page - 1) * pageSize

	result := r.db.Limit(pageSize).Offset(offset).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (r *UserRepositoryImpl) Update(user *entity.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepositoryImpl) Delete(id uint64) error {
	return r.db.Delete(&entity.User{}, id).Error
}

func (r *UserRepositoryImpl) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
