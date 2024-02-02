package lib

import (
	"context"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/starter-go/libredis"
)

// DefaultSourceFactory ...
type DefaultSourceFactory struct {

	//starter:component
	_as func(libredis.SourceFactory) //starter:as("#")

	MaxAgeInMS int64 //starter:inject("${redis.default-max-age}")

}

func (inst *DefaultSourceFactory) _impl() {
	inst._as(inst)
}

func (inst *DefaultSourceFactory) getMaxAge() time.Duration {
	ms := inst.MaxAgeInMS
	if ms < 1 {
		ms = 1000
	}
	return time.Millisecond * time.Duration(ms)
}

// Open ...
func (inst *DefaultSourceFactory) Open(cfg *libredis.Configuration) (libredis.Source, error) {

	host := cfg.Host
	port := cfg.Port
	user := cfg.Username
	pass := cfg.Password
	db := cfg.DB

	// max-age
	maxAge := inst.getMaxAge()

	// port
	const defaultPort = 6379
	if port < 1 {
		port = defaultPort
	}

	// new client
	addr := host + ":" + strconv.Itoa(port)
	rawClient := redis.NewClient(&redis.Options{
		Addr:     addr, // "localhost:6379",
		Password: pass, // no password set
		Username: user, // user
		DB:       db,   // use default DB
	})

	// new source
	s := &sourceImpl{
		client:        rawClient,
		defaultMaxAge: maxAge,
	}
	ctx := context.Background()
	c := s.Client()
	err := c.Ping(ctx)
	if err != nil {
		return nil, err
	}
	return s, nil
}
