package metrics

import (
	"fmt"
	"node_metrics_go/global"
	am "node_metrics_go/internal/alert"
	rs "node_metrics_go/internal/rules"
	ut "node_metrics_go/internal/utils"

	"go.uber.org/zap"
)

type NodeMetrics struct {
	*BaseMetrics
	nodeName                string
	cpuUsage                float32
	before1DayCpuUsage      float32
	before1WeekCpuUsage     float32
	memUsage                float32
	before1DayMemUsage      float32
	before1WeekMemUsage     float32
	diskUsage               float32
	before1DayDiskUsage     float32
	before1WeekDiskUsage    float32
	tcpConnUsage            float32
	before1DayTcpConnUsage  float32
	before1WeekTcpConnUsage float32
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
func (sr *NodeMetrics) Filter(alertMsgChan chan<- am.AlertInfo) (string, bool) {
	// 若该指标项未匹配到规则
	if sr.RuleItf == nil {
		return "", true
	}

	cpuInc1Day := ut.IncreaseRate(sr.before1DayCpuUsage, sr.cpuUsage)
	cpuInc1Week := ut.IncreaseRate(sr.before1WeekCpuUsage, sr.cpuUsage)
	memInc1Day := ut.IncreaseRate(sr.before1DayMemUsage, sr.memUsage)
	memInc1Week := ut.IncreaseRate(sr.before1WeekMemUsage, sr.memUsage)
	diskInc1Day := ut.IncreaseRate(sr.before1DayDiskUsage, sr.diskUsage)
	diskInc1Week := ut.IncreaseRate(sr.before1WeekDiskUsage, sr.diskUsage)
	tcpConnInc1Day := ut.IncreaseRate(sr.before1DayTcpConnUsage, sr.tcpConnUsage)
	tcpConnInc1Week := ut.IncreaseRate(sr.before1WeekTcpConnUsage, sr.tcpConnUsage)

	/* 一周增长率过滤
	 */
	if alertM, ok := rs.WithCpuIncrease1WeekRuleFilter(cpuInc1Week)(sr.RuleItf); !ok {
		alertMsgChan <- am.NewNodeAlertMessage(sr.GetJob(), sr.instance, sr.nodeName, CPU_RATE_LIMIT_1WEEK, fmt.Sprintf("%.2f%%", alertM.(float32)), fmt.Sprintf("%.2f%%", cpuInc1Week))
		global.Logger.Info(CPU_RATE_LIMIT_1WEEK, zap.String("job", sr.GetJob()), zap.String("instance", sr.instance), zap.String("cpu_increase_usage_1week", fmt.Sprintf("%.2f%%", cpuInc1Week)))
		return CPU_RATE_LIMIT_1WEEK, false
	}
	if alertM, ok := rs.WithMemIncrease1WeekRuleFilter(memInc1Week)(sr.RuleItf); !ok {
		alertMsgChan <- am.NewNodeAlertMessage(sr.GetJob(), sr.instance, sr.nodeName, MEM_RATE_LIMIT_1WEEK, fmt.Sprintf("%.2f%%", alertM.(float32)), fmt.Sprintf("%.2f%%", memInc1Week))
		global.Logger.Info(MEM_RATE_LIMIT_1WEEK, zap.String("job", sr.GetJob()), zap.String("instance", sr.instance), zap.String("mem_increase_usage_1week", fmt.Sprintf("%.2f%%", memInc1Week)))
		return MEM_RATE_LIMIT_1WEEK, false
	}
	if alertM, ok := rs.WithTcpConnIncrease1WeekRuleFilter(tcpConnInc1Week)(sr.RuleItf); !ok {
		alertMsgChan <- am.NewNodeAlertMessage(sr.GetJob(), sr.instance, sr.nodeName, TCP_CONN_RATE_LIMIT_1WEEK, fmt.Sprintf("%.2f%%", alertM.(float32)), fmt.Sprintf("%.2f%%", tcpConnInc1Week))
		global.Logger.Info(TCP_CONN_RATE_LIMIT_1WEEK, zap.String("job", sr.GetJob()), zap.String("instance", sr.instance), zap.String("tcp_conn_increase_counts_1week", fmt.Sprintf("%.2f%%", tcpConnInc1Week)))
		return TCP_CONN_RATE_LIMIT_1WEEK, false
	}
	if alertM, ok := rs.WithDiskIncrease1WeekRuleFilter(diskInc1Week)(sr.RuleItf); !ok {
		alertMsgChan <- am.NewNodeAlertMessage(sr.GetJob(), sr.instance, sr.nodeName, DISK_RATE_LIMIT_1WEEK, fmt.Sprintf("%.2f%%", alertM.(float32)), fmt.Sprintf("%.2f%%", diskInc1Week))
		global.Logger.Info(DISK_RATE_LIMIT_1WEEK, zap.String("job", sr.GetJob()), zap.String("instance", sr.instance), zap.String("disk_increase_usage_1week", fmt.Sprintf("%.2f%%", diskInc1Week)))
		return DISK_RATE_LIMIT_1WEEK, false
	}

	/* 一天增长率过滤
	 */
	if alertM, ok := rs.WithCpuIncrease1DayRuleFilter(cpuInc1Day)(sr.RuleItf); !ok {
		alertMsgChan <- am.NewNodeAlertMessage(sr.GetJob(), sr.instance, sr.nodeName, CPU_RATE_LIMIT_1DAY, fmt.Sprintf("%.2f%%", alertM.(float32)), fmt.Sprintf("%.2f%%", cpuInc1Day))
		global.Logger.Info(CPU_RATE_LIMIT_1DAY, zap.String("job", sr.GetJob()), zap.String("instance", sr.instance), zap.String("cpu_increase_usage_1day", fmt.Sprintf("%.2f%%", cpuInc1Day)))
		return CPU_RATE_LIMIT_1DAY, false
	}
	if alertM, ok := rs.WithMemIncrease1DayRuleFilter(memInc1Day)(sr.RuleItf); !ok {
		alertMsgChan <- am.NewNodeAlertMessage(sr.GetJob(), sr.instance, sr.nodeName, MEM_RATE_LIMIT_1DAY, fmt.Sprintf("%.2f%%", alertM.(float32)), fmt.Sprintf("%.2f%%", memInc1Day))
		global.Logger.Info(MEM_RATE_LIMIT_1DAY, zap.String("job", sr.GetJob()), zap.String("instance", sr.instance), zap.String("mem_increase_usage_1day", fmt.Sprintf("%.2f%%", memInc1Day)))
		return MEM_RATE_LIMIT_1DAY, false
	}
	if alertM, ok := rs.WithTcpConnIncrease1DayRuleFilter(tcpConnInc1Day)(sr.RuleItf); !ok {
		alertMsgChan <- am.NewNodeAlertMessage(sr.GetJob(), sr.instance, sr.nodeName, TCP_CONN_RATE_LIMIT_1DAY, fmt.Sprintf("%.2f%%", alertM.(float32)), fmt.Sprintf("%.2f%%", tcpConnInc1Day))
		global.Logger.Info(TCP_CONN_RATE_LIMIT_1DAY, zap.String("job", sr.GetJob()), zap.String("instance", sr.instance), zap.String("tcp_conn_increase_counts_1day", fmt.Sprintf("%.2f%%", tcpConnInc1Day)))
		return TCP_CONN_RATE_LIMIT_1DAY, false
	}
	if alertM, ok := rs.WithDiskIncrease1DayRuleFilter(diskInc1Day)(sr.RuleItf); !ok {
		alertMsgChan <- am.NewNodeAlertMessage(sr.GetJob(), sr.instance, sr.nodeName, DISK_RATE_LIMIT_1DAY, fmt.Sprintf("%.2f%%", alertM.(float32)), fmt.Sprintf("%.2f%%", diskInc1Day))
		global.Logger.Info(DISK_RATE_LIMIT_1DAY, zap.String("job", sr.GetJob()), zap.String("instance", sr.instance), zap.String("disk_increase_usage_1day", fmt.Sprintf("%.2f%%", diskInc1Day)))
		return DISK_RATE_LIMIT_1DAY, false
	}

	/*瞬时值过滤
	 */
	if alertM, ok := rs.WithCpuRuleFilter(sr.cpuUsage)(sr.RuleItf); !ok {
		alertMsgChan <- am.NewNodeAlertMessage(sr.GetJob(), sr.instance, sr.nodeName, CPU_LIMIT, fmt.Sprintf("%.2f%%", alertM.(float32)), fmt.Sprintf("%.2f%%", sr.cpuUsage))
		global.Logger.Info(CPU_LIMIT, zap.String("job", sr.GetJob()), zap.String("instance", sr.instance), zap.String("cpu_usage", fmt.Sprintf("%.2f%%", sr.cpuUsage)))
		return CPU_LIMIT, false
	}
	if alertM, ok := rs.WithMemRuleFilter(sr.memUsage)(sr.RuleItf); !ok {
		alertMsgChan <- am.NewNodeAlertMessage(sr.GetJob(), sr.instance, sr.nodeName, MEM_LIMIT, fmt.Sprintf("%.2f%%", alertM.(float32)), fmt.Sprintf("%.2f%%", sr.memUsage))
		global.Logger.Info("mem exceeds the threshold", zap.String("job", sr.GetJob()), zap.String("instance", sr.instance), zap.String("mem_usage", fmt.Sprintf("%.2f%%", sr.memUsage)))
		return MEM_LIMIT, false
	}
	if alertM, ok := rs.WithTcpConnRuleFilter(sr.tcpConnUsage)(sr.RuleItf); !ok {
		alertMsgChan <- am.NewNodeAlertMessage(sr.GetJob(), sr.instance, sr.nodeName, TCP_CONN_LIMIT, fmt.Sprintf("%.2f", alertM.(float32)), fmt.Sprintf("%.2f", sr.tcpConnUsage))
		global.Logger.Info("tcp conn counts exceeds the threshold", zap.String("job", sr.GetJob()), zap.String("instance", sr.instance), zap.Float32("tcp_conn_counts", sr.tcpConnUsage))
		return TCP_CONN_LIMIT, false
	}
	if alertM, ok := rs.WithDiskRuleFilter(sr.diskUsage)(sr.RuleItf); !ok {
		alertMsgChan <- am.NewNodeAlertMessage(sr.GetJob(), sr.instance, sr.nodeName, DISK_LIMIT, fmt.Sprintf("%.2f%%", alertM.(float32)), fmt.Sprintf("%.2f%%", sr.diskUsage))
		global.Logger.Info("disk exceeds the threshold", zap.String("job", sr.GetJob()), zap.String("instance", sr.instance), zap.String("disk_usage", fmt.Sprintf("%.2f%%", sr.diskUsage)))
		return DISK_LIMIT, false
	}

	return "", true
}

func (sr *NodeMetrics) Print() string {
	return fmt.Sprintf("## job: %s,nodeName: %s,instance: %s,cpuUsage: %.2f%%,cpuUsageBefore1Day: %.2f%%,cpuUsageBefore1Week: %.2f%%,memUsage: %.2f%%,memUsageBefore1Day: %.2f%%,memUsageBefore1Week: %.2f%%,diskUsage: %.2f%%,diskUsageBefore1Day: %.2f%%,diskUsageBefore1Week: %.2f%%,tcpConns: %.2f,tcpConnsBefore1Day: %.2f,tcpConnsBefore1Week: %.2f", sr.job, sr.nodeName, sr.instance, sr.cpuUsage, sr.before1DayCpuUsage, sr.before1WeekCpuUsage, sr.memUsage, sr.before1DayMemUsage, sr.before1WeekMemUsage, sr.diskUsage, sr.before1DayDiskUsage, sr.before1WeekDiskUsage, sr.tcpConnUsage, sr.before1DayTcpConnUsage, sr.before1WeekTcpConnUsage)
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
func WithBefore1DayCpuUsage(beforeCpuUsage float32) MetricsOption {
	return func(sr MetricsItf) {
		sr.(*NodeMetrics).before1DayCpuUsage = beforeCpuUsage
	}
}
func WithBefore1WeekCpuUsage(beforeCpuUsage float32) MetricsOption {
	return func(sr MetricsItf) {
		sr.(*NodeMetrics).before1WeekCpuUsage = beforeCpuUsage
	}
}
func WithMemUsage(memUsage float32) MetricsOption {
	return func(sr MetricsItf) {
		sr.(*NodeMetrics).memUsage = memUsage
	}
}
func WithBefore1DayMemUsage(beforeMemUsage float32) MetricsOption {
	return func(sr MetricsItf) {
		sr.(*NodeMetrics).before1DayMemUsage = beforeMemUsage
	}
}
func WithBefore1WeekMemUsage(beforeMemUsage float32) MetricsOption {
	return func(sr MetricsItf) {
		sr.(*NodeMetrics).before1WeekMemUsage = beforeMemUsage
	}
}
func WithDiskUsage(diskUsage float32) MetricsOption {
	return func(sr MetricsItf) {
		sr.(*NodeMetrics).diskUsage = diskUsage
	}
}
func WithBefore1DayDiskUsage(beforeDiskUsage float32) MetricsOption {
	return func(sr MetricsItf) {
		sr.(*NodeMetrics).before1DayDiskUsage = beforeDiskUsage
	}
}
func WithBefore1WeekDiskUsage(beforeDiskUsage float32) MetricsOption {
	return func(sr MetricsItf) {
		sr.(*NodeMetrics).before1WeekDiskUsage = beforeDiskUsage
	}
}
func WithTcpConnUsage(tcpConnUsage float32) MetricsOption {
	return func(sr MetricsItf) {
		sr.(*NodeMetrics).tcpConnUsage = tcpConnUsage
	}
}
func WithBefore1DayTcpConnUsage(beforeTcpConnUsage float32) MetricsOption {
	return func(sr MetricsItf) {
		sr.(*NodeMetrics).before1DayTcpConnUsage = beforeTcpConnUsage
	}
}
func WithBefore1WeekTcpConnUsage(beforeTcpConnUsage float32) MetricsOption {
	return func(sr MetricsItf) {
		sr.(*NodeMetrics).before1WeekTcpConnUsage = beforeTcpConnUsage
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
