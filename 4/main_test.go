package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
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
	time.Sleep(200 * time.Millisecond)

	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		_ = app.ShutdownWithContext(ctx)
	}()

	testCases := []struct {
		name          string
		postID        string
		requestMethod string
		wantCode      int
		want          string
	}{
		{
			name:          "post not found",
			postID:        "12345",
			requestMethod: fiber.MethodGet,
			wantCode:      http.StatusNotFound,
			want:          "",
		},
		{
			name:          "post increment 1",
			postID:        "12345",
			requestMethod: fiber.MethodPost,
			wantCode:      http.StatusCreated,
			want:          "1",
		},
		{
			name:          "post increment 2",
			postID:        "12345",
			requestMethod: fiber.MethodPost,
			wantCode:      http.StatusOK,
			want:          "2",
		},
		{
			name:          "get incremented post",
			postID:        "12345",
			requestMethod: fiber.MethodGet,
			wantCode:      http.StatusOK,
			want:          "2",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			url := fmt.Sprintf("http://%s/likes/%s", addr, tc.postID)
			req, tErr := http.NewRequest(tc.requestMethod, url, nil)
			tr := require.New(t)
			tr.NoError(tErr)

			httpClient := http.Client{}
			resp, tErr := httpClient.Do(req)
			tr.NoError(tErr)
			defer resp.Body.Close()

			tr.Equal(tc.wantCode, resp.StatusCode)
			if tc.want != "" {
				body, rErr := io.ReadAll(resp.Body)
				tr.NoError(rErr)
				tr.Equal(tc.want, string(body))
			}
		})
	}
}
