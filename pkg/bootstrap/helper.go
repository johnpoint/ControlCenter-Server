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
		ctx:        ctx,
		components: components,
		logger:     NewDefaultLogger(),
	}
}

type Helper struct {
	ctx        context.Context
	logger     Logger
	components []Component
}

func (i *Helper) WithLogger(logger Logger) *Helper {
	i.logger = logger
	return i
}

func (i *Helper) WithContext(ctx context.Context) *Helper {
	i.ctx = ctx
	return i
}

func (i *Helper) loadGlobalComponent() error {
	for j := range globalComponent {
		i.logger.Info("Bootstrap", zap.String("step", reflect.TypeOf(globalComponent[j]).String()))
		err := globalComponent[j].Init(i.ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *Helper) loadComponent() error {
	for j := range i.components {
		i.logger.Info("Bootstrap", zap.String("step", reflect.TypeOf(i.components[j]).String()))
		err := i.components[j].Init(i.ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *Helper) InitWithoutGlobalComponent() error {
	if i.logger == nil {
		i.logger = NewDefaultLogger()
	}
	if i.ctx == nil {
		i.ctx = context.TODO()
	}
	i.logger.Info("Bootstrap", zap.String("step", "start"))
	rand.Seed(time.Now().UnixNano())
	err := i.loadComponent()
	if err != nil {
		return err
	}
	i.logger.Info("Bootstrap", zap.String("step", "finish"))
	return nil
}

func (i *Helper) Init() error {
	if i.logger == nil {
		i.logger = NewDefaultLogger()
	}
	i.logger.Info("Bootstrap", zap.String("step", "start"))
	rand.Seed(time.Now().UnixNano())
	err := i.loadGlobalComponent()
	if err != nil {
		return err
	}
	err = i.loadComponent()
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
