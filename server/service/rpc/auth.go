package rpc

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type AuthHandler struct {
	ClientSecret string
}
type IOMHandler struct {
	Auth *AuthHandler
}

func (a *AuthHandler) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{"client_secret": a.ClientSecret}, nil
}

func (a *AuthHandler) RequireTransportSecurity() bool {
	return false
}

func (a *AuthHandler) Check(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Errorf(codes.Unauthenticated, "获取 metaData 失败")
	}

	var clientSecret string
	if value, ok := md["client_secret"]; ok {
		clientSecret = value[0]
		return clientSecret, nil
	} else {
		// 没有客户端密钥
		return "", status.Errorf(codes.Unauthenticated, "客户端密钥错误或不存在")
	}
}
