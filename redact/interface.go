// Package redact provides interfaces and methods to help implement redaction.
package redact

import "context"

// Redactor provides the method to be used to Redact
type Redactor interface {
	Redact()
}

// Apply will apply redaction on the input, if it implements Redactor.
// It will do nothing if the object does not implement the interface.
func Apply(in interface{}) {
	if red, ok := in.(Redactor); ok {
		red.Redact()
	}
}

// Bypass provides a way to bypass the internal marked methods to be ran
// through clients
type Bypass interface {
	CheckInternal(ctx context.Context, in interface{}) bool
}

// Wrap: helps to implement Bypass
type Wrap func(ctx context.Context, in interface{}) bool

func (w Wrap) CheckInternal(ctx context.Context, in interface{}) bool { return w(ctx, in) }
