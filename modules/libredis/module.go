package libredis

import (
	"github.com/starter-go/application"
	"github.com/starter-go/libredis"
	"github.com/starter-go/libredis/gen/gen4lib"
)

// Module ... 导出模块：github.com/starter-go/libredis
func Module() application.Module {
	mb := libredis.ModuleT()
	mb.Components(gen4lib.ComponentsForLibRedis)
	return mb.Create()
}
