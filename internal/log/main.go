package log

import (
	"github.com/rs/zerolog"
	"os"
	"time"
)

var Log zerolog.Logger

func init() {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	Log = zerolog.New(output).With().Timestamp().Logger()
}
