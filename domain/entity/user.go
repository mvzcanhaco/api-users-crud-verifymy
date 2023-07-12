package entity

import (
	"errors"
	"time"

	"github.com/mvzcanhaco/api-users-crud-verifymy/domain/utils"
)

type User struct {
	ID        uint64   `gorm:"primaryKey" json:"id,omitempty"`
	Name      string   `gorm:"not null" json:"name,omitempty" validate:"nonzero"`
	Email     string   `gorm:"not null;unique" json:"email,omitempty"`
	Password  string   `gorm:"not null" json:"password,omitempty"`
	BirthDate string   `gorm:"not null" json:"birthDate,omitempty"`
	Age       int      `json:"age,omitempty"`
	Profile   string   `gorm:"not null" json:"Profile,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

func (u *User) Validate() error {
	// Validate the Name field
	if u.Name == "" {
		return errors.New("Name cannot be empty")
	}

	// Validate the Email field
	isValid, err := utils.IsValidEmail(u.Email)
	if !isValid {
		return errors.New("Email must be a valid email address")
	}

	// Validate the Password field
	if len(u.Password) < 6 {
		return errors.New("Password must be at least 6 characters long")
	}

	// Validate the BirthDate field
	_, err = time.Parse("2006-01-02", u.BirthDate)
	if err != nil {
		return errors.New("BirthDate must be a valid date in the format YYYY-MM-DD")
	}

	// Validate the Profile field
	if u.Profile != "admin" && u.Profile != "user" {
		return errors.New("Profile must be either 'admin' or 'user'")
	}

	// Validate the Address field
	if u.Address != nil {
		if u.Address.Street == "" {
			return errors.New("Street cannot be empty")
		}
		if u.Address.City == "" {
			return errors.New("City cannot be empty")
		}
		if u.Address.State == "" {
			return errors.New("State cannot be empty")
		}
		if u.Address.Country == "" {
			return errors.New("PostalCode cannot be empty")
		}
	}

	// If all validations pass, return nil
	return nil
}
