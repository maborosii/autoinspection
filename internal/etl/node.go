package etl

import "fmt"

type NodeMetrics struct {
	*MetricsInfo
	NodeName       string
	CpuUsage       string
	BeforeCpuUsage string
	MemUsage       string
	BeforeMemUsage string
	// diskUsage        []Disk
}

// type Disk struct {
// 	mountPoint   string
// 	usage        string
// 	increaseDisk string
// }

type NotifyRule struct {
	job string
}

type NodeOption func(*NodeMetrics)

func WithNodeJob(job string) NodeOption {
	return func(sr *NodeMetrics) {
		sr.Job = job
	}
}
func WithNodeName(nodeName string) NodeOption {
	return func(sr *NodeMetrics) {
		sr.NodeName = nodeName
	}
}

func WithCpuUsage(cpuUsage string) NodeOption {
	return func(sr *NodeMetrics) {
		sr.CpuUsage = cpuUsage
	}
}
func WithBeforeCpuUsage(beforeCpuUsage string) NodeOption {
	return func(sr *NodeMetrics) {
		sr.BeforeCpuUsage = beforeCpuUsage
	}
}

func WithMemUsage(memUsage string) NodeOption {
	return func(sr *NodeMetrics) {
		sr.MemUsage = memUsage
	}
}

func WithBeforeMemUsage(beforeMemUsage string) NodeOption {
	return func(sr *NodeMetrics) {
		sr.BeforeMemUsage = beforeMemUsage
	}
}

func NewNodeMetrics(instance string, options ...NodeOption) *NodeMetrics {
	mi := &MetricsInfo{Instance: instance}
	sr := &NodeMetrics{
		MetricsInfo: mi,
	}
	for _, option := range options {
		option(sr)
	}
	return sr
}

func (sr *MetricsInfo) GetInstance() string {
	return sr.Instance
}
func (sr *MetricsInfo) GetJob() string {
	return sr.Job
}
func (sr *NodeMetrics) Print() string {
	return fmt.Sprintf("## job: %s,nodeName: %s,instance: %s,cpuUsage: %s,cpuUsageBefore: %s,memUsage: %s,memUsageBefore: %s", sr.Job, sr.NodeName, sr.Instance, sr.CpuUsage, sr.BeforeCpuUsage, sr.MemUsage, sr.BeforeMemUsage)
}
func (sr *NodeMetrics) ModifyStoreResult(options ...NodeOption) {
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
		sr.CpuUsage,
		sr.BeforeCpuUsage,
		sr.MemUsage,
		sr.BeforeMemUsage,
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
func (srs *NodeMetricsSlice) CreateOrModifyStoreResults(instance string, options ...NodeOption) {
	ok, index := (*srs).findInstance(instance)
	if ok {
		(*srs)[index].ModifyStoreResult(options...)
	} else {
		sr := NewNodeMetrics(instance, options...)
		*srs = append(*srs, sr)
	}
}
