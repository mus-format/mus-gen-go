// Package veropts provides options for versioned types.
package veropts

import (
	"errors"
	"reflect"
)

// Options specifies configuration options for versioned types.
type Options struct {
	Versions []Version
}

// Validate validates the Options struct.
func Validate(o *Options) error {
	foundCurrVersion := false
	for _, v := range o.Versions {
		if v.Migration == "" {
			if foundCurrVersion {
				return errors.New("multiple current versions found")
			}
			foundCurrVersion = true
		}
	}
	return nil
}

// Version represents a specific version of a type.
type Version struct {
	Type      reflect.Type
	Migration string
}

// SetOption defines a function type that applies a configuration option to an
// Options struct.
type SetOption func(o *Options)

// WithVersion returns a SetOption that adds an old type version with the
// specified migration function to the Options struct.
//
// The migration param must be a function name that accepts a type version and
// returns the target type.
//
// For example:
//
//	type Foo FooV2  // target type Foo
//	type FooV2        // current type version
//	type FooV1        // old type version
//
//	func MigrateFooV1(v FooV1) Foo { ... }
func WithVersion(t reflect.Type, migration string) SetOption {
	return func(o *Options) {
		o.Versions = append(o.Versions, Version{
			Type:      t,
			Migration: migration,
		})
	}
}

// WithCurrentVersion returns a SetOption that adds the current type version to
// the Options struct.
//
// For example:
//
//	type Foo FooV2 // target type Foo
//	type FooV2       // current type version
func WithCurrentVersion(t reflect.Type) SetOption {
	return func(o *Options) {
		o.Versions = append(o.Versions, Version{
			Type:      t,
			Migration: "",
		})
	}
}

// Apply applies a set of options to the provided Options struct.
func Apply(o *Options, opts ...SetOption) error {
	for _, opt := range opts {
		if opt != nil {
			opt(o)
		}
	}
	return Validate(o)
}
