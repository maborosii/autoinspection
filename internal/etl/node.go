package etl

import (
	"fmt"
	"node_metrics_go/global"
	rs "node_metrics_go/pkg/rules"

	"go.uber.org/zap"
)

type NodeMetrics struct {
	*BaseMetrics
	nodeName           string
	cpuUsage           float32
	beforeCpuUsage     float32
	memUsage           float32
	beforeMemUsage     float32
	diskUsage          float32
	beforeDiskUsage    float32
	tcpConnUsage       float32
	beforeTcpConnUsage float32
}

type NodeOutputMessage struct {
	job, instance, nodeName, alertMessage string
	alertMetricsLimit, alertMetricsUsage  float32
}

func NewNodeOutputMessage(job, instance, nodeName, alertMessage string, alertMetricsLimit, alertMetricsUsage float32) *NodeOutputMessage {
	return &NodeOutputMessage{
		job:               job,
		instance:          instance,
		nodeName:          nodeName,
		alertMessage:      alertMessage,
		alertMetricsLimit: alertMetricsLimit,
		alertMetricsUsage: alertMetricsUsage,
	}
}

func (n *NodeOutputMessage) Print() string {
	return fmt.Sprintf("主机指标异常 >>> job: %s, instance: %s, 主机名:%s, 告警信息:%s, 当前值:%f, 预警值：%f\n", n.job, n.instance, n.nodeName, n.alertMessage, n.alertMetricsUsage, n.alertMetricsLimit)
}

func (b *NodeMetrics) GetJob() string {
	return b.job
}
func (b *NodeMetrics) GetInstance() string {
	return b.instance
}

