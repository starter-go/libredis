package libredis

import (
	"github.com/starter-go/application"
	"github.com/starter-go/keyvalues/modules/keyvalues"
	"github.com/starter-go/libredis"
	"github.com/starter-go/libredis/gen/lib4libredis"
	"github.com/starter-go/libredis/gen/test4libredis"
)

// Module ... 导出模块：github.com/starter-go/libredis
func Module() application.Module {
	mb := libredis.NewModuleForLib()
	mb.Components(lib4libredis.ComponentsForLibRedis)
	// mb.Depend(starter.Module())
	mb.Depend(keyvalues.Module())
	return mb.Create()
}

// ModuleForTest ... 导出模块：github.com/starter-go/libredis#test
func ModuleForTest() application.Module {
	mb := libredis.NewModuleForTest()
	mb.Components(test4libredis.ComponentsForTestLibRedis)
	mb.Depend(Module())
	return mb.Create()
}
