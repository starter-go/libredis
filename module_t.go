package libredis

import (
	"embed"

	"github.com/starter-go/application"
)

const (
	theModuleName = "github.com/starter-go/libredis"
	theModuleVer  = "v0.0.4"
	theModuleRev  = 5
)

////////////////////////////////////////////////////////////////////////////////

const (
	theLibModuleResPath  = "src/main/resources"
	theTestModuleResPath = "src/test/resources"
)

//go:embed "src/main/resources"
var theLibModuleResFS embed.FS

//go:embed "src/test/resources"
var theTestModuleResFS embed.FS

////////////////////////////////////////////////////////////////////////////////

// NewModuleForLib module template for []
func NewModuleForLib() *application.ModuleBuilder {

	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName + "#lib")
	mb.Version(theModuleVer)
	mb.Revision(theModuleRev)
	mb.EmbedResources(theLibModuleResFS, theLibModuleResPath)
	return mb
}

// NewModuleForTest module template for []
func NewModuleForTest() *application.ModuleBuilder {

	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName + "#test")
	mb.Version(theModuleVer)
	mb.Revision(theModuleRev)
	mb.EmbedResources(theTestModuleResFS, theTestModuleResPath)
	return mb
}

////////////////////////////////////////////////////////////////////////////////
