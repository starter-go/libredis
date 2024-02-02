package libredis

import (
	"context"
	"time"
)

// Want 表示一个 redis 查询请求
type Want struct {
	// Context    context.Context
	Class      string
	ID         string
	Expiration time.Duration

	// Object     any // the data object

}

// Client 是面向业务的 redis 客户端
type Client interface {
	NewClient() Client

	// Context() context.Context

	Ping(ctx context.Context) error

	Exists(ctx context.Context, want *Want) (bool, error)
	Get(ctx context.Context, want *Want) (string, error)
	Put(ctx context.Context, want *Want, value string) error

	// GetText(want *Want) (string, error)
	// PutText(want *Want, text string) error
	// GetJSON(want *Want) error
	// PutJSON(want *Want) error

}
