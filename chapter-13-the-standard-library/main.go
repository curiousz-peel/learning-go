package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type logger struct {
	l *slog.Logger
}

func newJSONLogger() *logger {
	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	return &logger{
		l: slog.New(logHandler),
	}
}

func (logger *logger) logAddr(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.l.Info("address log", "address", r.RemoteAddr, "method", r.Method)
		h.ServeHTTP(w, r)
	})
}

func buildJSON(now time.Time) string {
	timeStruct := struct {
		DayOfWeek  string `json:"day_of_Week"`
		DayOfMonth int    `json:"day_of_month"`
		Month      string `json:"month"`
		Year       int    `json:"year"`
		Hour       int    `json:"hour"`
		Minute     int    `json:"minute"`
		Second     int    `json:"second"`
	}{
		DayOfWeek:  now.Weekday().String(),
		DayOfMonth: now.Day(),
		Month:      now.Month().String(),
		Year:       now.Year(),
		Hour:       now.Hour(),
		Minute:     now.Minute(),
		Second:     now.Second(),
	}
	jsonTime, _ := json.Marshal(timeStruct)
	return string(jsonTime)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /time", func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		var response string
		switch respFormat := r.Header.Get("Accept"); respFormat {
		case "application/json":
			w.Header().Set("Content-Type", "application/json")
			response = buildJSON(now)
		default:
			w.Header().Set("Content-Type", "text/plain")
			response = now.Format(time.RFC3339)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	})

	log := newJSONLogger()

	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 2 * time.Second,
		IdleTimeout:  3 * time.Second,
		Handler:      log.logAddr(mux),
	}

	err := s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}
