package handler

import (
	"context"
	"net/http"
)

type UsernameKey struct{}

func GetUsername(ctx context.Context) string {
	usernameCtx := ctx.Value(UsernameKey{})

	if usernameStr, ok := usernameCtx.(string); ok {
		return usernameStr
	}
	return ""
}

func SetUsername(req *http.Request, username string) *http.Request {
	ctx := req.Context()
	ctx = context.WithValue(ctx, UsernameKey{}, username)
	req = req.WithContext(ctx)
	return req
}
