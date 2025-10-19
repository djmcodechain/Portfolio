package logging

import (
	"log/slog"
	"os"
)

var infoOpts = slog.HandlerOptions{
	AddSource: true,
}
var info = slog.New(slog.NewJSONHandler(os.Stdout, &infoOpts))
