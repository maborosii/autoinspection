package starters

import (
	"fmt"
	"log"
	"node_metrics_go/global"
	"node_metrics_go/infra"
	rs "node_metrics_go/internal/rules"
	"node_metrics_go/pkg/setting"
	"sync"

	"github.com/mitchellh/mapstructure"
)

type RulesStarter struct {
	infra.BaseStarter
}

func (d *RulesStarter) Setup(conf *setting.Config) {
	d.setupRules(conf)
}

var doReg sync.Once
var registerMap = make(map[string]ruleRegister)

func (d *RulesStarter) setupRules(conf *setting.Config) {
	log.Println("init rules setting ...")
	doReg.Do(func() {
		registerMap["node"] = new(nodeRulesRegister)
		registerMap["redis"] = new(redisRulesRegister)
		registerMap["kafka"] = new(kafkaRulesRegister)
		registerMap["rabbitmq"] = new(rabbitMQRulesRegister)
		registerMap["es"] = new(elasticSearchRulesRegister)
	})

	for tt, j := range conf.Rules {
		// parse template config
		if _, ok := registerMap[tt]; !ok {
			panic(fmt.Sprintf("not suitable rule type in config, rule.type: %s", tt))
		}
		registerMap[tt].register(j)
	}
}

type ruleRegister interface {
	register([]interface{})
}

type nodeRulesRegister struct{}

func (n nodeRulesRegister) register(confRules []interface{}) {
	for _, jj := range confRules {
		nodeRule := new(rs.NodeRule)
		err := mapstructure.Decode(jj, nodeRule)
		if err != nil {
			panic(fmt.Sprintf("mapstructure rules for node occur error: %s", err))
		}
		global.NotifyRules[nodeRule.GetRuleJob()] = nodeRule
	}
}

type redisRulesRegister struct{}

func (n redisRulesRegister) register(confRules []interface{}) {
	for _, jj := range confRules {
		redisRule := new(rs.RedisRule)
		err := mapstructure.Decode(jj, redisRule)
		if err != nil {
			panic(fmt.Sprintf("mapstructure rules for redis occur error: %s", err))
		}
		global.NotifyRules[redisRule.GetRuleJob()] = redisRule
	}
}

type kafkaRulesRegister struct{}

func (n kafkaRulesRegister) register(confRules []interface{}) {
	for _, jj := range confRules {
		kafkaRule := new(rs.KafkaRule)
		err := mapstructure.Decode(jj, kafkaRule)
		if err != nil {
			panic(fmt.Sprintf("mapstructure rules for kafka occur error: %s", err))
		}
		global.NotifyRules[kafkaRule.GetRuleJob()] = kafkaRule
	}
}

type rabbitMQRulesRegister struct{}

func (n rabbitMQRulesRegister) register(confRules []interface{}) {
	for _, jj := range confRules {
		rabbitMQRule := new(rs.RabbitMQRule)
		err := mapstructure.Decode(jj, rabbitMQRule)
		if err != nil {
			panic(fmt.Sprintf("mapstructure rules for rabbitmq occur error: %s", err))
		}
		global.NotifyRules[rabbitMQRule.GetRuleJob()] = rabbitMQRule
	}
}

type elasticSearchRulesRegister struct{}

func (n elasticSearchRulesRegister) register(confRules []interface{}) {
	for _, jj := range confRules {
		elasticSearchRule := new(rs.ElasticSearchRule)
		err := mapstructure.Decode(jj, elasticSearchRule)
		if err != nil {
			panic(fmt.Sprintf("mapstructure rules for elasticSearch occur error: %s", err))
		}
		global.NotifyRules[elasticSearchRule.GetRuleJob()] = elasticSearchRule
	}
}
