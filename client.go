package libredis

import (
	"context"
	"time"
)

// Want 表示一个 redis 查询请求
type Want struct {
	Class      string
	ID         string
	Object     any // the data object
	Expiration time.Duration
}

// Client 是面向业务的 redis 客户端
type Client interface {
	NewClient(c context.Context) Client

	Context() context.Context

	Ping() error

	Exists(want *Want) (bool, error)

	GetText(want *Want) (string, error)

	PutText(want *Want, text string) error

	GetJSON(want *Want) error

	PutJSON(want *Want) error
}
