package main

import (
	"os"

	"github.com/starter-go/libredis/modules/libredis"
	"github.com/starter-go/starter"
)

func main() {
	i := starter.Init(os.Args)
	i.MainModule(libredis.ModuleForTest())
	i.WithPanic(true).Run()
}
