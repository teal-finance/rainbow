// Copyright 2019 dfuse Platform Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logging

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type registerConfig struct {
	shortName      string
	isTraceEnabled *bool
	onUpdate       func(newLogger *zap.Logger)
}

// RegisterOption are option parameters that you can set when registering a new logger
// in the system using `Register` function.
type RegisterOption interface {
	apply(config *registerConfig)
}

type registerOptionFunc func(config *registerConfig)

func (f registerOptionFunc) apply(config *registerConfig) {
	f(config)
}

// RegisterOnUpdate enable you to have a hook function that will receive the new logger
// that is going to be assigned to your logger instance. This is useful in some situation
// where you need to update other instances or re-configuring a bit the logger when
// a new one is attached.
//
// This is called **after** the instance has been re-assigned.
func RegisterOnUpdate(onUpdate func(newLogger *zap.Logger)) RegisterOption {
	return registerOptionFunc(func(config *registerConfig) {
		config.onUpdate = onUpdate
	})
}

func registerShortName(shortName string) RegisterOption {
	return registerOptionFunc(func(config *registerConfig) {
		config.shortName = shortName
	})
}

func registerWithTracer(isEnabled *bool) RegisterOption {
	return registerOptionFunc(func(config *registerConfig) {
		config.isTraceEnabled = isEnabled
	})
}

type LoggerExtender func(*zap.Logger) *zap.Logger

type registryEntry struct {
	packageID    string
	shortName    string
	traceEnabled *bool
	logPtr       **zap.Logger
	onUpdate     func(newLogger *zap.Logger)
}

var globalRegistry = newRegistry()
var defaultLogger = zap.NewNop()

func Register(packageID string, zlogPtr **zap.Logger, options ...RegisterOption) {
	register(globalRegistry, packageID, zlogPtr, options...)
}

func register2(registry *registry, shortName string, packageID string, zlogPtr **zap.Logger, options ...RegisterOption) Tracer {
	tracer := boolTracer{new(bool)}

	allOptions := append([]RegisterOption{
		registerShortName(shortName),
		registerWithTracer(tracer.value),
	}, options...)

	register(registry, packageID, zlogPtr, allOptions...)
	return tracer
}

func register(registry *registry, packageID string, zlogPtr **zap.Logger, options ...RegisterOption) {
	if zlogPtr == nil {
		panic("the zlog pointer (of type **zap.Logger) must be set")
	}

	config := registerConfig{}
	for _, opt := range options {
		opt.apply(&config)
	}

	entry := &registryEntry{
		packageID:    packageID,
		shortName:    config.shortName,
		traceEnabled: config.isTraceEnabled,
		logPtr:       zlogPtr,
		onUpdate:     config.onUpdate,
	}

	registry.addEntry(entry)

	logger := defaultLogger
	if *zlogPtr != nil {
		logger = *zlogPtr
	}

	// The tracing has already been set, so we can go unspecified here to not change anything
	setLogger(entry, logger, unspecifiedTracing)
}

func Set(logger *zap.Logger, regexps ...string) {
	for name, entry := range globalRegistry.entriesByPackageID {
		if len(regexps) == 0 {
			setLogger(entry, logger, unspecifiedTracing)
		} else {
			for _, re := range regexps {
				regex, err := regexp.Compile(re)
				if (err == nil && regex.MatchString(name)) || (err != nil && name == re) {
					setLogger(entry, logger, unspecifiedTracing)
				}
			}
		}
	}
}

// Extend is different than `Set` by being able to re-configure the existing logger set for
// all registered logger in the registry. This is useful for example to add a field to the
// currently set logger:
//
// ```
// logger.Extend(func (current *zap.Logger) { return current.With("name", "value") }, "github.com/dfuse-io/app.*")
// ```
func Extend(extender LoggerExtender, regexps ...string) {
	extend(extender, unspecifiedTracing, regexps...)
}

func extend(extender LoggerExtender, tracing tracingType, regexps ...string) {
	for name, entry := range globalRegistry.entriesByPackageID {
		if *entry.logPtr == nil {
			continue
		}

		if len(regexps) == 0 {
			setLogger(entry, extender(*entry.logPtr), tracing)
		} else {
			for _, re := range regexps {
				if regexp.MustCompile(re).MatchString(name) {
					setLogger(entry, extender(*entry.logPtr), tracing)
				}
			}
		}
	}
}

// Override sets the given logger on previously registered and next
// registrations.  Useful in tests.
func Override(logger *zap.Logger) {
	defaultLogger = logger
	Set(logger)
}

// TestingOverride calls `Override` (or `Set`, see below) with a development
// logger setup correctly with the right level based on some environment variables.
//
// By default, override using a `zap.NewDevelopment` logger (`info`), if
// environment variable `DEBUG` is set to anything or environment variable `TRACE`
// is set to `true`, logger is set in `debug` level.
//
// If `DEBUG` is set to something else than `true` and/or if `TRACE` is set
// to something else than
func TestingOverride() {
	debug := os.Getenv("DEBUG")
	trace := os.Getenv("TRACE")
	if debug == "" && trace == "" {
		return
	}

	logger, _ := zap.NewDevelopment()

	regex := ""
	if debug != "true" {
		regex = debug
	}

	if regex == "" && trace != "true" {
		regex = trace
	}

	if regex == "" {
		Override(logger)
	} else {
		for _, regexPart := range strings.Split(regex, ",") {
			regexPart = strings.TrimSpace(regexPart)
			if regexPart != "" {
				Set(logger, regexPart)
			}
		}
	}
}

