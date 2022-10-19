package starters

import (
	"fmt"
	"log"
	"node_metrics_go/global"
	"node_metrics_go/infra"
	rs "node_metrics_go/internal/rules"
	"node_metrics_go/pkg/setting"

	"github.com/mitchellh/mapstructure"
)

type RulesStarter struct {
	infra.BaseStarter
}

func (d *RulesStarter) Setup(conf *setting.Config) {
	d.setupRules(conf)
}

func (d *RulesStarter) setupRules(conf *setting.Config) {
	log.Println("init rules setting ...")
	for tt, j := range conf.Rules {
		// parse template config
		if tt == "node" {
			for _, jj := range j {
				nodeRule := new(rs.NodeRule)
				err := mapstructure.Decode(jj, nodeRule)
				if err != nil {
					panic(fmt.Sprintf("mapstructure rules for node_rule occur error: %s", err))
				}
				// global.NotifyRules = append(global.NotifyRules, nodeRule)
				global.NotifyRules[nodeRule.GetRuleJob()] = nodeRule
			}
		}
	}
}
