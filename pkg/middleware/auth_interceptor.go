package middleware

import (
	"context"
	"log"
	"strings"

	"github.com/adhyttungga/go-grpc-test/pkg/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// excludeMethods List the full method name to exclude from the interceptor
// Next: Get login full method name
var excludeMethods = map[string]bool{
	"/auth.AuthService/Login": true,
}

// Auth unaryinterceptor
func AuthInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	log.Println("Method: ", info.FullMethod)

	// Check if the method is in the exclusion list
	if excludeMethods[info.FullMethod] {
		return handler(ctx, req)
	}

	// Extracting metadata
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "error missing metadata")
	}

	// Checking for the presence of x-link-service header
	lsHeader := md.Get("x-link-service")
	if len(lsHeader) == 0 {
		return nil, status.Error(codes.Unauthenticated, "error missing x-link-service header")
	}

	// Add section to ctx
	section := lsHeader[0]
	ctx = context.WithValue(ctx, "section", section)

	// Checking for the presence and validity of access token
	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		return nil, status.Error(codes.Unauthenticated, "error missing authhorizaton header")
	}

	var userId string
	accessToken := strings.TrimPrefix(authHeader[0], "Bearer ")
	if !utils.ValidateToken(accessToken, &userId) {
		return nil, status.Error(codes.Unauthenticated, "invalid access token")
	}

	// Set user id to context
	ctx = context.WithValue(ctx, "user_id", userId)
	return handler(ctx, req)
}
