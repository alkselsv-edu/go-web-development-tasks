package main

import (
	"context"
	"net"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPractice(t *testing.T) {
	users = map[string]User{}
	app := SetupApp()

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)
	addr := ln.Addr().String()

	go func() {
		_ = app.Listener(ln)
	}()
	time.Sleep(200 * time.Millisecond)

	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		_ = app.ShutdownWithContext(ctx)
	}()

	testCases := []struct {
		name          string
		requestPath   string
		requestBody   string
		requestMethod string
		wantCode      int
	}{
		{
			name:          "positive case",
			requestPath:   "/users",
			requestMethod: http.MethodPost,
			requestBody:   `{"username": "john1", "email": "john@doe.com", "age": 18, "country": "Japan"}`,
			wantCode:      http.StatusOK,
		},
		{
			name:          "positive case vietnam",
			requestPath:   "/users",
			requestMethod: http.MethodPost,
			requestBody:   `{"username": "alice", "email": "alice@doe.com", "age": 130, "country": "Vietnam"}`,
			wantCode:      http.StatusOK,
		},
		{
			name:          "positive case malaysia",
			requestPath:   "/users",
			requestMethod: http.MethodPost,
			requestBody:   `{"username": "bob2", "email": "bob@doe.com", "age": 25, "country": "Malaysia"}`,
			wantCode:      http.StatusOK,
		},
		{
			name:          "positive case thailand",
			requestPath:   "/users",
			requestMethod: http.MethodPost,
			requestBody:   `{"username": "carol3", "email": "carol@doe.com", "age": 40, "country": "Thailand"}`,
			wantCode:      http.StatusOK,
		},
		{
			name:          "overwrite existing username",
			requestPath:   "/users",
			requestMethod: http.MethodPost,
			requestBody:   `{"username": "john1", "email": "john.updated@doe.com", "age": 26, "country": "Japan"}`,
			wantCode:      http.StatusOK,
		},
		{
			name:          "non-specified username",
			requestPath:   "/users",
			requestMethod: http.MethodPost,
			requestBody:   `{"email": "john@doe.com", "age": 25, "country": "Japan"}`,
			wantCode:      http.StatusUnprocessableEntity,
		},
		{
			name:          "invalid username uppercase and underscore",
			requestPath:   "/users",
			requestMethod: http.MethodPost,
			requestBody:   `{"username": "John_1", "email": "john@doe.com", "age": 25, "country": "Japan"}`,
			wantCode:      http.StatusUnprocessableEntity,
		},
		{
			name:          "non-specified email",
			requestPath:   "/users",
			requestMethod: http.MethodPost,
			requestBody:   `{"username": "john1", "age": 25, "country": "Japan"}`,
			wantCode:      http.StatusUnprocessableEntity,
		},
		{
			name:          "invalid email",
			requestPath:   "/users",
			requestMethod: http.MethodPost,
			requestBody:   `{"username": "john1", "email": "john.com", "age": 25, "country": "Japan"}`,
			wantCode:      http.StatusUnprocessableEntity,
		},
		{
			name:          "invalid age < 18",
			requestPath:   "/users",
			requestMethod: http.MethodPost,
			requestBody:   `{"username": "john1", "email": "john@doe.com", "age": 17, "country": "Japan"}`,
			wantCode:      http.StatusUnprocessableEntity,
		},
		{
			name:          "invalid age > 130",
			requestPath:   "/users",
			requestMethod: http.MethodPost,
			requestBody:   `{"username": "john1", "email": "john@doe.com", "age": 131, "country": "Japan"}`,
			wantCode:      http.StatusUnprocessableEntity,
		},
		{
			name:          "invalid country",
			requestPath:   "/users",
			requestMethod: http.MethodPost,
			requestBody:   `{"username": "john1", "email": "john@doe.com", "age": 130, "country": "Unknown"}`,
			wantCode:      http.StatusUnprocessableEntity,
		},
		{
			name:          "non-specified age",
			requestPath:   "/users",
			requestMethod: http.MethodPost,
			requestBody:   `{"username": "john1", "email": "john@doe.com", "country": "Japan"}`,
			wantCode:      http.StatusUnprocessableEntity,
		},
		{
			name:          "non-specified country",
			requestPath:   "/users",
			requestMethod: http.MethodPost,
			requestBody:   `{"username": "john1", "email": "john@doe.com", "age": 18}`,
			wantCode:      http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, tErr := http.NewRequest(
				tc.requestMethod,
				"http://"+addr+tc.requestPath,
				strings.NewReader(tc.requestBody),
			)
			tr := require.New(t)
			tr.NoError(tErr)

			req.Header.Set("Content-Type", "application/json")

			httpClient := http.Client{}
			resp, tErr := httpClient.Do(req)
			tr.NoError(tErr)
			defer resp.Body.Close()

			tr.Equal(tc.wantCode, resp.StatusCode)
		})
	}

	tr := require.New(t)
	user, ok := users["john1"]
	tr.True(ok)
	tr.Equal("john.updated@doe.com", user.Email)
	tr.Equal(26, user.Age)
}
