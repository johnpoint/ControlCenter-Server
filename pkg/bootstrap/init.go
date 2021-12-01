package bootstrap

import (
	"context"
	"fmt"
	"log"
	"reflect"
)

type Helper struct {
	*log.Logger
	components []Component
}

func (i *Helper) Init(ctx context.Context) error {
	fmt.Println("[Bootstrap] Start")
	for j := range i.components {
		fmt.Println(fmt.Sprintf("[Bootstrap] %s", reflect.TypeOf(i.components[j])))
		err := i.components[j].Init(ctx)
		if err != nil {
			return err
		}
	}
	fmt.Println("[Bootstrap] Finish")
	return nil
}

func (i *Helper) AddComponent(components ...Component) *Helper {
	for j := range components {
		i.components = append(i.components, components[j])
	}
	return i
}
