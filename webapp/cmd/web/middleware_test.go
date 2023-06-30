package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

var app application

func Test_application_addIpToCtx(t *testing.T) {
	tests := []struct {
		headerName string
		headerVal  string
		ip         string
		emptyIp    bool
	}{
		{"X-Forwarded-For", "127.0.0.1", "", false},
		{"", "", "", false},
		{"", "", "", true},
		{"", "", "hello:dear", false},
	}

	// Create dummy handler
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Make sure the value exists in ctx
		if _, ok := r.Context().Value(ctxUserKey).(string); !ok {
			t.Error("not string")
		}

		// Make sure we got an string back
		ip := r.Context().Value(ctxUserKey)
		if ip == nil {
			t.Error("not found")
		}
		t.Log(ip)
	})

	for _, tt := range tests {
		// Create handler to test
		handlerToTest := app.addIpToCtx(nextHandler)
		req := httptest.NewRequest("GET", "http://testing", nil)

		if tt.emptyIp {
			req.RemoteAddr = ""
		}

		if len(tt.headerName) > 0 {
			req.Header.Add(tt.headerName, tt.headerVal)
		}

		if len(tt.ip) > 0 {
			req.RemoteAddr = tt.ip
		}

		handlerToTest.ServeHTTP(httptest.NewRecorder(), req)
	}

}

func Test_application_ipFromCtx(t *testing.T) {
	// Create an app var of type application
	var app application
	// Get a context
	ctx := context.Background()
	// Put something in the context
	ctx = context.WithValue(ctx, ctxUserKey, "hello")
	// Call the ipFromCtx function
	ip := app.ipFromCtx(ctx)
	// Perform assertions
	if ip != "hello" {
		t.Error("wrong value returned from ipFromCtx")
	}
}
