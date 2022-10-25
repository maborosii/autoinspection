package rules

type NodeRule struct {
	BaseRule
	Job                  string  `mapstructure:"job"`
	Cpu                  float32 `mapstructure:"cpu"`
	CpuIncrease1Day      float32 `mapstructure:"cpuIncrease1Day"`
	CpuIncrease1Week     float32 `mapstructure:"cpuIncrease1Week"`
	Mem                  float32 `mapstructure:"mem"`
	MemIncrease1Day      float32 `mapstructure:"memIncrease1Day"`
	MemIncrease1Week     float32 `mapstructure:"memIncrease1Week"`
	Disk                 float32 `mapstructure:"disk"`
	DiskIncrease1Day     float32 `mapstructure:"diskIncrease1Day"`
	DiskIncrease1Week    float32 `mapstructure:"diskIncrease1Week"`
	TcpConn              float32 `mapstructure:"tcpConn"`
	TcpConnIncrease1Day  float32 `mapstructure:"tcpConnIncrease1Day"`
	TcpConnIncrease1Week float32 `mapstructure:"tcpConnIncrease1Week"`
}

func (n *NodeRule) GetRuleJob() string {
	return n.Job
}

// 指标规则过滤
// cpu 使用率判断
func WithCpuRuleFilter(nums float32) RuleOption {
	return func(r RuleItf) (interface{}, bool) {
		limit := r.(*NodeRule).Cpu
		if nums < limit {
			return limit, true
		}
		return limit, false
	}
}

// cpu 一天增长率判断
func WithCpuIncrease1DayRuleFilter(nums float32) RuleOption {
	return func(r RuleItf) (interface{}, bool) {
		limit := r.(*NodeRule).CpuIncrease1Day
		if nums < limit {
			return limit, true
		}
		return limit, false
	}
}

// cpu 一周增长率判断
func WithCpuIncrease1WeekRuleFilter(nums float32) RuleOption {
	return func(r RuleItf) (interface{}, bool) {
		limit := r.(*NodeRule).CpuIncrease1Week
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
func WithTcpConnRuleFilter(nums float32) RuleOption {
	return func(r RuleItf) (interface{}, bool) {
		limit := r.(*NodeRule).TcpConn
		if nums < limit {
			return limit, true
		}
		return limit, false
	}
}

// tcp 连接数一天增长率判断
func WithTcpConnIncrease1DayRuleFilter(nums float32) RuleOption {
	return func(r RuleItf) (interface{}, bool) {
		limit := r.(*NodeRule).TcpConnIncrease1Day
		if nums < limit {
			return limit, true
		}
		return limit, false
	}
}

// tcp 连接数一周增长率判断
func WithTcpConnIncrease1WeekRuleFilter(nums float32) RuleOption {
	return func(r RuleItf) (interface{}, bool) {
		limit := r.(*NodeRule).TcpConnIncrease1Week
		if nums < limit {
			return limit, true
		}
		return limit, false
	}
}