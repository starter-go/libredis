package gen4test
import (
    p91658cc84 "github.com/starter-go/libredis"
    pe547391eb "github.com/starter-go/libredis/src/test/golang"
     "github.com/starter-go/application"
)

// type pe547391eb.Tester in package:github.com/starter-go/libredis/src/test/golang
//
// id:com-e547391eb9a4a6b9-golang-Tester
// class:
// alias:
// scope:singleton
//
type pe547391eb9_golang_Tester struct {
}

func (inst* pe547391eb9_golang_Tester) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-e547391eb9a4a6b9-golang-Tester"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pe547391eb9_golang_Tester) new() any {
    return &pe547391eb.Tester{}
}

func (inst* pe547391eb9_golang_Tester) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pe547391eb.Tester)
	nop(ie, com)

	
    com.SourceManager = inst.getSourceManager(ie)


    return nil
}


func (inst*pe547391eb9_golang_Tester) getSourceManager(ie application.InjectionExt)p91658cc84.SourceManager{
    return ie.GetComponent("#alias-91658cc84667f4d073289b7614060648-SourceManager").(p91658cc84.SourceManager)
}


