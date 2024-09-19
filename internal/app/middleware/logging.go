package middleware

import (
	"go.uber.org/zap"
	"net/http"
	"time"
)

type (
	data struct {
		status int
		size   int
	}

	writer struct {
		writer *http.ResponseWriter
		data   *data
	}
)

func (w *writer) Header() http.Header {
	return (*w.writer).Header()
}

func (w *writer) Write(p []byte) (int, error) {
	size, err := (*w.writer).Write(p)
	w.data.size += size
	return size, err
}

func (w *writer) WriteHeader(statusCode int) {
	(*w.writer).WriteHeader(statusCode)
	w.data.status = statusCode
}

func WithLogging(sugar *zap.SugaredLogger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			data := data{
				status: 0,
				size:   0,
			}
			writer := writer{
				writer: &w,
				data:   &data,
			}

			defer func() {
				sugar.Infow(
					"Request completed",
					"uri",
					r.RequestURI,
					"method",
					r.Method,
					"status",
					data.status,
					"duration",
					time.Since(start), "size",
					data.size,
				)
			}()

			next.ServeHTTP(&writer, r)
		})
	}
}
