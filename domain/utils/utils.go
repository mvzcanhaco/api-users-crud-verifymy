package utils

import (
	"errors"
	"fmt"
	"regexp"
	"time"
)

// Função para calcular a idade com base na data de nascimento
func CalculateAge(BirthDate string) (int, error) {
	// Verifica se a data de nascimento foi informada
	if BirthDate == "" {
		return 0, errors.New("A data de nascimento deve ser informada")
	}

	// Calcula a idade com base na data de nascimento
	birthDate, err := time.Parse("2006-01-02", BirthDate)
	if err != nil {
		// Tratar erro na conversão da data de nascimento
		return 0, err
	}

	now := time.Now()
	age := now.Year() - birthDate.Year()

	// Ajusta a idade se ainda não fez aniversário neste ano
	if now.Month() < birthDate.Month() || (now.Month() == birthDate.Month() && now.Day() < birthDate.Day()) {
		age--
	}

	if age < 0 {
		return 0, errors.New("A data informada é do futuro")
	}

	return age, nil
}

func IsValidEmail(email string) (bool, error) {
	// Define a regular expression to match valid email addresses
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	// Check if the email address matches the regular expression
	if !re.MatchString(email) {
		return false, fmt.Errorf("invalid email address")
	}

	return true, nil
}
