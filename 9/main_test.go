package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestPractice(t *testing.T) {
	r := require.New(t)

	app := SetupApp()

	ln, err := net.Listen("tcp", "127.0.0.1:3000")
	r.NoError(err)
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

	// Send 4 parallel requests: 2 to /foo and 2 to /bar
	// 1 request to /foo and 1 request to bar should be successful
	// other requests should be rejected with 429 status code

	requests := []*http.Request{
		request(r, "http://"+addr+"/foo"),
		request(r, "http://"+addr+"/bar"),
		request(r, "http://"+addr+"/foo"),
		request(r, "http://"+addr+"/bar"),
	}

	wg := sync.WaitGroup{}

	for _, req := range requests {
		wg.Add(1)
		go func(req *http.Request) {
			defer wg.Done()

			_, gErr := http.DefaultClient.Do(req)
			r.NoError(gErr)
		}(req)
	}

	wg.Wait()

	data, _ := os.ReadFile(".log")
	output := string(data)

	lines := strings.Split(output, "\n")

	expectedOutputs := []string{
		": GET /foo - 200",
		": GET /bar - 200",
		": GET /foo - 429",
		": GET /bar - 429",
	}

	for _, expectedOutput := range expectedOutputs {
		r.Contains(output, expectedOutput)

		for _, line := range lines {
			if !strings.HasSuffix(line, expectedOutput) {
				continue
			}

			requestID := strings.TrimSuffix(line, expectedOutput)
			r.True(IsValidUUID(requestID), "Invalid request ID in line: %s", line)
			break
		}
	}
}

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

func request(r *require.Assertions, url string) *http.Request {
	req, tErr := http.NewRequest(http.MethodGet, url, nil)
	r.NoError(tErr)

	return req
}
