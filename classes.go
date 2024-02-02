package libredis

import (
	"context"
	"time"
)

// ClassInfo ...
type ClassInfo struct {

	// the namespace name for the class
	Namespace string

	// the full-name for the class
	FullName string

	// aka: simple name
	Alias string
}

// Class ... Data Access Object
type Class interface {
	GetInfo() *ClassInfo

	SetMaxAge(maxAge time.Duration)
}

// TextClass ...
type TextClass interface {
	Class

	Clone() TextClass

	Put(ctx context.Context, id string, text string) error

	Get(ctx context.Context, id string) (string, error)

	GetOrNew(ctx context.Context, id string, template ...string) (string, error)
}

// BinaryClass ...
type BinaryClass interface {
	Class

	Clone() BinaryClass

	Put(ctx context.Context, id string, data []byte) error

	Get(ctx context.Context, id string) ([]byte, error)

	GetOrNew(ctx context.Context, id string, template []byte) ([]byte, error)
}

// JSONClass ...
type JSONClass interface {
	Class

	Clone() JSONClass

	Put(ctx context.Context, id string, o any) error

	Get(ctx context.Context, id string, o any) error

	GetOrNew(ctx context.Context, id string, o any) error
}

// ClassManager ...
type ClassManager interface {
	GetTextClass(namespace, class string) (TextClass, error)

	GetBinaryClass(namespace, class string) (BinaryClass, error)

	GetJSONClass(namespace, class string) (JSONClass, error)
}

// ClassRegistration ...
type ClassRegistration struct {
	ClassInfo

	Enabled bool

	// the source name of this class
	Source string

	MaxAge time.Duration
}

// ClassRegistry ...
type ClassRegistry interface {
	ListClasses() []*ClassRegistration
}
