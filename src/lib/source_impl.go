package lib

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/libredis"
	"github.com/starter-go/vlog"
)

type sourceImpl struct {
	client        *redis.Client
	defaultMaxAge time.Duration // 默认的 max-age
}

func (inst *sourceImpl) _impl(a libredis.Source) {
	a = inst
}

func (inst *sourceImpl) Client() libredis.Client {
	return &clientImpl{
		source: inst,
		client: inst.client,
	}
}

////////////////////////////////////////////////////////////////////////////////

type clientImpl struct {
	// ctx    context.Context
	client *redis.Client
	source *sourceImpl
}

func (inst *clientImpl) _impl(a libredis.Client) {
	a = inst
}

func (inst *clientImpl) NewClient() libredis.Client {
	return inst.source.Client()
}

// func (inst *clientImpl) Context() context.Context {
// 	return inst.ctx
// }

func (inst *clientImpl) keyFor(want *libredis.Want) string {
	cl := want.Class
	id := want.ID
	return "obj:" + cl + "/" + id
}

func (inst *clientImpl) Ping(ctx context.Context) error {
	now := lang.Now()
	str1 := now.String()
	want := &libredis.Want{
		Class: "test",
		ID:    "ping",
	}
	err := inst.Put(ctx, want, str1)
	if err != nil {
		return err
	}
	str2, err := inst.Get(ctx, want)
	if err != nil {
		return err
	}
	vlog.Debug("[OK] ping redis with data1:%s data2:%s", str1, str2)
	return nil
}

func (inst *clientImpl) Exists(ctx context.Context, want *libredis.Want) (bool, error) {
	// ctx := inst.ctx
	key := inst.keyFor(want)
	result, err := inst.client.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return result == 1, nil
}

// get as text
func (inst *clientImpl) Get(ctx context.Context, want *libredis.Want) (string, error) {
	// ctx := inst.ctx
	key := inst.keyFor(want)
	return inst.client.Get(ctx, key).Result()
}

// put as text
func (inst *clientImpl) Put(ctx context.Context, want *libredis.Want, text string) error {
	// ctx := inst.ctx
	key := inst.keyFor(want)
	exp := want.Expiration
	if exp < 1 {
		exp = inst.source.defaultMaxAge
	}
	_, err := inst.client.Set(ctx, key, text, exp).Result()
	return err
}

// func (inst *clientImpl) GetJSON(want *libredis.Want) error {
// 	o := want.Object
// 	if o == nil {
// 		return fmt.Errorf("no model to get object from redis, class:" + want.Class)
// 	}
// 	str, err := inst.GetText(want)
// 	if err != nil {
// 		return err
// 	}
// 	return json.Unmarshal([]byte(str), o)
// }

// func (inst *clientImpl) PutJSON(want *libredis.Want) error {
// 	str := ""
// 	o := want.Object
// 	if o != nil {
// 		data, err := json.Marshal(o)
// 		if err != nil {
// 			return err
// 		}
// 		str = string(data)
// 	}
// 	return inst.PutText(want, str)
// }
