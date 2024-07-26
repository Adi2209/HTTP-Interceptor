package main

import (
	"bytes"
	"io"
	"net/http"

	"github.com/golang/snappy"
)

func compressionInterceptor(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Accept-Encoding") == "snappy" {
			var buf bytes.Buffer
			snappyWriter := snappy.NewBufferedWriter(&buf)
			wrappedWriter := &snappyResponseWriter{ResponseWriter: w, Writer: snappyWriter, Buffer: &buf}

			next.ServeHTTP(wrappedWriter, r)

			snappyWriter.Close()
			w.Header().Set("Content-Encoding", "snappy")
			io.Copy(w, &buf)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

type snappyResponseWriter struct {
	http.ResponseWriter
	Writer io.Writer
	Buffer *bytes.Buffer
}

func (w *snappyResponseWriter) Write(data []byte) (int, error) {
	return w.Writer.Write(data)
}
