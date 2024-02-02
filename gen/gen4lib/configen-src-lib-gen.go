package gen4lib
import (
    p0ef6f2938 "github.com/starter-go/application"
    p91658cc84 "github.com/starter-go/libredis"
    p3c2ff7f68 "github.com/starter-go/libredis/src/lib"
     "github.com/starter-go/application"
)

// type p3c2ff7f68.ClassManagerImpl in package:github.com/starter-go/libredis/src/lib
//
// id:com-3c2ff7f68e8b3d61-lib-ClassManagerImpl
// class:
// alias:alias-91658cc84667f4d073289b7614060648-ClassManager
// scope:singleton
//
type p3c2ff7f68e_lib_ClassManagerImpl struct {
}

func (inst* p3c2ff7f68e_lib_ClassManagerImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-3c2ff7f68e8b3d61-lib-ClassManagerImpl"
	r.Classes = ""
	r.Aliases = "alias-91658cc84667f4d073289b7614060648-ClassManager"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p3c2ff7f68e_lib_ClassManagerImpl) new() any {
    return &p3c2ff7f68.ClassManagerImpl{}
}

func (inst* p3c2ff7f68e_lib_ClassManagerImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p3c2ff7f68.ClassManagerImpl)
	nop(ie, com)

	
    com.ClassRegs = inst.getClassRegs(ie)
    com.SourceManager = inst.getSourceManager(ie)


    return nil
}


func (inst*p3c2ff7f68e_lib_ClassManagerImpl) getClassRegs(ie application.InjectionExt)[]p91658cc84.ClassRegistry{
    dst := make([]p91658cc84.ClassRegistry, 0)
    src := ie.ListComponents(".class-91658cc84667f4d073289b7614060648-ClassRegistry")
    for _, item1 := range src {
        item2 := item1.(p91658cc84.ClassRegistry)
        dst = append(dst, item2)
    }
    return dst
}


func (inst*p3c2ff7f68e_lib_ClassManagerImpl) getSourceManager(ie application.InjectionExt)p91658cc84.SourceManager{
    return ie.GetComponent("#alias-91658cc84667f4d073289b7614060648-SourceManager").(p91658cc84.SourceManager)
}



// type p3c2ff7f68.DefaultSourceFactory in package:github.com/starter-go/libredis/src/lib
//
// id:com-3c2ff7f68e8b3d61-lib-DefaultSourceFactory
// class:
// alias:alias-91658cc84667f4d073289b7614060648-SourceFactory
// scope:singleton
//
type p3c2ff7f68e_lib_DefaultSourceFactory struct {
}

func (inst* p3c2ff7f68e_lib_DefaultSourceFactory) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-3c2ff7f68e8b3d61-lib-DefaultSourceFactory"
	r.Classes = ""
	r.Aliases = "alias-91658cc84667f4d073289b7614060648-SourceFactory"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p3c2ff7f68e_lib_DefaultSourceFactory) new() any {
    return &p3c2ff7f68.DefaultSourceFactory{}
}

func (inst* p3c2ff7f68e_lib_DefaultSourceFactory) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p3c2ff7f68.DefaultSourceFactory)
	nop(ie, com)

	
    com.MaxAgeInMS = inst.getMaxAgeInMS(ie)


    return nil
}


func (inst*p3c2ff7f68e_lib_DefaultSourceFactory) getMaxAgeInMS(ie application.InjectionExt)int64{
    return ie.GetInt64("${redis.default-max-age}")
}



// type p3c2ff7f68.DefaultSourceManager in package:github.com/starter-go/libredis/src/lib
//
// id:com-3c2ff7f68e8b3d61-lib-DefaultSourceManager
// class:
// alias:alias-91658cc84667f4d073289b7614060648-SourceManager
// scope:singleton
//
type p3c2ff7f68e_lib_DefaultSourceManager struct {
}

func (inst* p3c2ff7f68e_lib_DefaultSourceManager) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-3c2ff7f68e8b3d61-lib-DefaultSourceManager"
	r.Classes = ""
	r.Aliases = "alias-91658cc84667f4d073289b7614060648-SourceManager"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p3c2ff7f68e_lib_DefaultSourceManager) new() any {
    return &p3c2ff7f68.DefaultSourceManager{}
}

