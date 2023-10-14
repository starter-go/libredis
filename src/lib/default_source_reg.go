package lib

import (
	"strings"

	"github.com/starter-go/application"
	"github.com/starter-go/application/properties"
	"github.com/starter-go/libredis"
	"github.com/starter-go/vlog"
)

const (
	datasourcePropertyNamePrefix = "redis-source."
	datasourcePropertyNameSuffix = ".host"
)

// DefaultSourceRegistry ...
type DefaultSourceRegistry struct {

	//starter:component
	_as func(libredis.SourceRegistry) //starter:as(".")

	AC       application.Context    //starter:inject("context")
	Factory  libredis.SourceFactory //starter:inject("#")
	NameList string                 //starter:inject("${redis.sources}")

	namelistAll     []string
	namelistWant    []string
	namelistEnabled []string
	all             []*libredis.SourceRegistration
}

func (inst *DefaultSourceRegistry) _impl() {
	inst._as(inst)
}

// ListRegistrations ...
func (inst *DefaultSourceRegistry) ListRegistrations() []*libredis.SourceRegistration {

	props := inst.AC.GetProperties()
	src := inst.listSourceNames(props)
	dst := make([]*libredis.SourceRegistration, 0)
	nlistEnabled := make([]string, 0)

	for _, name := range src {
		cfg, err := inst.getSourceConfig(name, props)
		if err != nil {
			vlog.Warn("bad redis source configuration: %s", err.Error())
			continue
		}
		if cfg.Enabled {
			nlistEnabled = append(nlistEnabled, name)
		}
		reg := &libredis.SourceRegistration{}
		reg.Name = name
		reg.Config = cfg
		reg.Enabled = cfg.Enabled
		reg.Factory = inst.Factory
		dst = append(dst, reg)
	}

	inst.namelistAll = src
	inst.namelistEnabled = nlistEnabled
	inst.namelistWant = inst.loadWantNameList()

	return dst
}

func (inst *DefaultSourceRegistry) loadWantNameList() []string {
	src := strings.Split(inst.NameList, ",")
	dst := make([]string, 0)
	for _, name := range src {
		name = strings.TrimSpace(name)
		if name == "" {
			continue
		}
		dst = append(dst, name)
	}
	return dst
}

func (inst *DefaultSourceRegistry) getSourceConfig(name string, p properties.Table) (*libredis.Configuration, error) {
	const (
		prefix = datasourcePropertyNamePrefix
	)
	keyPrefix := prefix + name + "."
	cfg := &libredis.Configuration{}
	getter := p.Getter().Required()

	cfg.Name = name
	cfg.Enabled = getter.GetBool(keyPrefix + "enabled")
	cfg.Host = getter.GetString(keyPrefix + "host")
	cfg.Port = getter.GetInt(keyPrefix + "port")
	cfg.Username = getter.GetString(keyPrefix + "username")
	cfg.Password = getter.GetString(keyPrefix + "password")
	cfg.DB = getter.GetInt(keyPrefix + "db")

	err := getter.Error()
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func (inst *DefaultSourceRegistry) listSourceNames(p properties.Table) []string {
	const (
		prefix = datasourcePropertyNamePrefix
		suffix = datasourcePropertyNameSuffix
	)
	src := p.Names()
	dst := make([]string, 0)
	for _, key := range src {
		if strings.HasPrefix(key, prefix) && strings.HasSuffix(key, suffix) {
			i1 := len(prefix)
			i2 := len(key) - len(suffix)
			name := key[i1:i2]
			dst = append(dst, name)
		}
	}
	return dst
}
