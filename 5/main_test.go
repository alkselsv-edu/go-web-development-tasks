package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"testing"
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
	time.Sleep(200 * time.Millisecond)

	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		_ = app.ShutdownWithContext(ctx)
	}()

	testCases := []struct {
		name        string
		requestBody string
		wantCode    int
		wantIndex   int
		wantError   string
	}{
		{
			name:        "invalid JSON request body",
			requestBody: `{"numbe`,
			wantCode:    http.StatusBadRequest,
			wantIndex:   -1,
			wantError:   "Invalid JSON",
		},
		{
			name:        "target number is not found",
			requestBody: `{"numbers": [1, 2, 3, 5], "target": 4}`,
			wantCode:    http.StatusNotFound,
			wantIndex:   -1,
			wantError:   "Target was not found",
		},
		{
			name:        "numbers count is even. Target number is found",
			requestBody: `{"numbers": [1, 2, 3, 7, 99, 100, 250, 1000], "target": 99}`,
			wantCode:    http.StatusOK,
			wantIndex:   4,
			wantError:   "",
		},
		{
			name:        "numbers count is odd. Target number is found",
			requestBody: `{"numbers": [1, 2, 7, 99, 100, 250, 1000], "target": 99}`,
			wantCode:    http.StatusOK,
			wantIndex:   3,
			wantError:   "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, tErr := http.NewRequest(
				http.MethodPost,
				fmt.Sprintf("http://%s/search", addr),
				bytes.NewReader([]byte(tc.requestBody)),
			)
			req.Header.Set("Content-Type", "application/json")
			tr := require.New(t)
			tr.NoError(tErr)

			httpClient := http.Client{}
			resp, tErr := httpClient.Do(req)
			tr.NoError(tErr)
			defer resp.Body.Close()

			bodyBytes, tErr := io.ReadAll(resp.Body)
			tr.NoError(tErr)

			tr.Equal(tc.wantCode, resp.StatusCode)

			bsResp := BinarySearchResponse{}
			tErr = json.Unmarshal(bodyBytes, &bsResp)
			tr.NoError(tErr)

			tr.Equal(tc.wantIndex, bsResp.TargetIndex)
			tr.Equal(tc.wantError, bsResp.Error)
		})
	}
}
