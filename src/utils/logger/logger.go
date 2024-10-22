package logger

import (
	"context"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

var (
	_logger *logger
)

type Event struct {
	instance *zerolog.Logger
}

type logger struct {
	instance *zerolog.Logger
	options  loggerOptions
}

func (l *Event) Infof(format string, args ...interface{}) {
	l.instance.Info().Msgf(format, args...)
}

func (l *Event) Debugf(format string, args ...interface{}) {
	l.instance.Debug().Msgf(format, args...)
}

func (l *Event) Errorf(format string, args ...interface{}) {
	l.instance.Error().Msgf(format, args...)
}

func (l *Event) Warnf(format string, args ...interface{}) {
	l.instance.Warn().Msgf(format, args...)
}

func (l *Event) Panicf(format string, args ...interface{}) {
	l.instance.Panic().Msgf(format, args...)
}

func (l *Event) Printf(format string, args ...interface{}) {
	l.instance.Printf(format, args...)
}

func (l *Event) LogUseCaseError(u interface{}, start time.Time, err error) (interface{}, error) {
	duration := time.Since(start)
	l.Errorf("Concluído com erro: %T em %s. Output: %s", u, duration, err.Error())
	return nil, err
}

func (l *Event) LogUseCaseSuccess(u interface{}, start time.Time, output interface{}) (interface{}, error) {
	duration := time.Since(start)

	outputLog := output

	if output == nil {
		outputLog = "{}"
	}

	l.Debugf("Concluído com sucesso: %T em %s. Output: %s", u, duration, outputLog)

	return output, nil
}

// Starts and configures the log Instance
// and should be called at the start of the application.
//
// # Notice
//
// At this point in time the underlying logger used is zero log https://github.com/rs/zerolog
func Init(opt ...LoggerOption) {
	var (
		l    = zerolog.New(os.Stdout)
		opts = defaultLoggerOptions
	)

	for _, o := range opt {
		o.apply(&opts)
	}

	switch opts.env {
	case string(test):
		l.Output(io.Discard)
	case string(development):
		l = l.Output(customConsoleWriter())
		l.Level(zerolog.DebugLevel)
	case string(production), string(staging):
		l.Level(zerolog.InfoLevel)
	default:
		l.Level(zerolog.DebugLevel)
	}

	_logger = &logger{instance: &l, options: opts}
}

// Gets a basic logger Instance with no tracing.
//
// # Note
//   - This function is only recommended for simple logging, like application startup and configuration.
//   - For more in-depth logging, please refer to the Trace method.
func Get() *Event {
	var (
		l                 zerolog.Logger
		_, fileName, _, _ = runtime.Caller(1)
	)

	l = _logger.instance.With().Str("file", fileName).Logger()
	return &Event{instance: &l}
}

// Gets a logger Instance with a trace ID.
//
// This function is recomended for a more in-depth logging.
//
// # Note
//
// The trace id key name in the log is set by the Option TraceKey. The default value is "cid".
func Trace(pctx context.Context) (log *Event, c context.Context) {
	return trace(pctx)
}

func trace(pctx context.Context) (log *Event, ctx context.Context) {
	var (
		current, fileName, _, _ = runtime.Caller(2)
		callerName              = runtime.FuncForPC(current).Name()
		fileDetail              = strings.Split(callerName, "/")
		pkg                     = fileDetail[len(fileDetail)-2]
		method                  = fileDetail[len(fileDetail)-1]
		methods                 = strings.Split(method, ".")
		ctxKey                  = _logger.options.contextKey
		traceKey                = _logger.options.traceKey
		traceId                 string
	)
	if len(methods) > 0 {
		method = methods[len(methods)-1]
	}

	if existingKey, ok := pctx.Value(ctxKey).(string); ok && existingKey != "" {
		traceId = existingKey
	} else {
		traceId = uuid.New().String()
	}

	ctx = context.WithValue(pctx, contextKey(ctxKey), traceId)

	// This cleans up a bit our logging in a development setting
	if _logger.options.env == string(development) {
		l := _logger.
			instance.
			With().
			Ctx(pctx).
			Str(traceKey, traceId).
			Logger()
		log = &Event{instance: &l}
		return log, ctx
	}

	l := _logger.
		instance.
		With().
		Ctx(pctx).
		Str(traceKey, traceId).
		Str("file", fileName).
		Str("method", method).
		Str("pkg", pkg).
		Logger()

	log = &Event{instance: &l}

	return log, ctx
}

func customConsoleWriter() zerolog.ConsoleWriter {
	return zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
		// Omit the timestamp completely
		FormatTimestamp: func(i interface{}) string {
			return ""
		},
		FormatLevel: func(i interface{}) string {
			// Add color to log levels
			if ll, ok := i.(string); ok {
				switch ll {
				case "debug":
					return "\033[36mDEBUG\033[0m" // Cyan
				case "info":
					return "\033[32mINFO\033[0m" // Green
				case "warn":
					return "\033[33mWARN\033[0m" // Yellow
				case "error":
					return "\033[31mERROR\033[0m" // Red
				case "fatal":
					return "\033[35mFATAL\033[0m" // Magenta
				case "panic":
					return "\033[41mPANIC\033[0m" // Red background
				default:
					return ll
				}
			}
			return "UNKNOWN"
		},
		FormatMessage: func(i interface{}) string {
			// Add color to the actual log message
			return fmt.Sprintf("\033[34m%s\033[0m", i) // Blue
		},
		FormatFieldName: func(i interface{}) string {
			// Make field names bold
			return fmt.Sprintf("\033[1m%s:\033[0m", i)
		},
		FormatFieldValue: func(i interface{}) string {
			// Make field values yellow
			return fmt.Sprintf("\033[33m%s\033[0m", i)
		},
	}
}
