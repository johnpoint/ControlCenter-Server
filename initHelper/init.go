package initHelper

import (
	"ControlCenter/config"
	"ControlCenter/initHelper/depend"
	"context"
	"fmt"
	"log"
	"math/rand"
	"reflect"
	"strings"
	"time"
)

type Helper struct {
	*log.Logger
	Depends []depend.Depend
}

func (i *Helper) Init(ctx context.Context) error {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("[depend] Start")
	for j := range i.Depends {
		fmt.Println(strings.Replace(fmt.Sprintf("[depend] %s", reflect.TypeOf(i.Depends[j])), "*depend.", "", 1))
		err := i.Depends[j].Init(ctx, config.Config)
		if err != nil {
			return err
		}
	}
	fmt.Println("[depend] Finish")
	return nil
}

func (i *Helper) AddDepend(depend ...depend.Depend) *Helper {
	for j := range depend {
		i.Depends = append(i.Depends, depend[j])
	}
	return i
}
