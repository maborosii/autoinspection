package starters

import (
	"log"
	"node_metrics_go/infra"
	"node_metrics_go/pkg/logger"
	"node_metrics_go/pkg/setting"
)

type LogStarter struct {
	infra.BaseStarter
}

func (l *LogStarter) Setup(conf *setting.Config) {
	log.Println("init logger ..")
	if err := logger.InitLogger(conf.LogConfig); err != nil {
		panic(err)
	}
}
