package golang

import (
	"context"
	"fmt"
	"time"

	"github.com/starter-go/application"
	"github.com/starter-go/libredis"
)

// TestClasses ...
type TestClasses struct {

	//starter:component
	_as func(libredis.ClassRegistry) //starter:as(".")

	Ser libredis.Service //starter:inject("#")

	NS        string //starter:inject("${test.libredis.class.namespace}")
	ClassName string //starter:inject("${test.libredis.class.alias}")

}

func (inst *TestClasses) _impl() (libredis.ClassRegistry, application.Lifecycle) {
	return inst, inst
}

// ListClasses ...
func (inst *TestClasses) ListClasses() []*libredis.ClassRegistration {
	cr1 := &libredis.ClassRegistration{
		MaxAge:  time.Minute,
		Enabled: true,
		Source:  "master",
	}
	info := &cr1.ClassInfo
	info.Alias = inst.ClassName
	info.Namespace = inst.NS
	return []*libredis.ClassRegistration{cr1}
}

// Life ...
func (inst *TestClasses) Life() *application.Life {
	return &application.Life{
		OnLoop: inst.run,
	}
}

func (inst *TestClasses) run() error {

	ctx := context.Background()
	ns := inst.NS
	cname := inst.ClassName
	id := "666"
	value1 := "hello,libredis"

	dao, err := inst.Ser.GetTextClass(ns, cname)
	if err != nil {
		return err
	}

	value2, err := dao.GetOrNew(ctx, id, value1)
	if err != nil {
		return err
	}

	value3, err := dao.Get(ctx, id)
	if err != nil {
		return err
	}

	if value1 != value2 || value2 != value3 {
		const f = "bad test result values, want=[%s]  have1=[%s]  have2=[%s]"
		return fmt.Errorf(f, value1, value2, value3)
	}

	return nil
}
