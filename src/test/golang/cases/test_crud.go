package cases

import (
	"fmt"

	"github.com/starter-go/application"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/keyvalues"
)

// TestCrud ...
type TestCrud struct {

	//starter:component

	// _as func(libredis.ClassRegistry) //starter:as(".")

	Service keyvalues.Service //starter:inject("#")

	NS    string //starter:inject("${test.libredis.class.namespace}")
	Alias string //starter:inject("${test.libredis.class.alias}")

}

// func (inst *TestCrud) _impl() (libredis.ClassRegistry, application.Lifecycle) {
// 	return inst, inst
// }

// Life  ...
func (inst *TestCrud) Life() *application.Life {
	return &application.Life{
		OnLoop: inst.run,
	}
}

func (inst *TestCrud) run() error {

	now := lang.Now()
	ser := inst.Service
	ns := inst.NS
	alias := inst.Alias
	id := "1234567"
	value1 := now.String()

	cl, err := ser.GetClassNS(keyvalues.NS(ns), keyvalues.ClassAlias(alias))
	if err != nil {
		return err
	}

	ent := cl.GetTextEntry(id)
	err = ent.Put(value1, nil)
	if err != nil {
		return err
	}

	value2, err := ent.Get()
	if err != nil {
		return err
	}

	if value1 != value2 {
		return fmt.Errorf("value1 != value2")
	}
	return nil
}
