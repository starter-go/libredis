package lib4libredis
import (
    p586008aca "github.com/starter-go/libredis/src/main/golang/driver"
     "github.com/starter-go/application"
)

// type p586008aca.RedisDriver in package:github.com/starter-go/libredis/src/main/golang/driver
//
// id:com-586008acad83d11b-driver-RedisDriver
// class:class-21f95db421796c61fc702c5dfd6515de-Registry
// alias:
// scope:singleton
//
type p586008acad_driver_RedisDriver struct {
}

func (inst* p586008acad_driver_RedisDriver) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-586008acad83d11b-driver-RedisDriver"
	r.Classes = "class-21f95db421796c61fc702c5dfd6515de-Registry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p586008acad_driver_RedisDriver) new() any {
    return &p586008aca.RedisDriver{}
}

func (inst* p586008acad_driver_RedisDriver) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p586008aca.RedisDriver)
	nop(ie, com)

	


    return nil
}


