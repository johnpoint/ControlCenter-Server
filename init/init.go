package init

import (
	"ControlCenter-Server/init/depend"
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"
)

type Helper struct {
	*log.Logger
	Depends []depend.Depend
}

func (i *Helper) Init(ctx context.Context) error {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("[init] Start")
	for j := range i.Depends {
		if i.Depends[j].GetEnable() {
			fmt.Println(fmt.Sprintf("[init] %s", i.Depends[j].GetName()))
			err := i.Depends[j].Init(ctx)
			if err != nil {
				return err
			}
		}
	}
	fmt.Println("[init] Finish")
	return nil
}

func (i *Helper) AddDepend(depend ...depend.Depend) {
	for j := range depend {
		depend[j].SetEnable(true)
		i.Depends = append(i.Depends, depend[j])
	}
}
