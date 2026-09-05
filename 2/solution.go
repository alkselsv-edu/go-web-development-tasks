package main

import (
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

// StartServer запускает сервер и возвращает его указатель для управления из тестов
func StartServer() *http.Server {
	cwd, _ := os.Getwd()
	logFile := filepath.Join(cwd, ".log")
	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	file, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		logger.Fatal(err)
	}
	// Не закрываем file здесь, чтобы логгер мог писать в файл до завершения сервера

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Go to /sum"))
	})

	mux.HandleFunc("/sum", func(w http.ResponseWriter, r *http.Request) {
		// BEGIN (write your solution here)
		// END
	})

	server := &http.Server{
		Addr:              ":3000",
		Handler:           mux,
		ReadHeaderTimeout: 10 * time.Second,
	}

	go func() {
		logWithPort := logrus.WithFields(logrus.Fields{
			"port": "3000",
		})
		logWithPort.Info("Starting a web-server on port")
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			logWithPort.Error(err)
		}
		// Закрываем файл лога после завершения сервера
		file.Close()
	}()

	return server
}

func main() {
	StartServer()
	select {} // Блокируем main, чтобы сервер не завершился сразу
}
