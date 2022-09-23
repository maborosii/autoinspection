package etl

import (
	"fmt"
	"node_metrics_go/global"
	rs "node_metrics_go/pkg/rules"

	"go.uber.org/zap"
)

type NodeMetrics struct {
	*BaseMetrics
	nodeName       string
	cpuUsage       float32
	beforeCpuUsage float32
	memUsage       float32
	beforeMemUsage float32
	// diskUsage        []Disk
}

func (b *NodeMetrics) GetJob() string {
	return b.job
}
func (b *NodeMetrics) GetInstance() string {
	return b.instance
}

// type Disk struct {
// 	mountPoint   string
// 	usage        string
// 	increaseDisk string
// }

// pre asset
// var a MetricsItf
// var _ = a.(*NodeMetrics)
func (sr *NodeMetrics) AdaptRules(r rs.RuleItf) {
	sr.RuleItf = r
}

// 指标过滤
func (sr *NodeMetrics) Filter() (string, bool) {
	// 若该指标项未匹配到规则
	if sr.RuleItf == nil {
		return "", true
	}

	increaseRate := func(a, b float32) float32 {
		return (b - a) / a * 100
	}
	cpuInc := increaseRate(sr.beforeCpuUsage, sr.cpuUsage)
	memInc := increaseRate(sr.beforeMemUsage, sr.memUsage)

	if ok := rs.WithCpuRuleFilter(sr.cpuUsage)(sr.RuleItf); !ok {
		global.Logger.Info("cpu exceeds the threshold", zap.String("job", sr.GetJob()), zap.String("instance", sr.instance), zap.Float32("cpu_usage", sr.cpuUsage))
		return CPU_LIMIT, false
	}
	if ok := rs.WithMemRuleFilter(sr.memUsage)(sr.RuleItf); !ok {
		global.Logger.Info("mem exceeds the threshold", zap.String("job", sr.GetJob()), zap.String("instance", sr.instance), zap.Float32("mem_usage", sr.memUsage))
		return MEM_LIMIT, false
	}
	if ok := rs.WithCpuIncreaseRuleFilter(cpuInc)(sr.RuleItf); !ok {
		global.Logger.Info("cpu rate exceeds the threshold", zap.String("job", sr.GetJob()), zap.String("instance", sr.instance), zap.Float32("cpu_increase_usage", cpuInc))
		return CPU_RATE_LIMIT, false
	}
	if ok := rs.WithMemIncreaseRuleFilter(memInc)(sr.RuleItf); !ok {
		global.Logger.Info("mem rate exceeds the threshold", zap.String("job", sr.GetJob()), zap.String("instance", sr.instance), zap.Float32("mem_increase_usage", memInc))
		return MEM_RATE_LIMIT, false
	}
	return "", true
}

func (sr *NodeMetrics) Print() string {
	return fmt.Sprintf("## job: %s,nodeName: %s,instance: %s,cpuUsage: %f,cpuUsageBefore: %f,memUsage: %f,memUsageBefore: %f", sr.job, sr.nodeName, sr.instance, sr.cpuUsage, sr.beforeCpuUsage, sr.memUsage, sr.beforeMemUsage)
}
func WithNodeJob(job string) MetricsOption {
	return func(sr MetricsItf) {
		sr.(*NodeMetrics).job = job
	}
}
func WithNodeName(nodeName string) MetricsOption {
	return func(sr MetricsItf) {
		sr.(*NodeMetrics).nodeName = nodeName
	}
}
func WithCpuUsage(cpuUsage float32) MetricsOption {
	return func(sr MetricsItf) {
		sr.(*NodeMetrics).cpuUsage = cpuUsage
	}
}
func WithBeforeCpuUsage(beforeCpuUsage float32) MetricsOption {
	return func(sr MetricsItf) {
		sr.(*NodeMetrics).beforeCpuUsage = beforeCpuUsage
	}
}
func WithMemUsage(memUsage float32) MetricsOption {
	return func(sr MetricsItf) {
		sr.(*NodeMetrics).memUsage = memUsage
	}
}
func WithBeforeMemUsage(beforeMemUsage float32) MetricsOption {
	return func(sr MetricsItf) {
		sr.(*NodeMetrics).beforeMemUsage = beforeMemUsage
	}
}
func NewNodeMetrics(instance string, options ...MetricsOption) *NodeMetrics {
	mi := &BaseMetrics{instance: instance}
	sr := &NodeMetrics{
		BaseMetrics: mi,
	}
	for _, option := range options {
		option(sr)
	}
	return sr
}

// 有序输出指标内容
// func (sr *NodeMetrics) ConvertToSlice() []string {
// 	return []string{
// 		sr.Job,
// 		sr.NodeName,
// 		sr.Instance,
// 		fmt.Sprintf("%f", sr.CpuUsage),
// 		fmt.Sprintf("%f", sr.BeforeCpuUsage),
// 		fmt.Sprintf("%f", sr.MemUsage),
// 		fmt.Sprintf("%f", sr.BeforeMemUsage),
// 	}
// }
