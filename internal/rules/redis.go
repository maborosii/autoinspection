package rules

type RedisRule struct {
	BaseRule
	Job                string  `mapstructure:"job"`
	Conns              float32 `mapstructure:"conns"`
	ConnsIncrease1Day  float32 `mapstructure:"connsIncrease1Day"`
	ConnsIncrease1Week float32 `mapstructure:"connsIncrease1Week"`
	Mem                float32 `mapstructure:"mem"`
	MemIncrease1Day    float32 `mapstructure:"memIncrease1Day"`
	MemIncrease1Week   float32 `mapstructure:"memIncrease1Week"`
}

func (n *RedisRule) GetRuleJob() string {
	return n.Job
}

// 指标规则过滤
// redis 连接数判断
func WithRedisConnsRuleFilter(nums float32) RuleOption {
	return func(r RuleItf) (interface{}, bool) {
		limit := r.(*RedisRule).Conns
		if nums < limit {
			return limit, true
		}
		return limit, false
	}
}

// redis 连接数一天增长率判断
func WithRedisConnsIncrease1DayRuleFilter(nums float32) RuleOption {
	return func(r RuleItf) (interface{}, bool) {
		limit := r.(*RedisRule).ConnsIncrease1Day
		if nums < limit {
			return limit, true
		}
		return limit, false
	}
}

// redis 连接数一周增长率判断
func WithRedisConnsIncrease1WeekRuleFilter(nums float32) RuleOption {
	return func(r RuleItf) (interface{}, bool) {
		limit := r.(*RedisRule).ConnsIncrease1Week
		if nums < limit {
			return limit, true
		}
		return limit, false
	}
}

// redis 内存使用率判断
func WithRedisMemRuleFilter(nums float32) RuleOption {
	return func(r RuleItf) (interface{}, bool) {
		limit := r.(*RedisRule).Mem
		if nums < limit {
			return limit, true
		}
		return limit, false
	}
}

// redis 连接数一天增长率判断
func WithRedisMemIncrease1DayRuleFilter(nums float32) RuleOption {
	return func(r RuleItf) (interface{}, bool) {
		limit := r.(*RedisRule).MemIncrease1Day
		if nums < limit {
			return limit, true
		}
		return limit, false
	}
}

// redis 连接数一周增长率判断
func WithRedisMemIncrease1WeekRuleFilter(nums float32) RuleOption {
	return func(r RuleItf) (interface{}, bool) {
		limit := r.(*RedisRule).MemIncrease1Week
		if nums < limit {
			return limit, true
		}
		return limit, false
	}
}
