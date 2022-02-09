package bootstrap

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"reflect"
	"time"
)

var globalComponent = make([]Component, 0)

func AddGlobalComponent(components ...Component) {
	globalComponent = append(globalComponent, components...)
}

func NewBoot(ctx context.Context, components ...Component) *Helper {
	return &Helper{
		components: components,
	}
}

type Helper struct {
	*log.Logger
	components []Component
}

func (i *Helper) loadGlobalComponent(ctx context.Context) error {
	for j := range globalComponent {
		fmt.Println(fmt.Sprintf("[Bootstrap] %s", reflect.TypeOf(globalComponent[j])))
		err := globalComponent[j].Init(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *Helper) loadComponent(ctx context.Context) error {
	for j := range i.components {
		fmt.Println(fmt.Sprintf("[Bootstrap] %s", reflect.TypeOf(i.components[j])))
		err := i.components[j].Init(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *Helper) InitWithoutGlobalComponent(ctx context.Context) error {
	fmt.Println("[Bootstrap] Start")
	rand.Seed(time.Now().UnixNano())
	err := i.loadComponent(ctx)
	if err != nil {
		return err
	}
	fmt.Println("[Bootstrap] Finish")
	return nil
}

func (i *Helper) Init(ctx context.Context) error {
	fmt.Println("[Bootstrap] Start")
	rand.Seed(time.Now().UnixNano())
	err := i.loadGlobalComponent(ctx)
	if err != nil {
		return err
	}
	err = i.loadComponent(ctx)
	if err != nil {
		return err
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
