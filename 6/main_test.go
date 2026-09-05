package main

import (
	"io"
	"net"
	"net/http"
	"strings"
	"testing"

	"context"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPractice(t *testing.T) {
	app := SetupApp()

	ln, err := net.Listen("tcp", "127.0.0.1:3000")
	require.NoError(t, err)
	addr := ln.Addr().String()

	go func() {
		_ = app.Listener(ln)
	}()
	// Даем серверу время на запуск
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
		wantBody      string
	}{
		{
			name:          "link not found",
			requestPath:   "/links/test",
			requestMethod: http.MethodGet,
			wantCode:      http.StatusNotFound,
			wantBody:      "Link not found",
		},
		{
			name:          "invalid JSON body",
			requestPath:   "/links",
			requestBody:   "",
			requestMethod: http.MethodPost,
			wantCode:      http.StatusBadRequest,
			wantBody:      "Invalid JSON",
		},
		{
			name:          "create a link",
			requestPath:   "/links",
			requestBody:   `{"external": "https://ly.com/kwz", "internal": "https://google.com?search=somequery"}`,
			requestMethod: http.MethodPost,
			wantCode:      http.StatusOK,
			wantBody:      "OK",
		},
		{
			name:          "get the created link",
			requestPath:   "/links/https%3A%2F%2Fly.com%2Fkwz",
			requestBody:   "",
			requestMethod: http.MethodGet,
			wantCode:      http.StatusOK,
			wantBody:      `{"internal":"https://google.com?search=somequery"}`,
		},
		{
			name:          "update link",
			requestPath:   "/links",
			requestBody:   `{"external": "https://ly.com/kwz", "internal": "https://google.com?search=updatedquery"}`,
			requestMethod: http.MethodPost,
			wantCode:      http.StatusOK,
			wantBody:      "OK",
		},
		{
			name:          "get updated link",
			requestPath:   "/links/https%3A%2F%2Fly.com%2Fkwz",
			requestBody:   "",
			requestMethod: http.MethodGet,
			wantCode:      http.StatusOK,
			wantBody:      `{"internal":"https://google.com?search=updatedquery"}`,
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

			bodyBytes, tErr := io.ReadAll(resp.Body)
			tr.NoError(tErr)

			tr.Equal(tc.wantCode, resp.StatusCode)
			tr.Equal(tc.wantBody, string(bodyBytes))
		})
	}
}
