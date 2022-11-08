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
		switch tt {
		case "node":
			for _, jj := range j {
				nodeRule := new(rs.NodeRule)
				err := mapstructure.Decode(jj, nodeRule)
				if err != nil {
					panic(fmt.Sprintf("mapstructure rules for node occur error: %s", err))
				}
				global.NotifyRules[nodeRule.GetRuleJob()] = nodeRule
			}
		case "redis":
			for _, jj := range j {
				redisRule := new(rs.RedisRule)
				err := mapstructure.Decode(jj, redisRule)
				if err != nil {
					panic(fmt.Sprintf("mapstructure rules for redis occur error: %s", err))
				}
				global.NotifyRules[redisRule.GetRuleJob()] = redisRule
			}
		case "kafka":
			for _, jj := range j {
				kafkaRule := new(rs.KafkaRule)
				err := mapstructure.Decode(jj, kafkaRule)
				if err != nil {
					panic(fmt.Sprintf("mapstructure rules for kafka occur error: %s", err))
				}
				global.NotifyRules[kafkaRule.GetRuleJob()] = kafkaRule
			}
		case "rabbitmq":
			for _, jj := range j {
				rabbitMQRule := new(rs.RabbitMQRule)
				err := mapstructure.Decode(jj, rabbitMQRule)
				if err != nil {
					panic(fmt.Sprintf("mapstructure rules for rabbitMQ occur error: %s", err))
				}
				global.NotifyRules[rabbitMQRule.GetRuleJob()] = rabbitMQRule
			}
		case "es":
			for _, jj := range j {
				elasticSearchRule := new(rs.ElasticSearchRule)
				err := mapstructure.Decode(jj, elasticSearchRule)
				if err != nil {
					panic(fmt.Sprintf("mapstructure rules for elasticSearch occur error: %s", err))
				}
				global.NotifyRules[elasticSearchRule.GetRuleJob()] = elasticSearchRule
			}
		default:
			panic(fmt.Sprintf("not suitable rule type in config, rule.type: %s", tt))
		}
	}
}
