package logger

import (
	"os"
	"strings"

	"go.elastic.co/ecszap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type (
	Logger struct {
		level uint32
	}

	Lvl uint8
)

const (
	DEBUG Lvl = iota + 1
	INFO
	WARN
	ERROR
	DPANIC
	PANIC
	FATAL
)

var level zapcore.Level

func init() {
	l := strings.ToUpper(os.Getenv("LOG_LEVEL"))
	switch l {
	// DebugLevel are typically voluminous, and are usualy disabled in production
	case "DEBUG":
		level = zap.DebugLevel
	// InfoLevel is the default logging priority
	case "INFO":
		level = zap.InfoLevel
	// WarnLevel logs are more important than Info, but don't need individual human review
	case "WARN":
		level = zap.WarnLevel
	// ErrorLevel logs are high-priority. If an application is running smoothly, it shouldn't
	// generate any error-level logs.
	case "ERROR":
		level = zap.ErrorLevel
	// DPanicLevel logs are particularly import errors. In development, the logger
	// panics after writing the message.
	case "DPANIC":
		level = zap.DPanicLevel
	// PanicLevel logs a message, then panics.
	case "PANIC":
		level = zap.PanicLevel
	// FatalLevel logs a message, then calls os.Exit(1)
	case "FATAL":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}
}

func Neww() (l *Logger) {
	l = &Logger{
		level: uint32(INFO),
	}
	
	return
}

// When performance and type safety are critical, use the `New`. It's even faster than
// the `SugaredLogger` and allocates far less, but it only supports structuted logging.
func New() *zap.Logger {
	encoderConfig := ecszap.NewDefaultEncoderConfig()
	core := ecszap.NewCore(encoderConfig, os.Stdout, level)
	logger := zap.New(core, zap.AddCaller())
	return logger
}

// In contexts where performance is nice, but not critical, use the `SugaredLogger`.
// It is 4-10x faster than other structured logging packages and includes both structured
// and `printf`-style APIs.
func NewWithSugaredLogger() *zap.SugaredLogger {
	log := New()
	logger := log.Sugar()
	return logger
}
