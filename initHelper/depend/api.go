package depend

import (
	"ControlCenter-Server/config"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Api struct {
	Enable bool
}

func (r *Api) Init(ctx context.Context) error {
	gin.SetMode(gin.ReleaseMode)
	routerGin := gin.New()

	go func() {
		fmt.Println("[init] HTTP Listen at " + config.Config.HttpServerListen)
		err := routerGin.Run(config.Config.HttpServerListen)
		if err != nil {
			panic(err)
		}
	}()
	return nil
}

func (r *Api) GetEnable() bool {
	return r.Enable
}

func (r *Api) SetEnable(enable bool) {
	r.Enable = enable
}

func (r *Api) GetName() string {
	return "API"
}

func (r *Api) GetDesc() string {
	return "个号日志API"
}
