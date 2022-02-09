package bootstrap

import (
	"context"
	"go.uber.org/zap"
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
		logger:     NewDefaultLogger(),
	}
}

type Helper struct {
	logger     Logger
	components []Component
}

func (i *Helper) loadGlobalComponent(ctx context.Context) error {
	for j := range globalComponent {
		i.logger.Info("Bootstrap", zap.String("load", reflect.TypeOf(globalComponent[j]).String()))
		err := globalComponent[j].Init(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *Helper) loadComponent(ctx context.Context) error {
	for j := range i.components {
		i.logger.Info("Bootstrap", zap.String("step", reflect.TypeOf(i.components[j]).String()))
		err := i.components[j].Init(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *Helper) InitWithoutGlobalComponent(ctx context.Context) error {
	if i.logger == nil {
		i.logger = NewDefaultLogger()
	}
	i.logger.Info("Bootstrap", zap.String("step", "start"))
	rand.Seed(time.Now().UnixNano())
	err := i.loadComponent(ctx)
	if err != nil {
		return err
	}
	i.logger.Info("Bootstrap", zap.String("step", "finish"))
	return nil
}

func (i *Helper) Init(ctx context.Context) error {
	if i.logger == nil {
		i.logger = NewDefaultLogger()
	}
	i.logger.Info("Bootstrap", zap.String("step", "start"))
	rand.Seed(time.Now().UnixNano())
	err := i.loadGlobalComponent(ctx)
	if err != nil {
		return err
	}
	err = i.loadComponent(ctx)
	if err != nil {
		return err
	}
	i.logger.Info("Bootstrap", zap.String("step", "finish"))
	return nil
}

func (i *Helper) AddComponent(components ...Component) *Helper {
	for j := range components {
		i.components = append(i.components, components[j])
	}
	return i
}
