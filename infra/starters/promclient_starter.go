package starters

import (
	"log"
	"node_metrics_go/global"
	"node_metrics_go/infra"
	"node_metrics_go/pkg/prom"
	"node_metrics_go/pkg/setting"
)

type PromClientsStarter struct {
	infra.BaseStarter
}

func (d *PromClientsStarter) Setup(conf *setting.Config) {
	d.setupPromClients(conf)
}

func (d *PromClientsStarter) setupPromClients(conf *setting.Config) {
	log.Println("init prometheus clients setting ...")
	for tt, j := range conf.Endpoints {
		log.Println(tt)
		global.PromClients[tt] = prom.ClientForProm(j)
	}
}
