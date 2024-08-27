package middleware

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

type (
	compressWriter struct {
		writer     http.ResponseWriter
		gzipWriter *gzip.Writer
	}
	compressReader struct {
		reader     io.ReadCloser
		gzipReader *gzip.Reader
	}
)

func (cw *compressWriter) Header() http.Header {
	return cw.writer.Header()
}

func (cw *compressWriter) WriteHeader(statusCode int) {
	if statusCode < 300 {
		cw.writer.Header().Set("Content-Encoding", "gzip")
	}
	cw.writer.WriteHeader(statusCode)
}

func (cw *compressWriter) Write(data []byte) (int, error) {
	return cw.gzipWriter.Write(data)
}

func (cw *compressWriter) Close() error {
	if cw.gzipWriter != nil {
		return cw.gzipWriter.Close()
	}

	return nil
}

func newCompressWriter(w http.ResponseWriter) *compressWriter {
	return &compressWriter{
		writer:     w,
		gzipWriter: gzip.NewWriter(w),
	}
}

func (cr *compressReader) Read(p []byte) (int, error) {
	return cr.gzipReader.Read(p)
}

func (cr *compressReader) Close() error {
	if err := cr.reader.Close(); err != nil {
		return err
	}
	return cr.gzipReader.Close()
}

func newCompressReader(r io.ReadCloser) (*compressReader, error) {
	gzipReader, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}

	return &compressReader{
		reader:     r,
		gzipReader: gzipReader,
	}, nil
}

func WithCompress() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ow := w

			acceptEncoding := r.Header.Get("Accept-Encoding")
			supportsGzip := strings.Contains(acceptEncoding, "gzip")
			if supportsGzip {
				cw := newCompressWriter(w)
				ow = cw
				defer cw.Close()
			}

			contentEncoding := r.Header.Get("Content-Encoding")
			sendsGzip := strings.Contains(contentEncoding, "gzip")
			if sendsGzip {
				cr, err := newCompressReader(r.Body)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}

				r.Body = cr
				defer cr.Close()
			}

			next.ServeHTTP(ow, r)
		})
	}
}