type tracingType uint8

const (
	unspecifiedTracing tracingType = iota
	enableTracing
	disableTracing
)

func setLogger(entry *registryEntry, logger *zap.Logger, tracing tracingType) {
	if entry == nil || logger == nil {
		return
	}

	*entry.logPtr = logger
	if entry.traceEnabled != nil && tracing != unspecifiedTracing {
		switch tracing {
		case enableTracing:
			*entry.traceEnabled = true
		case disableTracing:
			*entry.traceEnabled = false
		}
	}

	if entry.onUpdate != nil {
		entry.onUpdate(logger)
	}
}

type registry struct {
	sync.RWMutex

	factory            loggerFactory
	entriesByPackageID map[string]*registryEntry
	entriesByShortName map[string][]*registryEntry
}

func newRegistry() *registry {
	return &registry{
		entriesByPackageID: make(map[string]*registryEntry),
		entriesByShortName: make(map[string][]*registryEntry),
		factory: func(name string, level zapcore.Level) *zap.Logger {
			loggerOptions := newLoggerOptions("", WithAtomicLevel(zap.NewAtomicLevelAt(level)))
			if name != "" {
				loggerOptions.loggerName = name
			}

			return newLogger(&loggerOptions)
		},
	}
}

func (r *registry) addEntry(entry *registryEntry) {
	if entry == nil {
		panic("refusing to add a nil registry entry")
	}

	id := validateEntryIdentifier("package ID", entry.packageID, false)
	shortName := validateEntryIdentifier("short name", entry.shortName, true)

	if actual := r.entriesByPackageID[id]; actual != nil {
		panic(fmt.Sprintf("packageID %q is already registered", id))
	}

	entry.packageID = id
	entry.shortName = shortName

	r.entriesByPackageID[id] = entry
	if shortName != "" {
		r.entriesByShortName[shortName] = append(r.entriesByShortName[shortName], entry)
	}
}

func (r *registry) overrideFromSpec(spec *logLevelSpec, factory loggerFactory) {
	for _, specForKey := range spec.sortedSpecs() {
		if specForKey.key == "true" || specForKey.key == "*" {
			for _, entry := range r.entriesByPackageID {
				r.setLoggerForEntry(entry, specForKey.level, specForKey.trace, factory)
			}

			continue
		}

		r.setLoggerFromSpec(specForKey, factory)
	}
}

func (r *registry) setLoggerFromSpec(spec *levelSpec, factory loggerFactory) {
	entries, found := r.entriesByShortName[spec.key]
	if found {
		for _, entry := range entries {
			r.setLoggerForEntry(entry, spec.level, spec.trace, factory)
		}
		return
	}

	entry, found := r.entriesByPackageID[spec.key]
	if found {
		r.setLoggerForEntry(entry, spec.level, spec.trace, factory)
		return
	}

	regex, err := regexp.Compile(spec.key)
	for packageID, entry := range globalRegistry.entriesByPackageID {
		if (err == nil && regex.MatchString(packageID)) || (err != nil && packageID == spec.key) {
			r.setLoggerForEntry(entry, spec.level, spec.trace, factory)
		}
	}
}

func (r *registry) setLoggerForEntry(entry *registryEntry, level zapcore.Level, trace bool, factory loggerFactory) {
	if entry == nil {
		return
	}

	logger := factory(entry.shortName, level)

	*entry.logPtr = logger

	// It's possible for an entry to have no tracer registered, for example if the legacy
	// register method is used. We must protect from this and not set anything.
	if entry.traceEnabled != nil {
		*entry.traceEnabled = trace
	}

	if entry.onUpdate != nil {
		entry.onUpdate(logger)
	}
}

func validateEntryIdentifier(tag string, rawInput string, allowEmpty bool) string {
	input := strings.TrimSpace(rawInput)
	if input == "" && !allowEmpty {
		panic(fmt.Errorf("the %s %q is invalid, must not be empty", tag, input))
	}

	if input == "true" {
		panic(fmt.Errorf("the %s %q is invalid, the identifier 'true' is reserved", tag, input))
	}

	if input == "*" {
		panic(fmt.Errorf("the %s %q is invalid, the identifier '*' is reserved", tag, input))
	}

	if strings.HasPrefix(input, "-") {
		panic(fmt.Errorf("the %s %q is invalid, must not starts with the '-' character", tag, input))
	}

	if strings.Contains(input, ",") {
		panic(fmt.Errorf("the %s %q is invalid, must not contain the ',' character", tag, input))
	}

	if strings.Contains(input, "=") {
		panic(fmt.Errorf("the %s %q is invalid, must not contain the '=' character", tag, input))
	}

	return input
}
