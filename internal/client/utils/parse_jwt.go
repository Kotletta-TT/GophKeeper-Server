package utils

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func GetUserIdOnJWT(token string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Проверяем, что токен подписан с использованием алгоритма HS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("подпись токена неверна")
		}
		// Возвращаем секретный ключ для проверки подписи токена
		return []byte("ваш_секретный_ключ"), nil
	})
	if err != nil {
		fmt.Println("Ошибка парсинга токена:", err)
		return "", nil
	}

	// Проверяем, валиден ли токен
	if !token.Valid {
		fmt.Println("Токен невалиден")
		return "", nil
	}

	// Получаем claims из токена
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("Невозможно получить claims из токена")
		return "", nil
	}

	// Получаем user_id из claims
	userID, ok := claims["user_id"].(string)
	if !ok {
		fmt.Println("user_id не найден в токене")
		return "", nil
	}
	return userID, nil
}
