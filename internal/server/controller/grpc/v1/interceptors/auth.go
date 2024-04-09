package interceptors

import (
	"context"

	"github.com/Kotletta-TT/GophKeeper/internal/server/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// AuthInterceptor - интерфейс для аутентификации
type AuthInterceptor struct {
	AuthService service.AuthService // AuthService - сервис аутентификации
	Methods     map[string]bool
}

// Unary - унарный серверный интерфейс
func (ai *AuthInterceptor) Unary(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// Проверяем, должен ли интерсептор быть применен к этому методу
	if ai.Methods[info.FullMethod] {
		// Получаем метаданные из контекста
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
		}

		// Получаем токен аутентификации из метаданных
		token := md.Get("authorization")
		if len(token) == 0 {
			return nil, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
		}

		// Проверяем токен аутентификации с помощью сервиса аутентификации
		if _, err := ai.AuthService.ValidateToken(token[0]); err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "authentication failed: %v", err)
		}
	}

	// Продолжаем выполнение цепочки обработчиков
	return handler(ctx, req)
}
