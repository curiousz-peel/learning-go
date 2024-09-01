package main

import (
	"context"
	"fmt"
	"net/http"
)

type logKey int

const (
	_ logKey = iota
	key
)

type Level string

const (
	Debug Level = "debug"
	Info  Level = "info"
)

func setLogLevel(ctx context.Context, level Level) context.Context {
	return context.WithValue(ctx, key, level)
}

func getLogLevel(ctx context.Context) (Level, bool) {
	lvl, ok := ctx.Value(key).(Level)
	return lvl, ok
}

func logMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logLvl := r.URL.Query().Get("log_level")
		ctx := setLogLevel(r.Context(), Level(logLvl))
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func Log(ctx context.Context, level Level, message string) {
	var inLevel Level
	inLevel, ok := getLogLevel(ctx)
	if !ok {
		return
	}
	if level == Debug && inLevel == Debug {
		fmt.Println(message)
	}
	if level == Info && (inLevel == Debug || inLevel == Info) {
		fmt.Println(message)
	}
}
