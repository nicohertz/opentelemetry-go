// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package instrumentation // import "go.opentelemetry.io/otel/sdk/instrumentation"

import (
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// Scope represents the instrumentation scope.
type Scope struct {
	// Name is the name of the instrumentation scope. This should be the
	// Go package name of that scope.
	Name string
	// Version is the version of the instrumentation scope.
	Version string
	// SchemaURL of the telemetry emitted by the scope.
	SchemaURL string
	// Attributes of the telemetry emitted by the scope.
	Attributes attribute.Set
	// AutoProfiling indicates whether automatic runtime profiling is enabled.
	AutoProfiling bool
	// SkipProfiling indicates whether profiling is skipped by default at the
	// scope level.
	SkipProfiling bool

	spanStartOptions *[]trace.SpanStartOption
}

// SpanOptions returns the default span start options for this scope.
// The options are computed once on first call and then cached.
func (s *Scope) SpanOptions() []trace.SpanStartOption {
	if s.spanStartOptions == nil {
		var opts []trace.SpanStartOption
		if s.AutoProfiling {
			opts = append(opts, trace.ProfileRegion())
		}
		opts = append(opts, trace.WithSkipProfiling(s.SkipProfiling))
		s.spanStartOptions = &opts
	}
	return *s.spanStartOptions
}
