package logger

import (
	"log/slog"
)

func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "Ошибка",
		Value: slog.StringValue(err.Error()),
	}
}
