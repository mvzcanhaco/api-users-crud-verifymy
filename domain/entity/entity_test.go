package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser_Validate(t *testing.T) {
	// Create a new User instance with valid data
	user := User{
		Name:      "John Doe",
		Email:     "johndoe@example.com",
		Password:  "password123",
		BirthDate: "2006-01-02",
		Profile:   "user",
		Address: &Address{
			Street:  "123 Main St",
			City:    "Anytown",
			State:   "CA",
			Country: "12345",
		},
	}

	// Call the Validate method on the user instance
	err := user.Validate()

	// Check that there are no errors
	if err != nil {
		t.Errorf("Expected no errors, but got %v", err)
	}
}

func TestUser_Validate_InvalidName(t *testing.T) {
	// Create a new User instance with an invalid name
	user := User{
		Name:      "",
		Email:     "johndoe@example.com",
		Password:  "password123",
		BirthDate: "2006-01-02",
		Profile:   "user",
		Address: &Address{
			Street:  "123 Main St",
			City:    "Anytown",
			State:   "CA",
			Country: "US",
		},
	}

	// Call the Validate method on the user instance
	err := user.Validate()

	// Check that there is an error for the invalid name
	if err == nil || err.Error() != "Name cannot be empty" {
		t.Errorf("Expected error for invalid name, but got %v", err)
	}
}

func TestUser_Validate_InvalidEmail(t *testing.T) {
	// Create a new User instance with an invalid email
	user := User{
		Name:      "John Doe",
		Email:     "invalidemail",
		Password:  "password123",
		BirthDate: "2006-01-02",
		Profile:   "user",
		Address: &Address{
			Street:  "123 Main St",
			City:    "Anytown",
			State:   "CA",
			Country: "US",
		},
	}

	// Call the Validate method on the user instance
	err := user.Validate()

	// Check that there is an error for the invalid email
	if err == nil || err.Error() != "Email must be a valid email address" {
		t.Errorf("Expected error for invalid email, but got %v", err)
	}
}

func TestUser_Validate_InvalidPassword(t *testing.T) {
	// Create a new User instance with an invalid password
	user := User{
		Name:      "John Doe",
		Email:     "johndoe@example.com",
		Password:  "1234",
		BirthDate: "2006-01-02",
		Profile:   "user",
		Address: &Address{
			Street:  "123 Main St",
			City:    "Anytown",
			State:   "CA",
			Country: "US",
		},
	}

	// Call the Validate method on the user instance
	err := user.Validate()

	// Check that there is an error for the invalid password
	if err == nil || err.Error() != "Password must be at least 6 characters long" {
		t.Errorf("Expected error for invalid password, but got %v", err)
	}
}

func TestAddress_Value(t *testing.T) {
	// Create a new Address instance with valid data
	address := Address{
		Street:  "123 Main St",
		City:    "Anytown",
		State:   "CA",
		Country: "USA",
	}

	// Call the Value method on the address instance
	value, err := address.Value()

	// Check that there are no errors
	assert.NoError(t, err)

	// Check that the value is a string
	assert.IsType(t, "", value)

	// Check that the value is the expected JSON string
	expectedValue := `{"street":"123 Main St","city":"Anytown","state":"CA","country":"USA"}`
	assert.Equal(t, expectedValue, value)
}
