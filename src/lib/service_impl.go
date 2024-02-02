package lib

import "github.com/starter-go/libredis"

// ServiceImpl ...
type ServiceImpl struct {

	//starter:component
	_as func(libredis.Service) //starter:as("#")

	ClassMan  libredis.ClassManager  //starter:inject("#")
	SourceMan libredis.SourceManager //starter:inject("#")

}

func (inst *ServiceImpl) _impl() libredis.Service {
	return inst
}

// GetClassManager ...
func (inst *ServiceImpl) GetClassManager() libredis.ClassManager {
	return inst.ClassMan
}

// GetSourceManager ...
func (inst *ServiceImpl) GetSourceManager() libredis.SourceManager {
	return inst.SourceMan
}

// GetTextClass ...
func (inst *ServiceImpl) GetTextClass(ns, alias string) (libredis.TextClass, error) {
	return inst.ClassMan.GetTextClass(ns, alias)
}

// GetBinaryClass ...
func (inst *ServiceImpl) GetBinaryClass(ns, alias string) (libredis.BinaryClass, error) {
	return inst.ClassMan.GetBinaryClass(ns, alias)
}

// GetJSONClass ...
func (inst *ServiceImpl) GetJSONClass(ns, alias string) (libredis.JSONClass, error) {
	return inst.ClassMan.GetJSONClass(ns, alias)
}
