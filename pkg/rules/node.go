package rules

type NodeRule struct {
	BaseRule
	Job             string  `mapstructure:"job"`
	Cpu             float32 `mapstructure:"cpu"`
	CpuIncrease     float32 `mapstructure:"cpuIncrease"`
	Mem             float32 `mapstructure:"mem"`
	MemIncrease     float32 `mapstructure:"memIncrease"`
	Disk            float32 `mapstructure:"disk"`
	DiskIncrease    float32 `mapstructure:"diskIncrease"`
	TcpConn         float32 `mapstructure:"tcpConn"`
	TcpConnIncrease float32 `mapstructure:"tcpConnIncrease"`
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

// cpu 一周增长率判断
func WithCpuIncreaseRuleFilter(nums float32) RuleOption {
	return func(r RuleItf) (interface{}, bool) {
		limit := r.(*NodeRule).CpuIncrease
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

// 内存一周增长判断
func WithMemIncreaseRuleFilter(nums float32) RuleOption {
	return func(r RuleItf) (interface{}, bool) {
		limit := r.(*NodeRule).MemIncrease
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

// 磁盘一周增长率判断
func WithDiskIncreaseRuleFilter(nums float32) RuleOption {
	return func(r RuleItf) (interface{}, bool) {
		limit := r.(*NodeRule).DiskIncrease
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

// tcp 连接数一周增长率判断
func WithTcpConnIncreaseRuleFilter(nums float32) RuleOption {
	return func(r RuleItf) (interface{}, bool) {
		limit := r.(*NodeRule).TcpConnIncrease
		if nums < limit {
			return limit, true
		}
		return limit, false
	}
}
