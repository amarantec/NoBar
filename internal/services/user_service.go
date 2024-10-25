package services

import (
    "context"
    "errors"
    "github.com/amarantec/nobar/internal/models"
    "github.com/amarantec/nobar/internal/utils"
    "gorm.io/gorm"
    "fmt"
)

func (s *ServicePostgres) Register (ctx context.Context, user models.Users) (bool, error) {
    hashedPassword, err := utils.HashPassword(user.Password)
    if err != nil {
        return false, err
    }

    user.Password = hashedPassword

    if err :=
        s.Db.WithContext(ctx).
        Create(&user).Error; err != nil {
            return false, err
        }

    return true, nil
}

func (s *ServicePostgres) Login (ctx context.Context, user models.Users) (uint, error) {
    userScan := models.Users{}
    if err :=
        s.Db.WithContext(ctx).
        Model(&models.Users{}).
        Select("id, password").
        Where("name = ?", user.Name).
        Scan(&userScan).Error; err != nil {
            if errors.Is(err, gorm.ErrRecordNotFound) {
                return 0, nil
            }
            return 0, err
        }

    passwordIsValid :=
        utils.CheckPassword(user.Password, userScan.Password)
    if !passwordIsValid {
        return 0, fmt.Errorf("wrong password")
    }

    return userScan.ID, nil
}

