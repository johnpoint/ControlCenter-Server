package depend

import (
	"ControlCenter/pkg/bootstrap"
	"context"
)

type PerformanceCollector struct{}

var _ bootstrap.Component = (*PerformanceCollector)(nil)

func (p *PerformanceCollector) Init(ctx context.Context) error {
	return nil
}
