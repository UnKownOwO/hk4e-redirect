package app

import (
	"context"
	"hk4e-redirect/common/config"
	"hk4e-redirect/pkg/logger"
	"os"
	"os/signal"
	"syscall"
)

func Run(ctx context.Context, configFile string) error {
	config.InitConfig(configFile)

	logger.InitLogger("hk4e-redirect")
	logger.Warn("hk4e-redirect start")
	defer func() {
		logger.CloseLogger()
	}()

	_ = NewController()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		select {
		case <-ctx.Done():
			return nil
		case s := <-c:
			logger.Warn("get a signal %s", s.String())
			switch s {
			case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
				logger.Warn("dispatch exit")
				return nil
			case syscall.SIGHUP:
			default:
				return nil
			}
		}
	}
}
