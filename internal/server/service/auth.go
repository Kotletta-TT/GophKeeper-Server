package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// AuthService - сервис аутентификации
type AuthService struct {
	// Секретный ключ для подписи и верификации токенов
	secretKey []byte
}

// NewAuthService - конструктор для создания экземпляра сервиса аутентификации
func NewAuthService(secretKey string) *AuthService {
	return &AuthService{
		secretKey: []byte(secretKey),
	}
}

// ValidateToken - метод для проверки токена аутентификации
func (as *AuthService) ValidateToken(tokenString string) (*jwt.Token, error) {
	// Парсим токен
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Проверяем, что алгоритм подписи подходит
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		// Возвращаем секретный ключ для верификации
		return as.secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	// Проверяем, что токен валиден
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token, nil
}

// GenerateToken - метод для генерации JWT токена
func (as *AuthService) GenerateToken(userID string) (string, error) {
	// Создаем структуру с метаданными токена
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Токен действителен в течение 24 часов
		"iat":     time.Now().Unix(),
	})

	// Подписываем токен с использованием секретного ключа
	tokenString, err := claims.SignedString(as.secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
