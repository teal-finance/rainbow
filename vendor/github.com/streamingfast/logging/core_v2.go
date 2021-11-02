package logging

import (
	"fmt"
	"net/http"
	"os"

	"github.com/blendle/zapdriver"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/crypto/ssh/terminal"
)

type loggerFactory func(name string, level zapcore.Level) *zap.Logger

// This v2 version of `core.go` is a work in progress without any backwar compatibility
// version. It might not made it to an official version of the library so you can depend
// at your own risk.

type loggerOptions struct {
	encoderVerbosity         *int
	level                    *zap.AtomicLevel
	reportAllErrors          *bool
	serviceName              *string
	switcherServerAutoStart  *bool
	switcherServerListenAddr string
	zapOptions               []zap.Option
	registerOptions          []RegisterOption

	// Use internally only, no With... value defined for it
	loggerName string
}

func newLoggerOptions(shortName string, opts ...LoggerOption) loggerOptions {
	loggerOptions := loggerOptions{switcherServerListenAddr: "127.0.0.1:1065"}
	for _, opt := range opts {
		opt.apply(&loggerOptions)
	}

	if loggerOptions.reportAllErrors == nil {
		WithReportAllErrors().apply(&loggerOptions)
	}

	if loggerOptions.serviceName == nil && shortName != "" {
		WithServiceName(shortName).apply(&loggerOptions)
	}

	if loggerOptions.switcherServerAutoStart == nil && isProductionEnvironment() {
		WithSwitcherServerAutoStart().apply(&loggerOptions)
	}

	return loggerOptions
}

type LoggerOption interface {
	apply(o *loggerOptions)
}

type loggerFuncOption func(o *loggerOptions)

func (f loggerFuncOption) apply(o *loggerOptions) {
	f(o)
}

func WithSwitcherServerAutoStart() LoggerOption {
	return loggerFuncOption(func(o *loggerOptions) {
		o.switcherServerAutoStart = ptrBool(true)
	})
}

func WithAtomicLevel(level zap.AtomicLevel) LoggerOption {
	return loggerFuncOption(func(o *loggerOptions) {
		o.level = ptrLevel(level)
	})
}

func WithReportAllErrors() LoggerOption {
	return loggerFuncOption(func(o *loggerOptions) {
		o.reportAllErrors = ptrBool(true)
	})
}

func WithServiceName(name string) LoggerOption {
	return loggerFuncOption(func(o *loggerOptions) {
		o.serviceName = ptrString(name)
	})
}

func WithZapOption(zapOption zap.Option) LoggerOption {
	return loggerFuncOption(func(o *loggerOptions) {
		o.zapOptions = append(o.zapOptions, zapOption)
	})
}

func WithOnUpdate(onUpdate func(newLogger *zap.Logger)) LoggerOption {
	return loggerFuncOption(func(o *loggerOptions) {
		o.registerOptions = append(o.registerOptions, RegisterOnUpdate(onUpdate))
	})
}

// LibraryLogger creates a new no-op logger (via `zap.NewNop`) and automatically registered it
// withing the logging registry with a tracer that can be be used for conditionally tracing
// code.
func LibraryLogger(shortName string, packageID string, logger **zap.Logger, registerOptions ...RegisterOption) Tracer {
	return libraryLogger(globalRegistry, shortName, packageID, logger, registerOptions...)
}

func libraryLogger(registry *registry, shortName string, packageID string, logger **zap.Logger, registerOptions ...RegisterOption) Tracer {
	return register2(registry, shortName, packageID, logger, registerOptions...)
}

// ApplicationLogger should be used to get a logger for a top-level binary application which will
// immediately activate all registered loggers with a logger. The actual logger for all component
// used is deried based on the identified environment and from environment variables.
//
// Here the set of rules used and the outcome they are giving:
//
//  1. If a production environment is detected (for now, only checking if file /.dockerenv exists)
//     Use a JSON StackDriver compatible format
//
//  2. Otherwise
//     Use a developer friendly colored format
//
//
// *Note* The ApplicationLogger should be start only once per processed. That could be enforced
//        in the future.
func ApplicationLogger(shortName string, packageID string, logger **zap.Logger, opts ...LoggerOption) Tracer {
	return applicationLogger(globalRegistry, os.Getenv, shortName, packageID, logger, opts...)
}

