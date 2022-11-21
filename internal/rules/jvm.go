package rules

type JVMRule struct {
	BaseRule
	Job                string  `mapstructure:"job"`
	BlockedThreadCount int8    `mapstructure:"blockedThreadCount"`
	GCTime             float32 `mapstructure:"gcTime"`
	GCCount            float32 `mapstructure:"gcCount"`
}

func (n *JVMRule) GetRuleJob() string {
	return n.Job
}

// 指标规则过滤
// JVM 当前阻塞线程数
func WithJVMBlockedThreadCountRuleFilter(nums int8) RuleOption {
	return func(r RuleItf) (interface{}, bool) {
		limit := r.(*JVMRule).BlockedThreadCount
		if nums == limit {
			return limit, true
		}
		return limit, false
	}
}

// JVM 当前 gc 时间
func WithJVMGarbageCollectTimeRuleFilter(nums float32) RuleOption {
	return func(r RuleItf) (interface{}, bool) {
		limit := r.(*JVMRule).GCTime
		if nums < limit {
			return limit, true
		}
		return limit, false
	}
}

// JVM 当前 gc 次数
func WithJVMGarbageCollectCountRuleFilter(nums float32) RuleOption {
	return func(r RuleItf) (interface{}, bool) {
		limit := r.(*JVMRule).GCCount
		if nums < limit {
			return limit, true
		}
		return limit, false
	}
}
