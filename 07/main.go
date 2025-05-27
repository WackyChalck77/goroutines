package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// реализуем 2 функции, шифрование в контекст токена
func AddJWTToContext(ctx context.Context, userID int) (context.Context, error) {
	//создаем токен JWT
	token := jwt.New(jwt.SigningMethodHS256)

	// Создание кастомного клейма с номером пользователя
	claims := token.Claims.(jwt.MapClaims)
	claims["userID"] = userID
	//Пропускаем установку срока действия
	// Подпись токена с секретным ключом
	secretKey := []byte("highLowTopSlow")
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println("Ошибка при подписи токена")
		return ctx, err
	}
	//запишем токен в контекст
	ctx = context.WithValue(ctx, "jwt_token", tokenString)
	fmt.Println("Контекст после добавления в него токена\n", ctx)
	//Token &{ 0xc0000080c0 map[alg:HS256 typ:JWT] map[userID:45]  false}
	//ctx  context.Background.WithValue(jwt_token, eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVC
	//J9.eyJ1c2VySUQiOjQ1fQ.UI7WKzQZzyVkaokY_DCaTuxJiOte4sjQiHMYh8VJbPk)
	return ctx, err
}

// дешифрование токена из контекста
func ExtractUserIDFromContext(ctx context.Context) (int, error) {
	//получаем токен из контекста
	tokenString, ok := ctx.Value("jwt_token").(string)
	if !ok {
		return 0, fmt.Errorf("токена в контексте нет")
	}
	//парсинг токена
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("highLowTopSlow"), nil
	})
	if err != nil {
		return 0, fmt.Errorf("error parsing token: %w", err)
	}
	//проверяем клэймы
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, fmt.Errorf("invalid token or missing claims")
	}

	userID, ok := claims["userID"].(float64) //преобразование
	if !ok {
		return 0, fmt.Errorf("невозможно получить userID из токена")
	}
	return int(userID), nil
}

func main() {
	//создаем пустой контекст
	ctx := context.Background()
	var n int
	fmt.Println("Добавьте токен")
	fmt.Scanln(&n)
	ctx, err := AddJWTToContext(ctx, n)
	if err != nil {
		log.Printf("Ошибка при добавлении токена в контекст: %v", err)
		return
	}

	// fmt.Println("Выполняю извлечение из токена")
	// userID, err := ExtractUserIDFromContext(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//создадим горутину для извлечения из контекста
	go func(ctx context.Context) {
		userID, err := ExtractUserIDFromContext(ctx)
		if err != nil {
			fmt.Println("Ошибка в горутине, вытаскивающей токен из контекста", err)
			return
		}
		fmt.Println("Извлечен ID из контекста,", userID)
	}(ctx)

	time.Sleep(time.Microsecond * 500)

	//	fmt.Println("User ID:", userID)
}
