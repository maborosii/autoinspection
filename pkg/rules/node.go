package rules

type NodeRule struct {
	BaseRule
	Job         string  `mapstructure:"job"`
	Cpu         float32 `mapstructure:"cpu"`
	CpuIncrease float32 `mapstructure:"cpuIncrease"`
	Mem         float32 `mapstructure:"mem"`
	MemIncrease float32 `mapstructure:"memIncrease"`
}

func (n *NodeRule) GetRuleJob() string {
	return n.Job
}
func WithCpuRuleFilter(nums float32) RuleOption {
	return func(r RuleItf) bool {
		if nums < r.(*NodeRule).Cpu {
			return true
		}
		return false
	}

}
func WithMemRuleFilter(nums float32) RuleOption {
	return func(r RuleItf) bool {
		if nums < r.(*NodeRule).Mem {
			return true
		}
		return false
	}
}
func WithMemIncreaseRuleFilter(nums float32) RuleOption {
	return func(r RuleItf) bool {
		if nums < r.(*NodeRule).MemIncrease {
			return true
		}
		return false
	}
}
func WithCpuIncreaseRuleFilter(nums float32) RuleOption {
	return func(r RuleItf) bool {
		if nums < r.(*NodeRule).CpuIncrease {
			return true
		}
		return false
	}
}