func applicationLogger(
	registry *registry,
	envGet func(string) string,
	shortName string,
	packageID string,
	logger **zap.Logger,
	opts ...LoggerOption,
) Tracer {
	loggerOptions := newLoggerOptions(shortName, opts...)
	tracer := register2(registry, shortName, packageID, logger, loggerOptions.registerOptions...)

	registry.factory = func(name string, level zapcore.Level) *zap.Logger {
		clonedOptions := loggerOptions
		if name != "" {
			clonedOptions.loggerName = name
		}

		// If the level was specified up-front, let's not use the one received
		if loggerOptions.level != nil {
			return newLogger(&clonedOptions)
		}

		clonedOptions.level = ptrLevel(zap.NewAtomicLevelAt(level))

		return newLogger(&clonedOptions)
	}

	// We must keep the pointer because it could be moved in the override below
	initialLogger := *logger

	logLevelSpec := newLogLevelSpec(envGet)
	registry.overrideFromSpec(logLevelSpec, registry.factory)

	appEntry := registry.entriesByPackageID[packageID]
	if *appEntry.logPtr != nil && *appEntry.logPtr == initialLogger {
		// No environment override the default logger, let's force INFO to be used in this case
		registry.setLoggerForEntry(appEntry, zapcore.InfoLevel, false, registry.factory)
	}

	appLogger := zap.NewNop()
	if *appEntry.logPtr != nil {
		appLogger = *appEntry.logPtr
	}

	// Hijack standard Golang `log` and redirect it to our common logger
	zap.RedirectStdLogAt(appLogger, zap.DebugLevel)

	if loggerOptions.switcherServerAutoStart != nil && *loggerOptions.switcherServerAutoStart {
		go func() {
			listenAddr := loggerOptions.switcherServerListenAddr
			appLogger.Debug("starting atomic level switcher", zap.String("listen_addr", listenAddr))

			handler := &switcherServerHandler{registry: registry}
			if err := http.ListenAndServe(listenAddr, handler); err != nil {
				appLogger.Warn("failed starting atomic level switcher", zap.Error(err), zap.String("listen_addr", listenAddr))
			}
		}()
	}

	return tracer
}

// NewLogger creates a new logger with sane defaults based on a varity of rules described
// below and automatically registered withing the logging registry.
func NewLogger(opts ...LoggerOption) *zap.Logger {
	logger, err := MaybeNewLogger(opts...)
	if err != nil {
		panic(fmt.Errorf("unable to create logger (in production? %t): %w", isProductionEnvironment(), err))
	}

	return logger
}

func MaybeNewLogger(opts ...LoggerOption) (*zap.Logger, error) {
	options := loggerOptions{}
	for _, opt := range opts {
		opt.apply(&options)
	}

	logger, err := maybeNewLogger(&options)
	if err != nil {
		return nil, err
	}

	return logger, nil
}

func newLogger(opts *loggerOptions) *zap.Logger {
	logger, err := maybeNewLogger(opts)
	if err != nil {
		panic(fmt.Errorf("unable to create logger (in production? %t): %w", isProductionEnvironment(), err))
	}

	return logger
}

func maybeNewLogger(opts *loggerOptions) (logger *zap.Logger, err error) {
	defer func() {
		if logger != nil && opts.loggerName != "" {
			logger = logger.Named(opts.loggerName)
		}
	}()

	zapOptions := opts.zapOptions

	if isProductionEnvironment() {
		reportAllErrors := opts.reportAllErrors != nil
		serviceName := opts.serviceName

		if reportAllErrors && opts.serviceName != nil {
			zapOptions = append(zapOptions, zapdriver.WrapCore(zapdriver.ReportAllErrors(true), zapdriver.ServiceName(*serviceName)))
		} else if reportAllErrors {
			zapOptions = append(zapOptions, zapdriver.WrapCore(zapdriver.ReportAllErrors(true)))
		} else if opts.serviceName != nil {
			zapOptions = append(zapOptions, zapdriver.WrapCore(zapdriver.ServiceName(*serviceName)))
		}

		return zapdriver.NewProductionConfig().Build(zapOptions...)
	}

	// Development logger
	isTTY := terminal.IsTerminal(int(os.Stderr.Fd()))
	logStdoutWriter := zapcore.Lock(os.Stderr)
	verbosity := 1
	if opts.encoderVerbosity != nil {
		verbosity = *opts.encoderVerbosity
	}

	return zap.New(zapcore.NewCore(NewEncoder(verbosity, isTTY), logStdoutWriter, opts.level), zapOptions...), nil
}

func isProductionEnvironment() bool {
	_, err := os.Stat("/.dockerenv")

	return !os.IsNotExist(err)
}

type Tracer interface {
	Enabled() bool
}

type boolTracer struct {
	value *bool
}

func (t boolTracer) Enabled() bool {
	if t.value == nil {
		return false
	}

	return *t.value
}

func ptrBool(value bool) *bool                        { return &value }
func ptrInt(value int) *int                           { return &value }
func ptrString(value string) *string                  { return &value }
func ptrLevel(value zap.AtomicLevel) *zap.AtomicLevel { return &value }
