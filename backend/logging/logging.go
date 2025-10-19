package logging

import (
	"log/slog"
	"os"
)

var infoOpts = slog.HandlerOptions{
	AddSource: true,
}
var Logger = slog.New(slog.NewJSONHandler(os.Stdout, &infoOpts))
