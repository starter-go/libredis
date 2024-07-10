package test4libredis
import (
    p21f95db42 "github.com/starter-go/keyvalues"
    p19ce63675 "github.com/starter-go/libredis/src/test/golang/cases"
     "github.com/starter-go/application"
)

// type p19ce63675.TestCrud in package:github.com/starter-go/libredis/src/test/golang/cases
//
// id:com-19ce6367568731c1-cases-TestCrud
// class:
// alias:
// scope:singleton
//
type p19ce636756_cases_TestCrud struct {
}

func (inst* p19ce636756_cases_TestCrud) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-19ce6367568731c1-cases-TestCrud"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p19ce636756_cases_TestCrud) new() any {
    return &p19ce63675.TestCrud{}
}

func (inst* p19ce636756_cases_TestCrud) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p19ce63675.TestCrud)
	nop(ie, com)

	
    com.Service = inst.getService(ie)
    com.NS = inst.getNS(ie)
    com.Alias = inst.getAlias(ie)


    return nil
}


func (inst*p19ce636756_cases_TestCrud) getService(ie application.InjectionExt)p21f95db42.Service{
    return ie.GetComponent("#alias-21f95db421796c61fc702c5dfd6515de-Service").(p21f95db42.Service)
}


func (inst*p19ce636756_cases_TestCrud) getNS(ie application.InjectionExt)string{
    return ie.GetString("${test.libredis.class.namespace}")
}


func (inst*p19ce636756_cases_TestCrud) getAlias(ie application.InjectionExt)string{
    return ie.GetString("${test.libredis.class.alias}")
}


