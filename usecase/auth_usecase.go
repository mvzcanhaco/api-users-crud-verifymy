package usecase

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func (u *UserUseCaseImpl) AuthenticateUser(email, password string) (string, error) {
	user, err := u.userRepo.FindByEmail(email)
	if err != nil {
		return "", err
	}

	// Verificar se a senha está correta
	if !checkPassword(password, user.Password) {
		return "", nil
	}

	token, err := generateAuthToken(user.ID, user.Profile)
	if err != nil {
		return "", err
	}

	// Retorne o token de autenticação
	return token, nil
}

// Função de geração de token de autenticação
func generateAuthToken(userID uint64, profile string) (string, error) {
	// Defina as informações do token, como claims e tempo de expiração
	claims := jwt.MapClaims{
		"id":      userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expira em 24 horas
		"profile": profile,
	}

	// Defina a assinatura do token (pode ser uma chave secreta)
	signingKey := []byte("VMYCRUDTEST")

	// Crie o token JWT com as informações e a assinatura
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func checkPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
