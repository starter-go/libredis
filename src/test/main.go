package main

import (
	"embed"
	"os"

	"github.com/starter-go/application"
	"github.com/starter-go/libredis/gen/gen4test"
	libredismod "github.com/starter-go/libredis/modules/libredis"
	"github.com/starter-go/starter"
)

func main() {
	i := starter.Init(os.Args)
	i.MainModule(theModule())
	i.WithPanic(true).Run()
}

const (
	mName    = "github.com/starter-go/libredis#test"
	mVer     = "v0.0.1"
	mRev     = 1
	mResPath = "resources"
)

//go:embed "resources"
var mResFS embed.FS

func theModule() application.Module {
	mb := application.ModuleBuilder{}
	mb.Name(mName).Version(mVer).Revision(mRev)
	mb.EmbedResources(mResFS, mResPath)
	mb.Components(gen4test.ComponentsForTestLibRedis)
	mb.Depend(libredismod.Module())
	return mb.Create()
}
