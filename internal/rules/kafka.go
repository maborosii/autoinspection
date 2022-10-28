package rules

type KafkaRule struct {
	BaseRule
	Job                 string  `mapstructure:"job"`
	LagSum              float32 `mapstructure:"lagSum"`
	LagSumIncrease1Day  float32 `mapstructure:"lagSumIncrease1Day"`
	LagSumIncrease1Week float32 `mapstructure:"lagSumIncrease1Week"`
}

func (n *KafkaRule) GetRuleJob() string {
	return n.Job
}

// 指标规则过滤
// kafka 总堆积量判断
func WithKafkaLagSumRuleFilter(nums float32) RuleOption {
	return func(r RuleItf) (interface{}, bool) {
		limit := r.(*KafkaRule).LagSum
		if nums < limit {
			return limit, true
		}
		return limit, false
	}
}

// kafka 总堆积量一天增长率判断
func WithKafkaLagSumIncrease1DayRuleFilter(nums float32) RuleOption {
	return func(r RuleItf) (interface{}, bool) {
		limit := r.(*KafkaRule).LagSumIncrease1Day
		if nums < limit {
			return limit, true
		}
		return limit, false
	}
}

// kafka 总堆积量一周增长率判断
func WithKafkaLagSumIncrease1WeekRuleFilter(nums float32) RuleOption {
	return func(r RuleItf) (interface{}, bool) {
		limit := r.(*KafkaRule).LagSumIncrease1Week
		if nums < limit {
			return limit, true
		}
		return limit, false
	}
}
