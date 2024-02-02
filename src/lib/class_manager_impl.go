package lib

import (
	"fmt"
	"strings"
	"sync"

	"github.com/starter-go/libredis"
)

// ClassManagerImpl ...
type ClassManagerImpl struct {

	//starter:component
	_as func(libredis.ClassManager) //starter:as("#")

	ClassRegs     []libredis.ClassRegistry //starter:inject(".")
	SourceManager libredis.SourceManager   //starter:inject("#")

	table map[string]*libredis.ClassRegistration
	mutex sync.Mutex
}

func (inst *ClassManagerImpl) _impl() libredis.ClassManager {
	return inst
}

func (inst *ClassManagerImpl) keyForItem(item *libredis.ClassRegistration) string {
	ns := item.Namespace
	class := item.Alias
	return inst.keyForItemName(ns, class)
}

func (inst *ClassManagerImpl) keyForItemName(ns, class string) string {
	a := strings.TrimSpace(ns)
	b := strings.TrimSpace(class)
	return strings.ToLower(a + "#" + b)
}

func (inst *ClassManagerImpl) loadTable() (map[string]*libredis.ClassRegistration, error) {
	src := inst.ClassRegs
	dst := make(map[string]*libredis.ClassRegistration)
	for _, r1 := range src {
		r2s := r1.ListClasses()
		for _, r2 := range r2s {
			if r2.Enabled {
				key := inst.keyForItem(r2)
				old := dst[key]
				if old != nil {
					return nil, fmt.Errorf("duplicate libredis class [%s]", key)
				}
				r2.FullName = key
				dst[key] = r2
			}
		}
	}
	return dst, nil
}

func (inst *ClassManagerImpl) getTable() (map[string]*libredis.ClassRegistration, error) {
	tab := inst.table
	if tab != nil {
		return tab, nil
	}
	tab, err := inst.loadTable()
	if err == nil {
		inst.table = tab
		return tab, nil
	}
	return nil, err
}

func (inst *ClassManagerImpl) findClass(ns, class string) (*libredis.ClassRegistration, error) {
	inst.mutex.Lock()
	defer func() {
		inst.mutex.Unlock()
	}()
	tab, err := inst.getTable()
	if err != nil {
		return nil, err
	}
	key := inst.keyForItemName(ns, class)
	item := tab[key]
	if item == nil {
		const f = "cannot find item with key [%s] in libredis class table"
		return nil, fmt.Errorf(f, key)
	}
	return item, nil
}

func (inst *ClassManagerImpl) getClassContext(ns, class string) (*classDaoContext, error) {

	cr, err := inst.findClass(ns, class)
	if err != nil {
		return nil, err
	}

	sourceName := cr.Source
	source, err := inst.SourceManager.GetSource(sourceName)
	if err != nil {
		return nil, err
	}

	client := source.Client()

	ctx := &classDaoContext{}
	ctx.maxAge = cr.MaxAge
	ctx.source = source
	ctx.client = client

	ctx.info.Namespace = cr.Namespace
	ctx.info.Alias = cr.Alias
	ctx.info.FullName = cr.FullName

	return ctx, nil
}

// GetTextClass ...
func (inst *ClassManagerImpl) GetTextClass(ns, class string) (libredis.TextClass, error) {
	ctx, err := inst.getClassContext(ns, class)
	if err != nil {
		return nil, err
	}
	return &classDaoForText{context: ctx}, nil
}

// GetBinaryClass ...
func (inst *ClassManagerImpl) GetBinaryClass(ns, class string) (libredis.BinaryClass, error) {
	ctx, err := inst.getClassContext(ns, class)
	if err != nil {
		return nil, err
	}
	return &classDaoForBinary{context: ctx}, nil
}

// GetJSONClass ...
func (inst *ClassManagerImpl) GetJSONClass(ns, class string) (libredis.JSONClass, error) {
	ctx, err := inst.getClassContext(ns, class)
	if err != nil {
		return nil, err
	}
	return &classDaoForJSON{context: ctx}, nil
}
