package rules

type ElasticSearchRule struct {
	BaseRule
	Job                     string `mapstructure:"job"`
	HealthStatus            int8   `mapstructure:"healthStatus"`
	HealthStatusChange1Day  int8   `mapstructure:"healthStatusChange1Day"`
	HealthStatusChange1Week int8   `mapstructure:"healthStatusChange1Week"`
}

func (n *ElasticSearchRule) GetRuleJob() string {
	return n.Job
}

// 指标规则过滤
// es 当前健康状态判断
func WithElasticSearchHealthStatusRuleFilter(nums int8) RuleOption {
	return func(r RuleItf) (interface{}, bool) {
		limit := r.(*ElasticSearchRule).HealthStatus
		if nums == limit {
			return limit, true
		}
		return limit, false
	}
}

// es 健康状态判断 一天变化
func WithElasticSearchHealthStatusChange1DayRuleFilter(nums int8) RuleOption {
	return func(r RuleItf) (interface{}, bool) {
		limit := r.(*ElasticSearchRule).HealthStatusChange1Day
		if nums == limit {
			return limit, true
		}
		return limit, false
	}
}

// es 健康状态判断 一周变化
func WithElasticSearchHealthStatusChange1WeekRuleFilter(nums int8) RuleOption {
	return func(r RuleItf) (interface{}, bool) {
		limit := r.(*ElasticSearchRule).HealthStatusChange1Week
		if nums == limit {
			return limit, true
		}
		return limit, false
	}
}
