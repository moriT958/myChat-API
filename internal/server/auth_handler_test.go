package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestSignup(t *testing.T) {
	reqBody := strings.NewReader(`{"username":"test-username-1","password":"test-password"}`)
	req, err := http.NewRequest(http.MethodPost, "/signup", reqBody)
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()
	srv := NewTodoServer(&MockAuthService{}, &MockChatService{})
	srv.Addr = "127.0.0.1:8080"
	srv.SignupHandler(rec, req)

	t.Run("correctly signup new user", func(t *testing.T) {
		want := map[string]string{"id": "test-userID-1"}
		var got map[string]string
		if err := json.NewDecoder(bytes.NewReader(rec.Body.Bytes())).Decode(&got); err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(want, got) {
			t.Errorf("expected %s, got %s", want, got)
		}
	})

	t.Run("correctly response code", func(t *testing.T) {
		want := http.StatusCreated
		got := rec.Code

		if want != got {
			t.Errorf("expect %d, got %d", want, got)
		}
	})
}

func TestLogin(t *testing.T) {
	reqBody := strings.NewReader(`{"username":"test-username-1","password":"test-password"}`)
	req, err := http.NewRequest(http.MethodPost, "/login", reqBody)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()

	srv := NewTodoServer(&MockAuthService{}, &MockChatService{})
	srv.Addr = "127.0.0.1:8080"

	srv.LoginHandler(rec, req)

	t.Run("correctly login and return token", func(t *testing.T) {
		want := map[string]string{"access_token": "test-token-1", "token_type": "Bearer"}
		var got map[string]string
		if err := json.NewDecoder(bytes.NewReader(rec.Body.Bytes())).Decode(&got); err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(want, got) {
			t.Errorf("expected %v, got %v", want, got)
		}
	})
}
