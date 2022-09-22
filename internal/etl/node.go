package etl

import "fmt"

type NodeMetrics struct {
	*BaseMetrics
	NodeName       string
	CpuUsage       float32
	BeforeCpuUsage float32
	MemUsage       float32
	BeforeMemUsage float32
	// diskUsage        []Disk
}

// type Disk struct {
// 	mountPoint   string
// 	usage        string
// 	increaseDisk string
// }

// type NodeOption func(*NodeMetrics)

// pre asset
// var a MetricsItf
// var _ = a.(*NodeMetrics)

func WithNodeJob(job string) MetricsOption {
	return func(sr MetricsItf) {
		sr.(*NodeMetrics).Job = job
	}
}
func WithNodeName(nodeName string) MetricsOption {
	return func(sr MetricsItf) {
		sr.(*NodeMetrics).NodeName = nodeName
	}
}

func WithCpuUsage(cpuUsage float32) MetricsOption {
	return func(sr MetricsItf) {
		sr.(*NodeMetrics).CpuUsage = cpuUsage
	}
}
func WithBeforeCpuUsage(beforeCpuUsage float32) MetricsOption {
	return func(sr MetricsItf) {
		sr.(*NodeMetrics).BeforeCpuUsage = beforeCpuUsage
	}
}

func WithMemUsage(memUsage float32) MetricsOption {
	return func(sr MetricsItf) {
		sr.(*NodeMetrics).MemUsage = memUsage
	}
}

func WithBeforeMemUsage(beforeMemUsage float32) MetricsOption {
	return func(sr MetricsItf) {
		sr.(*NodeMetrics).BeforeMemUsage = beforeMemUsage
	}
}

func NewNodeMetrics(instance string, options ...MetricsOption) *NodeMetrics {
	mi := &BaseMetrics{Instance: instance}
	sr := &NodeMetrics{
		BaseMetrics: mi,
	}
	for _, option := range options {
		option(sr)
	}
	return sr
}

func (sr *NodeMetrics) Print() string {
	return fmt.Sprintf("## job: %s,nodeName: %s,instance: %s,cpuUsage: %f,cpuUsageBefore: %f,memUsage: %f,memUsageBefore: %f", sr.Job, sr.NodeName, sr.Instance, sr.CpuUsage, sr.BeforeCpuUsage, sr.MemUsage, sr.BeforeMemUsage)
}
func (sr *NodeMetrics) ModifyStoreResult(options ...MetricsOption) {
	for _, option := range options {
		option(sr)
	}
}

// 有序输出指标内容
func (sr *NodeMetrics) ConvertToSlice() []string {
	return []string{
		sr.Job,
		sr.NodeName,
		sr.Instance,
		fmt.Sprintf("%f", sr.CpuUsage),
		fmt.Sprintf("%f", sr.BeforeCpuUsage),
		fmt.Sprintf("%f", sr.MemUsage),
		fmt.Sprintf("%f", sr.BeforeMemUsage),
	}
}

type NodeMetricsSlice []*NodeMetrics

func NewNodeMetricsSlice() NodeMetricsSlice {
	return []*NodeMetrics{}
}

func (srs NodeMetricsSlice) findInstance(instance string) (bool, int) {
	for i, sr := range srs {
		if sr.GetInstance() == instance {
			return true, i
		}
	}
	return false, -1
}

// 创建或更新主机指标
func (srs *NodeMetricsSlice) CreateOrModifyStoreResults(instance string, options ...MetricsOption) {
	ok, index := (*srs).findInstance(instance)
	if ok {
		(*srs)[index].ModifyStoreResult(options...)
	} else {
		sr := NewNodeMetrics(instance, options...)
		*srs = append(*srs, sr)
	}
}
