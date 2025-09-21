package middleware

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

func LoggingInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	start := time.Now()

	log.Printf("Received RPC call: Method=%s", info.FullMethod)

	resp, err := handler(ctx, req)

	duration := time.Since(start)

	if err != nil {
		log.Printf("RPC call failed: Method=%s, Error=%v, Duration=%s", info.FullMethod, err, duration)
	} else {
		log.Printf("RPC call completed: Method=%s, Duration=%s", info.FullMethod, duration)
	}

	return resp, err
}
