package utils

import "time"

// Função para calcular a idade com base na data de nascimento
func CalculateAge(BirthDate string) (int, error) {
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

	return age, err
}
