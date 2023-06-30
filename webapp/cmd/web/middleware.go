package main

import (
	"context"
	"net"
	"net/http"
	"strings"
)

type ctxKey string

const ctxUserKey ctxKey = "user_ip"

func (app *application) ipFromCtx(ctx context.Context) string {
	return ctx.Value(ctxUserKey).(string)
}

func (app *application) addIpToCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		ip, err := getIP(r)
		if err != nil {
			ip, _, _ = net.SplitHostPort(r.RemoteAddr)
			if len(ip) == 0 {
				ip = "unknown"
			}
			ctx = context.WithValue(r.Context(), ctxUserKey, ip)
		} else {
			ctx = context.WithValue(r.Context(), ctxUserKey, ip)
		}
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}

func getIP(r *http.Request) (string, error) {
	ip := r.Header.Get("X-Forwarded-For")
	if ip != "" {
		ipList := strings.Split(ip, ",")
		ip = strings.TrimSpace(ipList[0])
	} else {
		host, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr))
		if err != nil {
			return "", err
		}

		if host == "::1" {
			host = "127.0.0.1"
		}

		ip = host
	}

	return ip, nil
}
