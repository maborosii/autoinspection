package starters

import (
	"log"
	"node_metrics_go/global"
	"node_metrics_go/infra"
	"node_metrics_go/pkg/setting"
)

type MonitorStarter struct {
	infra.BaseStarter
}

func (d *MonitorStarter) Setup(conf *setting.Config) {
	d.setupMonitor(conf)
}

func (d *MonitorStarter) setupMonitor(conf *setting.Config) {
	log.Println("init monitor setting ...")
	global.MonitorSetting = conf
}
