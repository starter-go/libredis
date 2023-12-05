package libredis

import (
	"embed"

	"github.com/starter-go/application"
	"github.com/starter-go/starter"
)

const (
	theModuleName    = "github.com/starter-go/libredis"
	theModuleVer     = "v0.0.2"
	theModuleRev     = 3
	theModuleResPath = "src/lib/resources"
)

//go:embed "src/lib/resources"
var theModuleResFS embed.FS

// ModuleT module template for []
func ModuleT() *application.ModuleBuilder {

	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName)
	mb.Version(theModuleVer)
	mb.Revision(theModuleRev)
	mb.EmbedResources(theModuleResFS, theModuleResPath)

	mb.Depend(starter.Module())

	return mb
}
