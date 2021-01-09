package middlewares

import (
	"net/http"

	reqContext "github.com/4lexvav/hit-the-goal/context"
	"github.com/google/uuid"
)

const requestIDHeader = "X-Request-Id"

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			requestID := r.Header.Get(requestIDHeader)
			if requestID == "" {
				requestID = uuid.New().String()
			}

			ctx = reqContext.WithRequestID(ctx, requestID)

			next.ServeHTTP(w, r.WithContext(ctx))
		},
	)
}
