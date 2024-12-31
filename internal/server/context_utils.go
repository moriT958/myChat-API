package server

import (
	"context"
	"net/http"
)

type userIDKey struct{}

func getUserID(ctx context.Context) string {
	ctxUserID := ctx.Value(userIDKey{})

	if userID, ok := ctxUserID.(string); ok {
		return userID
	}
	return ""
}

func setUserID(req *http.Request, userID string) *http.Request {
	ctx := req.Context()
	ctx = context.WithValue(ctx, userIDKey{}, userID)
	req = req.WithContext(ctx)
	return req
}
