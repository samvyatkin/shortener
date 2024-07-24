package create

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"shortener/internal/app/handlers/create/mocks"
	"shortener/internal/app/storage"
	"shortener/internal/app/utils"
	"strings"
	"testing"
)

const (
	host = "http://example.com"
	URL  = "https://practicum.yandex.ru"
	UUID = "0190e4f6-ea36-71f6-a9c1-46010747f9aa"
)

func TestHandle(t *testing.T) {
	type want struct {
		code     int
		response string
		headers  map[string]string
	}

	tests := []struct {
		name          string
		method        string
		request       string
		body          string
		uuidGenerator utils.IdentifierGenerator
		storage       storage.Storage
		want          want
	}{
		{
			name:          "Create shorten URL (Success)",
			method:        http.MethodPost,
			request:       "/",
			body:          URL,
			uuidGenerator: mocks.NewUUIDGeneratorMock(UUID),
			storage:       mocks.NewInMemoryStorageMock(map[string]string{}),
			want: want{
				code:     http.StatusCreated,
				response: fmt.Sprintf("%s/%s", host, UUID),
				headers: map[string]string{
					"Content-Type": "text/plain",
				},
			},
		},
		{
			name:          "Get shorten URL (Success)",
			method:        http.MethodGet,
			request:       fmt.Sprintf("%s/%s", host, UUID),
			uuidGenerator: mocks.NewUUIDGeneratorMock(UUID),
			storage: mocks.NewInMemoryStorageMock(
				map[string]string{
					UUID: URL,
				},
			),
			want: want{
				code: http.StatusTemporaryRedirect,
				headers: map[string]string{
					"Content-Type": "text/plain",
					"Location":     URL,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := strings.NewReader(tt.body)

			req := httptest.NewRequest(tt.method, tt.request, body)
			rec := httptest.NewRecorder()

			h := New(tt.uuidGenerator, tt.storage)
			h.Handle(rec, req)

			res := rec.Result()
			assert.Equal(t, res.StatusCode, tt.want.code, "incorrect status code = %d, want = %d", res.StatusCode, tt.want.code)

			defer res.Body.Close()
			resBody, err := io.ReadAll(res.Body)

			require.NoError(t, err, "failed to read response body")
			assert.Equal(t, tt.want.response, string(resBody), "incorrect response body")

			for k, v := range tt.want.headers {
				assert.Equal(t, v, res.Header.Get(k), "header value is wrong %s, want = %s", res.Header.Get(k), v)
			}
		})
	}
}