// pre asset
// var a MetricsItf
// var _ = a.(*NodeMetrics)
// 适配过滤规则
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
		if a == 0 {
			return 0
		}
		return (b - a) / a * 100
	}
	cpuInc := increaseRate(sr.beforeCpuUsage, sr.cpuUsage)
	memInc := increaseRate(sr.beforeMemUsage, sr.memUsage)
	diskInc := increaseRate(sr.beforeDiskUsage, sr.diskUsage)
	tcpConnInc := increaseRate(sr.beforeTcpConnUsage, sr.tcpConnUsage)

	if alertM, ok := rs.WithCpuRuleFilter(sr.cpuUsage)(sr.RuleItf); !ok {
		nodeOutputMessageList = append(nodeOutputMessageList, NewNodeOutputMessage(sr.GetJob(), sr.instance, sr.nodeName, CPU_LIMIT, alertM.(float32), sr.cpuUsage))

		global.Logger.Info("cpu exceeds the threshold", zap.String("job", sr.GetJob()), zap.String("instance", sr.instance), zap.Float32("cpu_usage", sr.cpuUsage))
		return CPU_LIMIT, false
	}
	if alertM, ok := rs.WithCpuIncreaseRuleFilter(cpuInc)(sr.RuleItf); !ok {
		nodeOutputMessageList = append(nodeOutputMessageList, NewNodeOutputMessage(sr.GetJob(), sr.instance, sr.nodeName, CPU_RATE_LIMIT, alertM.(float32), cpuInc))

		global.Logger.Info("cpu rate exceeds the threshold", zap.String("job", sr.GetJob()), zap.String("instance", sr.instance), zap.Float32("cpu_increase_usage", cpuInc))
		return CPU_RATE_LIMIT, false
	}
	if alertM, ok := rs.WithMemRuleFilter(sr.memUsage)(sr.RuleItf); !ok {
		nodeOutputMessageList = append(nodeOutputMessageList, NewNodeOutputMessage(sr.GetJob(), sr.instance, sr.nodeName, MEM_LIMIT, alertM.(float32), sr.memUsage))

		global.Logger.Info("mem exceeds the threshold", zap.String("job", sr.GetJob()), zap.String("instance", sr.instance), zap.Float32("mem_usage", sr.memUsage))
		return MEM_LIMIT, false
	}
	if alertM, ok := rs.WithMemIncreaseRuleFilter(memInc)(sr.RuleItf); !ok {
		nodeOutputMessageList = append(nodeOutputMessageList, NewNodeOutputMessage(sr.GetJob(), sr.instance, sr.nodeName, MEM_RATE_LIMIT, alertM.(float32), memInc))

		global.Logger.Info("mem rate exceeds the threshold", zap.String("job", sr.GetJob()), zap.String("instance", sr.instance), zap.Float32("mem_increase_usage", memInc))
		return MEM_RATE_LIMIT, false
	}
	if alertM, ok := rs.WithDiskRuleFilter(sr.diskUsage)(sr.RuleItf); !ok {
		nodeOutputMessageList = append(nodeOutputMessageList, NewNodeOutputMessage(sr.GetJob(), sr.instance, sr.nodeName, DISK_LIMIT, alertM.(float32), sr.diskUsage))

		global.Logger.Info("disk exceeds the threshold", zap.String("job", sr.GetJob()), zap.String("instance", sr.instance), zap.Float32("disk_usage", sr.diskUsage))
		return DISK_LIMIT, false
	}
	if alertM, ok := rs.WithDiskIncreaseRuleFilter(diskInc)(sr.RuleItf); !ok {
		nodeOutputMessageList = append(nodeOutputMessageList, NewNodeOutputMessage(sr.GetJob(), sr.instance, sr.nodeName, DISK_RATE_LIMIT, alertM.(float32), diskInc))

		global.Logger.Info("disk rate exceeds the threshold", zap.String("job", sr.GetJob()), zap.String("instance", sr.instance), zap.Float32("disk_increase_usage", diskInc))
		return DISK_RATE_LIMIT, false
	}
	if alertM, ok := rs.WithTcpConnRuleFilter(sr.tcpConnUsage)(sr.RuleItf); !ok {
		nodeOutputMessageList = append(nodeOutputMessageList, NewNodeOutputMessage(sr.GetJob(), sr.instance, sr.nodeName, TCP_CONN_LIMIT, alertM.(float32), sr.tcpConnUsage))

		global.Logger.Info("tcp conn counts exceeds the threshold", zap.String("job", sr.GetJob()), zap.String("instance", sr.instance), zap.Float32("tcp_conn_counts", sr.tcpConnUsage))
		return TCP_CONN_LIMIT, false
	}
	if alertM, ok := rs.WithTcpConnIncreaseRuleFilter(tcpConnInc)(sr.RuleItf); !ok {
		nodeOutputMessageList = append(nodeOutputMessageList, NewNodeOutputMessage(sr.GetJob(), sr.instance, sr.nodeName, TCP_CONN_RATE_LIMIT, alertM.(float32), tcpConnInc))

		global.Logger.Info("tcp conn counts rate exceeds the threshold", zap.String("job", sr.GetJob()), zap.String("instance", sr.instance), zap.Float32("tcp_conn_increase_counts", tcpConnInc))
		return TCP_CONN_RATE_LIMIT, false
	}
	return "", true
}

func (sr *NodeMetrics) Print() string {
	return fmt.Sprintf("## job: %s,nodeName: %s,instance: %s,cpuUsage: %f,cpuUsageBefore: %f,memUsage: %f,memUsageBefore: %f,diskUsage: %f,diskUsageBefore: %f,tcpConns: %f,tcpConnsBefore: %f", sr.job, sr.nodeName, sr.instance, sr.cpuUsage, sr.beforeCpuUsage, sr.memUsage, sr.beforeMemUsage, sr.diskUsage, sr.beforeDiskUsage, sr.tcpConnUsage, sr.beforeTcpConnUsage)
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
func WithDiskUsage(diskUsage float32) MetricsOption {
	return func(sr MetricsItf) {
		sr.(*NodeMetrics).diskUsage = diskUsage
	}
}
func WithBeforeDiskUsage(beforeDiskUsage float32) MetricsOption {
	return func(sr MetricsItf) {
		sr.(*NodeMetrics).beforeDiskUsage = beforeDiskUsage
	}
}
func WithTcpConnUsage(tcpConnUsage float32) MetricsOption {
	return func(sr MetricsItf) {
		sr.(*NodeMetrics).tcpConnUsage = tcpConnUsage
	}
}
func WithBeforeTcpConnUsage(beforeTcpConnUsage float32) MetricsOption {
	return func(sr MetricsItf) {
		sr.(*NodeMetrics).beforeTcpConnUsage = beforeTcpConnUsage
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
