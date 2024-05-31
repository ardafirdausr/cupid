package logger

import (
	"os"

	"github.com/rs/zerolog"
)

var Log zerolog.Logger

func init() {
	Log = zerolog.New(os.Stdout).
		Level(zerolog.InfoLevel).
		With().
		Caller().
		Timestamp().
		Logger()
}

func SetLogLevel(level zerolog.Level) {
	Log = Log.Level(level)
}

func SetOutput(output *os.File) {
	Log = Log.Output(output)
}

func SetFormat(format string) {
	Log = Log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
}
