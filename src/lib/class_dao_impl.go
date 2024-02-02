package lib

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"time"

	"github.com/starter-go/libredis"
)

type classDaoContext struct {
	source libredis.Source
	client libredis.Client
	info   libredis.ClassInfo
	maxAge time.Duration
}

func (inst *classDaoContext) getInfo() *libredis.ClassInfo {
	dst := new(libredis.ClassInfo)
	*dst = inst.info
	return dst
}

func (inst *classDaoContext) makeWant(id string) *libredis.Want {
	want := new(libredis.Want)
	want.Class = inst.info.FullName
	want.ID = id
	want.Expiration = inst.maxAge
	return want
}

func (inst *classDaoContext) clone() *classDaoContext {
	dst := new(classDaoContext)
	*dst = *inst
	return dst
}

////////////////////////////////////////////////////////////////////////////////

type classDaoForText struct {
	context *classDaoContext
}

func (inst *classDaoForText) _impl() libredis.TextClass {
	return inst
}

func (inst *classDaoForText) Clone() libredis.TextClass {
	ctx := inst.context.clone()
	return &classDaoForText{context: ctx}
}

func (inst *classDaoForText) SetMaxAge(ma time.Duration) {
	inst.context.maxAge = ma
}

func (inst *classDaoForText) GetInfo() *libredis.ClassInfo {
	return inst.context.getInfo()
}

func (inst *classDaoForText) Put(ctx context.Context, id string, text string) error {
	client := inst.context.client
	want := inst.context.makeWant(id)
	return client.Put(ctx, want, text)
}

func (inst *classDaoForText) Get(ctx context.Context, id string) (string, error) {
	client := inst.context.client
	want := inst.context.makeWant(id)
	return client.Get(ctx, want)
}

func (inst *classDaoForText) GetOrNew(ctx context.Context, id string, template ...string) (string, error) {
	value, err := inst.Get(ctx, id)
	if err == nil {
		return value, nil // ok
	}
	value = ""
	if len(template) > 0 {
		value = template[0]
	}
	err = inst.Put(ctx, id, value)
	if err == nil {
		return value, nil // ok
	}
	return "", err
}

////////////////////////////////////////////////////////////////////////////////

type classDaoForBinary struct {
	context *classDaoContext
}

func (inst *classDaoForBinary) _impl() libredis.BinaryClass {
	return inst
}

func (inst *classDaoForBinary) Clone() libredis.BinaryClass {
	ctx := inst.context.clone()
	return &classDaoForBinary{context: ctx}
}

func (inst *classDaoForBinary) SetMaxAge(ma time.Duration) {
	inst.context.maxAge = ma
}

func (inst *classDaoForBinary) GetInfo() *libredis.ClassInfo {
	return inst.context.getInfo()
}

func (inst *classDaoForBinary) Put(ctx context.Context, id string, data []byte) error {
	text := base64.StdEncoding.EncodeToString(data)
	client := inst.context.client
	want := inst.context.makeWant(id)
	return client.Put(ctx, want, text)
}

func (inst *classDaoForBinary) Get(ctx context.Context, id string) ([]byte, error) {
	client := inst.context.client
	want := inst.context.makeWant(id)
	str, err := client.Get(ctx, want)
	if err != nil {
		return nil, err
	}
	return base64.StdEncoding.DecodeString(str)
}

func (inst *classDaoForBinary) GetOrNew(ctx context.Context, id string, template []byte) ([]byte, error) {
	// get
	data, err := inst.Get(ctx, id)
	if err == nil {
		return data, nil // ok
	}
	// prepare template
	if template == nil {
		template = make([]byte, 0)
	}
	data = template
	// put
	err = inst.Put(ctx, id, template)
	if err == nil {
		return data, nil // ok
	}
	return nil, err
}

////////////////////////////////////////////////////////////////////////////////

type classDaoForJSON struct {
	context *classDaoContext
}

func (inst *classDaoForJSON) _impl() libredis.JSONClass {
	return inst
}

func (inst *classDaoForJSON) Clone() libredis.JSONClass {
	ctx := inst.context.clone()
	return &classDaoForJSON{context: ctx}
}

func (inst *classDaoForJSON) SetMaxAge(ma time.Duration) {
	inst.context.maxAge = ma
}

func (inst *classDaoForJSON) GetInfo() *libredis.ClassInfo {
	return inst.context.getInfo()
}

func (inst *classDaoForJSON) Put(ctx context.Context, id string, o any) error {
	data, err := json.Marshal(o)
	if err != nil {
		return err
	}
	client := inst.context.client
	want := inst.context.makeWant(id)
	return client.Put(ctx, want, string(data))
}

func (inst *classDaoForJSON) Get(ctx context.Context, id string, o any) error {
	client := inst.context.client
	want := inst.context.makeWant(id)
	str, err := client.Get(ctx, want)
	if err != nil {
		return err
	}
	data := []byte(str)
	return json.Unmarshal(data, o)
}

func (inst *classDaoForJSON) GetOrNew(ctx context.Context, id string, o any) error {
	err := inst.Get(ctx, id, o)
	if err == nil {
		return nil // ok
	}
	return inst.Put(ctx, id, o)
}

////////////////////////////////////////////////////////////////////////////////
