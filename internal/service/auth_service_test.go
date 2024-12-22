package service_test

import (
	"context"
	"myChat-API2/internal/service"
	"myChat-API2/internal/service/testdata"
	"testing"
)

func TestSignup(t *testing.T) {
	userRepo := new(testdata.MockUserRepository)
	authService := service.NewAuthService(userRepo)

	ctx := context.Background()
	var tests = []struct {
		name     string
		username string
		password string
		wantErr  bool
	}{
		{name: "success signup", username: "user1", password: "pass1", wantErr: false},
		{name: "password create fail", username: "user1", password: "", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gotErr bool
			_, err := authService.Signup(ctx, tt.name, tt.password)
			if err != nil {
				gotErr = true
			}
			if gotErr != tt.wantErr {
				t.Errorf("Signup(ctx,%s,%s): want err(%t) ,got err(%t)", tt.username, tt.password, tt.wantErr, gotErr)
			}
		})
	}
}
