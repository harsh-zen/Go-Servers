package middleware

import (
	"bytes"
	"log"
	"net/http"
	"time"
)

//custom implementation of http.ResponseWriter
type responseCapturingWriter struct {
	http.ResponseWriter
	buffer          *bytes.Buffer
	responseSizeInBytes int64
}

// override Write method
func (w *responseCapturingWriter) Write(data []byte) (int, error) {
	bytesWritten, err := w.buffer.Write(data)
	if err != nil {
		return 0, err
	}
	w.responseSizeInBytes += int64(bytesWritten)
	return w.ResponseWriter.Write(data)
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		responseBuffer := &bytes.Buffer{}
		writerWithCapture := &responseCapturingWriter{
			ResponseWriter: w,
			buffer:         responseBuffer,
		}

		next.ServeHTTP(writerWithCapture, r)

		log.Printf("%s %s response time: %s, response size: %d bytes", r.Method, r.RequestURI, time.Since(startTime), writerWithCapture.responseSizeInBytes)
	})
}