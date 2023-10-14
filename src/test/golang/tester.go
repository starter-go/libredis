package golang

import (
	"github.com/starter-go/application"
	"github.com/starter-go/libredis"
)

// Tester ...
type Tester struct {

	//starter:component

	SourceManager libredis.SourceManager //starter:inject("#")

}

func (inst *Tester) _impl(a application.Lifecycle) {
	a = inst
}

// Life ...
func (inst *Tester) Life() *application.Life {
	return &application.Life{
		OnLoop: inst.run,
	}
}

func (inst *Tester) run() error {

	src, err := inst.SourceManager.GetSource(libredis.SourceMaster)
	if err != nil {
		return err
	}

	client := src.Client(nil)
	return client.Ping()
}