func (inst* p3c2ff7f68e_lib_DefaultSourceManager) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p3c2ff7f68.DefaultSourceManager)
	nop(ie, com)

	
    com.RegList = inst.getRegList(ie)
    com.Factory = inst.getFactory(ie)


    return nil
}


func (inst*p3c2ff7f68e_lib_DefaultSourceManager) getRegList(ie application.InjectionExt)[]p91658cc84.SourceRegistry{
    dst := make([]p91658cc84.SourceRegistry, 0)
    src := ie.ListComponents(".class-91658cc84667f4d073289b7614060648-SourceRegistry")
    for _, item1 := range src {
        item2 := item1.(p91658cc84.SourceRegistry)
        dst = append(dst, item2)
    }
    return dst
}


func (inst*p3c2ff7f68e_lib_DefaultSourceManager) getFactory(ie application.InjectionExt)p91658cc84.SourceFactory{
    return ie.GetComponent("#alias-91658cc84667f4d073289b7614060648-SourceFactory").(p91658cc84.SourceFactory)
}



// type p3c2ff7f68.DefaultSourceRegistry in package:github.com/starter-go/libredis/src/lib
//
// id:com-3c2ff7f68e8b3d61-lib-DefaultSourceRegistry
// class:class-91658cc84667f4d073289b7614060648-SourceRegistry
// alias:
// scope:singleton
//
type p3c2ff7f68e_lib_DefaultSourceRegistry struct {
}

func (inst* p3c2ff7f68e_lib_DefaultSourceRegistry) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-3c2ff7f68e8b3d61-lib-DefaultSourceRegistry"
	r.Classes = "class-91658cc84667f4d073289b7614060648-SourceRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p3c2ff7f68e_lib_DefaultSourceRegistry) new() any {
    return &p3c2ff7f68.DefaultSourceRegistry{}
}

func (inst* p3c2ff7f68e_lib_DefaultSourceRegistry) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p3c2ff7f68.DefaultSourceRegistry)
	nop(ie, com)

	
    com.AC = inst.getAC(ie)
    com.Factory = inst.getFactory(ie)
    com.NameList = inst.getNameList(ie)


    return nil
}


func (inst*p3c2ff7f68e_lib_DefaultSourceRegistry) getAC(ie application.InjectionExt)p0ef6f2938.Context{
    return ie.GetContext()
}


func (inst*p3c2ff7f68e_lib_DefaultSourceRegistry) getFactory(ie application.InjectionExt)p91658cc84.SourceFactory{
    return ie.GetComponent("#alias-91658cc84667f4d073289b7614060648-SourceFactory").(p91658cc84.SourceFactory)
}


func (inst*p3c2ff7f68e_lib_DefaultSourceRegistry) getNameList(ie application.InjectionExt)string{
    return ie.GetString("${redis.sources}")
}



// type p3c2ff7f68.ServiceImpl in package:github.com/starter-go/libredis/src/lib
//
// id:com-3c2ff7f68e8b3d61-lib-ServiceImpl
// class:
// alias:alias-91658cc84667f4d073289b7614060648-Service
// scope:singleton
//
type p3c2ff7f68e_lib_ServiceImpl struct {
}

func (inst* p3c2ff7f68e_lib_ServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-3c2ff7f68e8b3d61-lib-ServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-91658cc84667f4d073289b7614060648-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p3c2ff7f68e_lib_ServiceImpl) new() any {
    return &p3c2ff7f68.ServiceImpl{}
}

func (inst* p3c2ff7f68e_lib_ServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p3c2ff7f68.ServiceImpl)
	nop(ie, com)

	
    com.ClassMan = inst.getClassMan(ie)
    com.SourceMan = inst.getSourceMan(ie)


    return nil
}


func (inst*p3c2ff7f68e_lib_ServiceImpl) getClassMan(ie application.InjectionExt)p91658cc84.ClassManager{
    return ie.GetComponent("#alias-91658cc84667f4d073289b7614060648-ClassManager").(p91658cc84.ClassManager)
}


func (inst*p3c2ff7f68e_lib_ServiceImpl) getSourceMan(ie application.InjectionExt)p91658cc84.SourceManager{
    return ie.GetComponent("#alias-91658cc84667f4d073289b7614060648-SourceManager").(p91658cc84.SourceManager)
}


