package lib

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
)

// SLogBufferHandler is a custom log handler that writes logs to a bytes.Buffer.
type SLogBufferHandler struct {
	buffer     *bytes.Buffer
	attrs      []slog.Attr
	FormatFunc func(record slog.Record) string
}

// NewSLogBufferHandler creates a new SLogBufferHandler.
func NewSLogBufferHandler() *SLogBufferHandler {
	return &SLogBufferHandler{
		buffer: new(bytes.Buffer),
		FormatFunc: func(record slog.Record) string {
			return fmt.Sprintf("[%v] %v: %v\n",
				record.Level,
				record.Time,
				record.Message,
			)
		},
	}
}

// Enabled checks if the log level is enabled.
func (h *SLogBufferHandler) Enabled(ctx context.Context, level slog.Level) bool {
	// Here, you can implement logic to filter log levels if needed.
	// For simplicity, this example enables all log levels.
	return true
}

// Handle processes the log record.
func (h *SLogBufferHandler) Handle(ctx context.Context, record slog.Record) error {
	// Format the log record and write it to the buffer.
	// You may need to adjust the formatting based on your requirements.
	_, _ = fmt.Fprintf(h.buffer, h.FormatFunc(record))
	return nil
}

// WithAttrs returns a new handler with added attributes.
func (h *SLogBufferHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	newHandler := *h
	newHandler.attrs = append(newHandler.attrs, attrs...)
	return &newHandler
}

// WithGroup returns a new handler with the given group name.
// If the group name is empty, it returns the same handler.
func (h *SLogBufferHandler) WithGroup(name string) slog.Handler {
	if name == "" {
		return h
	}
	// Implement group handling logic if needed.
	return h // This is a simple implementation.
}

// Content returns the content of the log buffer.
func (h *SLogBufferHandler) Content() string {
	return h.buffer.String()
}
