package driver

import (
	"context"
	"crypto/tls"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/starter-go/keyvalues"
	"github.com/starter-go/vlog"
)

// RedisDriver ...
type RedisDriver struct {

	//starter:component

	_as func(keyvalues.Registry) //starter:as(".")
}

func (inst *RedisDriver) _impl() (keyvalues.Registry, keyvalues.Driver) {
	return inst, inst
}

// ListRegistrations ...
func (inst *RedisDriver) ListRegistrations() []*keyvalues.Registration {
	info := &keyvalues.DriverInfo{
		Name:   "redis",
		Driver: inst,
	}
	r1 := &keyvalues.Registration{
		Enabled:  true,
		Priority: 0,
		Driver:   info,
	}
	return []*keyvalues.Registration{r1}
}

// Open ...
func (inst *RedisDriver) Open(cfg *keyvalues.Configuration) (keyvalues.Store, error) {

	store := &myRedisStore{
		config: cfg,
	}

	err := store.open()
	if err != nil {
		return nil, err
	}

	return store, nil
}

////////////////////////////////////////////////////////////////////////////////

type myRedisStore struct {
	config *keyvalues.Configuration
	client *redis.Client
}

func (inst *myRedisStore) _impl() keyvalues.Store {
	return inst
}

func (inst *myRedisStore) open() error {

	const (
		defaultPort = 6379
	)

	cfg := inst.config
	host := cfg.Host
	port := cfg.Port

	if port < 1 {
		port = defaultPort
	}

	tlscfg := &tls.Config{
		MinVersion: tls.VersionTLS12,
		ServerName: "your domain",
	}

	opt := &redis.Options{}
	opt.TLSConfig = tlscfg
	opt.Addr = host + ":" + strconv.Itoa(port)

	opt.TLSConfig = nil

	db := redis.NewClient(opt)
	err := inst.checkConn(db)
	if err != nil {
		vlog.Warn(err.Error())
		return err
	}
	inst.client = db
	return nil
}

func (inst *myRedisStore) makeContext() (context.Context, func()) {
	now := time.Now()
	ctx := context.Background()
	timeout := time.Second * 30
	ctx, cancel := context.WithDeadline(ctx, now.Add(timeout))
	return ctx, cancel
}

func (inst *myRedisStore) checkConn(db *redis.Client) error {

	ctx, cancel := inst.makeContext()
	defer cancel()

	key := "test/demo1"
	maxage := time.Second * 10
	value1 := "hello,redis"

	status1 := db.Set(ctx, key, value1, maxage)
	err := status1.Err()
	if err != nil {
		return err
	}

	status2 := db.Get(ctx, key)
	err = status2.Err()
	if err != nil {
		return err
	}

	value2 := status2.Val()
	if value1 != value2 {
		return fmt.Errorf("bad redis client connection")
	}
	return nil
}

func (inst *myRedisStore) Put(key string, value []byte, opt *keyvalues.Options) error {

	ctx, cancel := inst.makeContext()
	defer cancel()

	if opt == nil {
		opt = &keyvalues.Options{
			MaxAge: time.Second * 60,
		}
	}

	db := inst.client
	status := db.Set(ctx, key, value, opt.MaxAge)
	return status.Err()
}

func (inst *myRedisStore) Get(key string) ([]byte, error) {

	ctx, cancel := inst.makeContext()
	defer cancel()

	db := inst.client
	status := db.Get(ctx, key)
	err := status.Err()
	if err != nil {
		return nil, err
	}
	return status.Bytes()
}

func (inst *myRedisStore) Contains(key string) (bool, error) {

	ctx, cancel := inst.makeContext()
	defer cancel()

	db := inst.client
	status := db.Get(ctx, key)
	err := status.Err()
	if err == nil {
		b, err := status.Bytes()
		if err == nil && b != nil {
			return true, nil
		}
	}
	return false, nil
}
