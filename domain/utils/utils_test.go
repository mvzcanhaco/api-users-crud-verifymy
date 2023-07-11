package utils

import (
	"testing"
)

func TestCalculateAge(t *testing.T) {
	// Test case 1: Valid date
	age, err := CalculateAge("1990-01-01")
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if age != 33 {
		t.Errorf("Expected age to be 29, but got %d", age)
	}

	// Test case 2: Invalid date format
	_, err = CalculateAge("01/02/2006")
	if err == nil {
		t.Error("Expected error, but got none")
	}

	// Test case 3: Date in the future
	_, err = CalculateAge("2050-01-01")
	if err == nil {
		t.Error("Expected error, but got none")
	}

	// Test case 4: Empty date
	_, err = CalculateAge("")
	if err == nil {
		t.Error("Expected error, but got none")
	}
}

func TestIsValidEmail(t *testing.T) {
	// Test valid email addresses
	validEmails := []string{"example@example.com", "example@w.com", "example@example.br"}
	for _, email := range validEmails {
		isValid, err := IsValidEmail(email)
		if err != nil {
			t.Errorf("Expected no error for %s, but got %v", email, err)
		}
		if !isValid {
			t.Errorf("Expected %s to be a valid email address", email)
		}
	}

	// Test invalid email addresses
	invalidEmails := []string{"example", "example@", "@example.com", "example@example.", "example@.com"}
	for _, email := range invalidEmails {
		isValid, err := IsValidEmail(email)
		if err != nil {
			t.Logf("Expected no error for %s, but got %v", email, err)
		}
		if isValid {
			t.Errorf("Expected %s to be an invalid email address", email)
		}
	}
}
