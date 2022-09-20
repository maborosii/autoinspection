package rules

type NodeRule struct {
	BaseRule
	Job         string  `mapstructure:"job"`
	Cpu         float32 `mapstructure:"cpu"`
	CpuIncrease float32 `mapstructure:"cpuIncrease"`
	Mem         float32 `mapstructure:"mem"`
	MemIncrease float32 `mapstructure:"memIncrease"`
}

func (n *NodeRule) GetJob() string {
	return n.Job
}
