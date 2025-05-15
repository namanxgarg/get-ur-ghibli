package repository

import (
    "github.com/jinzhu/gorm"
    "github.com/example/get-ur-ghibli/auth-service/internal/models"
)

type UserRepository interface {
    CreateUser(user *models.User) error
    FindByEmail(email string) (*models.User, error)
    UpdateUser(user *models.User) error
}

type userRepo struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepo{db: db}
}

func (r *userRepo) CreateUser(user *models.User) error {
    return r.db.Create(user).Error
}

func (r *userRepo) FindByEmail(email string) (*models.User, error) {
    var user models.User
    err := r.db.Where("email = ?", email).First(&user).Error
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *userRepo) UpdateUser(user *models.User) error {
    return r.db.Save(user).Error
}
