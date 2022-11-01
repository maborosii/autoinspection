package rules

type NodeRule struct {
	BaseRule
	Job                  string  `mapstructure:"job"`
	CPU                  float32 `mapstructure:"cpu"`
	CPUIncrease1Day      float32 `mapstructure:"cpuIncrease1Day"`
	CPUIncrease1Week     float32 `mapstructure:"cpuIncrease1Week"`
	Mem                  float32 `mapstructure:"mem"`
	MemIncrease1Day      float32 `mapstructure:"memIncrease1Day"`
	MemIncrease1Week     float32 `mapstructure:"memIncrease1Week"`
	Disk                 float32 `mapstructure:"disk"`
	DiskIncrease1Day     float32 `mapstructure:"diskIncrease1Day"`
	DiskIncrease1Week    float32 `mapstructure:"diskIncrease1Week"`
	TCPConn              float32 `mapstructure:"tcpConn"`
	TCPConnIncrease1Day  float32 `mapstructure:"tcpConnIncrease1Day"`
	TCPConnIncrease1Week float32 `mapstructure:"tcpConnIncrease1Week"`
}

func (n *NodeRule) GetRuleJob() string {
	return n.Job
}

// 指标规则过滤
// cpu 使用率判断
func WithCPURuleFilter(nums float32) RuleOption {
	return func(r RuleItf) (interface{}, bool) {
		limit := r.(*NodeRule).CPU
		if nums < limit {
			return limit, true
		}
		return limit, false
	}
}

// cpu 一天增长率判断
func WithCPUIncrease1DayRuleFilter(nums float32) RuleOption {
	return func(r RuleItf) (interface{}, bool) {
		limit := r.(*NodeRule).CPUIncrease1Day
		if nums < limit {
			return limit, true
		}
		return limit, false
	}
}

// cpu 一周增长率判断
func WithCPUIncrease1WeekRuleFilter(nums float32) RuleOption {
	return func(r RuleItf) (interface{}, bool) {
		limit := r.(*NodeRule).CPUIncrease1Week
		if nums < limit {
			return limit, true
		}
		return limit, false
	}
}

// 内存使用率判断
func WithMemRuleFilter(nums float32) RuleOption {
	return func(r RuleItf) (interface{}, bool) {
		limit := r.(*NodeRule).Mem
		if nums < limit {
			return limit, true
		}
		return limit, false
	}
}

// 内存一天增长判断
func WithMemIncrease1DayRuleFilter(nums float32) RuleOption {
	return func(r RuleItf) (interface{}, bool) {
		limit := r.(*NodeRule).MemIncrease1Day
		if nums < limit {
			return limit, true
		}
		return limit, false
	}
}

// 内存一周增长判断
func WithMemIncrease1WeekRuleFilter(nums float32) RuleOption {
	return func(r RuleItf) (interface{}, bool) {
		limit := r.(*NodeRule).MemIncrease1Week
		if nums < limit {
			return limit, true
		}
		return limit, false
	}
}

// 磁盘使用率判断
func WithDiskRuleFilter(nums float32) RuleOption {
	return func(r RuleItf) (interface{}, bool) {
		limit := r.(*NodeRule).Disk
		if nums < limit {
			return limit, true
		}
		return limit, false
	}
}

// 磁盘一天增长率判断
func WithDiskIncrease1DayRuleFilter(nums float32) RuleOption {
	return func(r RuleItf) (interface{}, bool) {
		limit := r.(*NodeRule).DiskIncrease1Day
		if nums < limit {
			return limit, true
		}
		return limit, false
	}
}

// 磁盘一周增长率判断
func WithDiskIncrease1WeekRuleFilter(nums float32) RuleOption {
	return func(r RuleItf) (interface{}, bool) {
		limit := r.(*NodeRule).DiskIncrease1Week
		if nums < limit {
			return limit, true
		}
		return limit, false
	}
}

// tcp 连接数判断
func WithTCPConnRuleFilter(nums float32) RuleOption {
	return func(r RuleItf) (interface{}, bool) {
		limit := r.(*NodeRule).TCPConn
		if nums < limit {
			return limit, true
		}
		return limit, false
	}
}

// tcp 连接数一天增长率判断
func WithTCPConnIncrease1DayRuleFilter(nums float32) RuleOption {
	return func(r RuleItf) (interface{}, bool) {
		limit := r.(*NodeRule).TCPConnIncrease1Day
		if nums < limit {
			return limit, true
		}
		return limit, false
	}
}

// tcp 连接数一周增长率判断
func WithTCPConnIncrease1WeekRuleFilter(nums float32) RuleOption {
	return func(r RuleItf) (interface{}, bool) {
		limit := r.(*NodeRule).TCPConnIncrease1Week
		if nums < limit {
			return limit, true
		}
		return limit, false
	}
}
