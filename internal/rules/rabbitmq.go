package rules

type RabbitMQRule struct {
	BaseRule
	Job          string `mapstructure:"job"`
	RunningNodes int8   `mapstructure:"runningNodes"`
	// RunningNodesChange1Day  int8    `mapstructure:"runningNodesChange1Day"`
	// RunningNodesChange1Week int8    `mapstructure:"runningNodesChange1Week"`
	LagSum              float32 `mapstructure:"lagSum"`
	LagSumIncrease1Day  float32 `mapstructure:"lagSumIncrease1Day"`
	LagSumIncrease1Week float32 `mapstructure:"lagSumIncrease1Week"`
}

func (n *RabbitMQRule) GetRuleJob() string {
	return n.Job
}

// 指标规则过滤
// rabbitMQ 当前节点数判断
func WithRabbitMQRunningNodesRuleFilter(nums int8) RuleOption {
	return func(r RuleItf) (interface{}, bool) {
		limit := r.(*RabbitMQRule).RunningNodes
		if nums == limit {
			return limit, true
		}
		return limit, false
	}
}

// rabbitMQ 节点数判断 一天变化数目
// func WithRabbitMQRunningNodesChange1DayRuleFilter(nums int8) RuleOption {
// 	return func(r RuleItf) (interface{}, bool) {
// 		limit := r.(*RabbitMQRule).RunningNodesChange1Day
// 		if nums == limit {
// 			return limit, true
// 		}
// 		return limit, false
// 	}
// }

// rabbitMQ 节点数判断 一周变化数目
// func WithRabbitMQRunningNodesChange1WeekRuleFilter(nums int8) RuleOption {
// 	return func(r RuleItf) (interface{}, bool) {
// 		limit := r.(*RabbitMQRule).RunningNodesChange1Week
// 		if nums == limit {
// 			return limit, true
// 		}
// 		return limit, false
// 	}
// }

// rabbitmq 总堆积量判断
func WithRabbitMQLagSumRuleFilter(nums float32) RuleOption {
	return func(r RuleItf) (interface{}, bool) {
		limit := r.(*RabbitMQRule).LagSum
		if nums < limit {
			return limit, true
		}
		return limit, false
	}
}

// rabbitmq 总堆积量一天增长率判断
func WithRabbitMQLagSumIncrease1DayRuleFilter(nums float32) RuleOption {
	return func(r RuleItf) (interface{}, bool) {
		limit := r.(*RabbitMQRule).LagSumIncrease1Day
		if nums < limit {
			return limit, true
		}
		return limit, false
	}
}

// rabbitmq 总堆积量一周增长率判断
func WithRabbitMQLagSumIncrease1WeekRuleFilter(nums float32) RuleOption {
	return func(r RuleItf) (interface{}, bool) {
		limit := r.(*RabbitMQRule).LagSumIncrease1Week
		if nums < limit {
			return limit, true
		}
		return limit, false
	}
}
