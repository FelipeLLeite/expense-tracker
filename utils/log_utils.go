package utils

import (
	"io"
	"log"
	"log/slog"
	"os"

	"github.com/gookit/color"
	// "image/color"
)

var (
	LogWarning *log.Logger
	LogInfo    *log.Logger
	LogError   *log.Logger
	OutputFile *os.File
	LogHmm     *slog.Logger
)

func init() {
	LogWarning = log.New(io.MultiWriter(os.Stdout, io.Discard), color.FgYellow.Render("[WARNING] "), log.Ldate|log.Ltime|log.Lshortfile)
	LogInfo = log.New(io.MultiWriter(os.Stdout, io.Discard), color.FgGreen.Render("[INFO] "), log.Ldate|log.Ltime|log.Lshortfile)
	LogError = log.New(io.MultiWriter(os.Stdout, io.Discard), color.FgRed.Render("[ERROR] "), log.Ldate|log.Ltime|log.Lshortfile)
	LogHmm = slog.New(
		slog.NewMultiHandler(
			slog.NewJSONHandler(os.Stdout, nil),
			// slog.DiscardHandler,
		),
	)
}
