package main

import (
	"fmt"
	"time"

	"github.com/yankooo/go-archaius"
	"github.com/yankooo/go-archaius/event"
	"github.com/yankooo/go-archaius/source/remote"
	_ "github.com/yankooo/go-archaius/source/remote/kie"
	"github.com/go-mesh/openlogging"
)

type Listener struct {
	Key string
}

func (li *Listener) Event(event *event.Event) {
	openlogging.GetLogger().Infof("change event :%+v", *event)
}

func main() {
	var kieObj = &archaius.RemoteInfo{
		DefaultDimension: map[string]string{
			remote.LabelApp:         "foo",
			remote.LabelService:     "bar",
			remote.LabelVersion:     "1.0.0",
			remote.LabelEnvironment: "prod",
		},
		URL:         "http://127.0.0.1:30110",
		RefreshMode: remote.ModeWatch,
	}
	if err := archaius.Init(archaius.WithRemoteSource(archaius.KieSource, kieObj)); err != nil {
		fmt.Println(err)
	}
	if err := archaius.RegisterListener(&Listener{}, "user"); err != nil {
		fmt.Println(err)
	}
	for {
		fmt.Println("current user: ", archaius.Get("user"))
		time.Sleep(time.Second * 3)
	}
}
