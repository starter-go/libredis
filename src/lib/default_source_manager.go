package lib

import (
	"fmt"

	"github.com/starter-go/application"
	"github.com/starter-go/libredis"
	"github.com/starter-go/vlog"
)

// DefaultSourceManager ...
type DefaultSourceManager struct {

	//starter:component
	_as func(libredis.SourceManager) //starter:as("#")

	RegList []libredis.SourceRegistry //starter:inject(".")
	Factory libredis.SourceFactory    //starter:inject("#")

	sources []*libredis.SourceRegistration
}

func (inst *DefaultSourceManager) _impl() application.Lifecycle {
	inst._as(inst)
	return inst
}

func (inst *DefaultSourceManager) loadAll() ([]*libredis.SourceRegistration, error) {
	src := inst.RegList
	dst := inst.sources
	dst = nil
	for _, r1 := range src {
		list1 := r1.ListRegistrations()
		for _, r2 := range list1 {
			if r2 == nil {
				continue
			}
			err := inst.loadSource(r2)
			if err != nil {
				return nil, err
			}
			dst = append(dst, r2)
		}
	}
	return dst, nil
}

func (inst *DefaultSourceManager) loadSource(reg *libredis.SourceRegistration) error {

	if !reg.Enabled {
		return nil
	}

	cfg := reg.Config
	factory := reg.Factory
	src := reg.Source

	if factory == nil {
		factory = inst.Factory
		reg.Factory = factory
	}

	if src == nil {
		src1, err := factory.Open(cfg)
		if err != nil {
			return err
		}
		src = src1
		reg.Source = src1
	}

	client := src.Client(nil)
	return client.Ping()
}

func (inst *DefaultSourceManager) listAll() ([]*libredis.SourceRegistration, error) {
	all := inst.sources
	if all == nil {
		list, err := inst.loadAll()
		if err != nil {
			return nil, err
		}
		all = list
		inst.sources = all
	}
	return all, nil
}

// Life ...
func (inst *DefaultSourceManager) Life() *application.Life {
	return &application.Life{
		OnCreate: inst.init,
	}
}

func (inst *DefaultSourceManager) init() error {
	list, err := inst.listAll()
	if err != nil {
		return err
	}
	for _, item := range list {
		en := "no"
		if item.Enabled {
			en = "yes"
		}
		vlog.Info("redis source name:%s enabled:%s", item.Name, en)
	}
	return nil
}

// GetSource ...
func (inst *DefaultSourceManager) GetSource(name string) (libredis.Source, error) {
	list, err := inst.listAll()
	if err != nil {
		return nil, err
	}
	for _, item := range list {
		if item.Name == name {
			if item.Source != nil {
				return item.Source, nil
			}
		}
	}
	return nil, fmt.Errorf("no redis source with name:%s", name)
}
