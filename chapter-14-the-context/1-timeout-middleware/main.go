package main

import (
	"context"
	"net/http"
	"time"
)

func addTimeout(ms int) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx, cancel := context.WithTimeout(ctx, time.Millisecond*time.Duration(ms))
			defer cancel()
			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func laggyFunction(ctx context.Context) (string, error) {
	select {
	case <-time.After(2 * time.Second):
		return "Hi there with a delay!", nil
	case <-ctx.Done():
		return "Hi there took too long :(", ctx.Err()
	}
}

func main() {
	timeout := http.NewServeMux()
	timeout.HandleFunc("GET /hi", func(w http.ResponseWriter, r *http.Request) {
		out, err := laggyFunction(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusGatewayTimeout)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		w.Write([]byte(out))
	})

	inTime := http.NewServeMux()
	inTime.HandleFunc("GET /hi", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hi there immediately!"))
	})

	// timeoutMiddleware := addTimeout(800)
	timeoutMiddleware := addTimeout(1200)
	// timeoutMiddleware := addTimeout(4000)

	mux := http.NewServeMux()
	mux.Handle("/timesOut/", http.StripPrefix("/timesOut", timeoutMiddleware(timeout)))
	mux.Handle("/inTime/", http.StripPrefix("/inTime", inTime))

	s := http.Server{
		Addr:        ":8081",
		Handler:     timeoutMiddleware(mux),
		ReadTimeout: 500 * time.Millisecond,
		// WriteTimeout can't be smaller than addTimeout value, otherwise server will abandon request
		WriteTimeout: 3 * time.Second,
	}

	err := s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}
