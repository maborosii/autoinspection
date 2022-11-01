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
	before1DayCPUUsage      float32
	before1WeekCPUUsage     float32
	memUsage                float32
	before1DayMemUsage      float32
	before1WeekMemUsage     float32
	diskUsage               float32
	before1DayDiskUsage     float32
	before1WeekDiskUsage    float32
	tcpConnUsage            float32
	before1DayTCPConnUsage  float32
	before1WeekTCPConnUsage float32
}

func (nm *NodeMetrics) GetJob() string {
	return nm.job
}
func (nm *NodeMetrics) GetInstance() string {
	return nm.instance
}

// pre asset
// var a MetricsItf
// var _ = a.(*NodeMetrics)
// 适配过滤规则
func (nm *NodeMetrics) AdaptRules(r rs.RuleItf) {
	nm.RuleItf = r
}

// 指标过滤
func (nm *NodeMetrics) Filter(alertMsgChan chan<- am.AlertInfo) (string, bool) {
	// 若该指标项未匹配到规则
	if nm.RuleItf == nil {
		return "", true
	}

	cpuInc1Day := ut.IncreaseRate(nm.before1DayCPUUsage, nm.cpuUsage)
	cpuInc1Week := ut.IncreaseRate(nm.before1WeekCPUUsage, nm.cpuUsage)
	memInc1Day := ut.IncreaseRate(nm.before1DayMemUsage, nm.memUsage)
	memInc1Week := ut.IncreaseRate(nm.before1WeekMemUsage, nm.memUsage)
	diskInc1Day := ut.IncreaseRate(nm.before1DayDiskUsage, nm.diskUsage)
	diskInc1Week := ut.IncreaseRate(nm.before1WeekDiskUsage, nm.diskUsage)
	tcpConnInc1Day := ut.IncreaseRate(nm.before1DayTCPConnUsage, nm.tcpConnUsage)
	tcpConnInc1Week := ut.IncreaseRate(nm.before1WeekTCPConnUsage, nm.tcpConnUsage)

	/* 一周增长率过滤
	 */
	if alertM, ok := rs.WithCPUIncrease1WeekRuleFilter(cpuInc1Week)(nm.RuleItf); !ok {
		alertMsgChan <- am.NewNodeAlertMessage(
			nm.GetJob(), nm.instance, nm.nodeName, CPU_RATE_LIMIT_1WEEK,
			ut.FormatF2S(alertM.(float32)), ut.FormatF2S(cpuInc1Week),
			ut.FormatF2S(nm.cpuUsage), ut.FormatF2S(nm.before1DayCPUUsage),
			ut.FormatF2S(nm.before1WeekCPUUsage))

		global.Logger.Info(CPU_RATE_LIMIT_1WEEK, zap.String("job", nm.GetJob()), zap.String("instance", nm.instance), zap.String("cpu_increase_usage_1week", ut.FormatF2S(cpuInc1Week)))
		return CPU_RATE_LIMIT_1WEEK, false
	}
	if alertM, ok := rs.WithMemIncrease1WeekRuleFilter(memInc1Week)(nm.RuleItf); !ok {
		alertMsgChan <- am.NewNodeAlertMessage(
			nm.GetJob(), nm.instance, nm.nodeName, MEM_RATE_LIMIT_1WEEK,
			ut.FormatF2S(alertM.(float32)), ut.FormatF2S(memInc1Week),
			ut.FormatF2S(nm.memUsage), ut.FormatF2S(nm.before1DayMemUsage),
			ut.FormatF2S(nm.before1WeekMemUsage))

		global.Logger.Info(MEM_RATE_LIMIT_1WEEK, zap.String("job", nm.GetJob()), zap.String("instance", nm.instance), zap.String("mem_increase_usage_1week", ut.FormatF2S(memInc1Week)))
		return MEM_RATE_LIMIT_1WEEK, false
	}
	if alertM, ok := rs.WithTCPConnIncrease1WeekRuleFilter(tcpConnInc1Week)(nm.RuleItf); !ok {
		alertMsgChan <- am.NewNodeAlertMessage(
			nm.GetJob(), nm.instance, nm.nodeName, TCP_CONN_RATE_LIMIT_1WEEK,
			ut.FormatF2S(alertM.(float32)), ut.FormatF2S(tcpConnInc1Week),
			ut.FormatF(nm.tcpConnUsage), ut.FormatF(nm.before1DayTCPConnUsage),
			ut.FormatF(nm.before1WeekTCPConnUsage))

		global.Logger.Info(TCP_CONN_RATE_LIMIT_1WEEK, zap.String("job", nm.GetJob()), zap.String("instance", nm.instance), zap.String("tcp_conn_increase_counts_1week", ut.FormatF2S(tcpConnInc1Week)))
		return TCP_CONN_RATE_LIMIT_1WEEK, false
	}
	if alertM, ok := rs.WithDiskIncrease1WeekRuleFilter(diskInc1Week)(nm.RuleItf); !ok {
		alertMsgChan <- am.NewNodeAlertMessage(
			nm.GetJob(), nm.instance, nm.nodeName, DISK_RATE_LIMIT_1WEEK,
			ut.FormatF2S(alertM.(float32)), ut.FormatF2S(diskInc1Week),
			ut.FormatF2S(nm.diskUsage), ut.FormatF2S(nm.before1DayDiskUsage),
			ut.FormatF2S(nm.before1WeekDiskUsage))

		global.Logger.Info(DISK_RATE_LIMIT_1WEEK, zap.String("job", nm.GetJob()), zap.String("instance", nm.instance), zap.String("disk_increase_usage_1week", ut.FormatF2S(diskInc1Week)))
		return DISK_RATE_LIMIT_1WEEK, false
	}

	/* 一天增长率过滤
	 */
	if alertM, ok := rs.WithCPUIncrease1DayRuleFilter(cpuInc1Day)(nm.RuleItf); !ok {
		alertMsgChan <- am.NewNodeAlertMessage(
			nm.GetJob(), nm.instance, nm.nodeName, CPU_RATE_LIMIT_1DAY,
			ut.FormatF2S(alertM.(float32)), ut.FormatF2S(cpuInc1Day),
			ut.FormatF2S(nm.cpuUsage), ut.FormatF2S(nm.before1DayCPUUsage),
			ut.FormatF2S(nm.before1WeekCPUUsage))

		global.Logger.Info(CPU_RATE_LIMIT_1DAY, zap.String("job", nm.GetJob()), zap.String("instance", nm.instance), zap.String("cpu_increase_usage_1day", ut.FormatF2S(cpuInc1Day)))
		return CPU_RATE_LIMIT_1DAY, false
	}
	if alertM, ok := rs.WithMemIncrease1DayRuleFilter(memInc1Day)(nm.RuleItf); !ok {
		alertMsgChan <- am.NewNodeAlertMessage(
			nm.GetJob(), nm.instance, nm.nodeName, MEM_RATE_LIMIT_1DAY,
			ut.FormatF2S(alertM.(float32)), ut.FormatF2S(memInc1Day),
			ut.FormatF2S(nm.memUsage), ut.FormatF2S(nm.before1DayMemUsage),
			ut.FormatF2S(nm.before1WeekMemUsage))

		global.Logger.Info(MEM_RATE_LIMIT_1DAY, zap.String("job", nm.GetJob()), zap.String("instance", nm.instance), zap.String("mem_increase_usage_1day", ut.FormatF2S(memInc1Day)))
		return MEM_RATE_LIMIT_1DAY, false
	}
	if alertM, ok := rs.WithTCPConnIncrease1DayRuleFilter(tcpConnInc1Day)(nm.RuleItf); !ok {
		alertMsgChan <- am.NewNodeAlertMessage(
			nm.GetJob(), nm.instance, nm.nodeName, TCP_CONN_RATE_LIMIT_1DAY,
			ut.FormatF2S(alertM.(float32)), ut.FormatF2S(tcpConnInc1Day),
			ut.FormatF(nm.tcpConnUsage), ut.FormatF(nm.before1DayTCPConnUsage),
			ut.FormatF(nm.before1WeekTCPConnUsage))

		global.Logger.Info(TCP_CONN_RATE_LIMIT_1DAY, zap.String("job", nm.GetJob()), zap.String("instance", nm.instance), zap.String("tcp_conn_increase_counts_1day", ut.FormatF2S(tcpConnInc1Day)))
		return TCP_CONN_RATE_LIMIT_1DAY, false
	}
	if alertM, ok := rs.WithDiskIncrease1DayRuleFilter(diskInc1Day)(nm.RuleItf); !ok {
		alertMsgChan <- am.NewNodeAlertMessage(
			nm.GetJob(), nm.instance, nm.nodeName, DISK_RATE_LIMIT_1DAY,
			ut.FormatF2S(alertM.(float32)), ut.FormatF2S(diskInc1Day),
			ut.FormatF2S(nm.diskUsage), ut.FormatF2S(nm.before1DayDiskUsage),
			ut.FormatF2S(nm.before1WeekDiskUsage))

		global.Logger.Info(DISK_RATE_LIMIT_1DAY, zap.String("job", nm.GetJob()), zap.String("instance", nm.instance), zap.String("disk_increase_usage_1day", ut.FormatF2S(diskInc1Day)))
		return DISK_RATE_LIMIT_1DAY, false
	}

	/*瞬时值过滤
	 */
	if alertM, ok := rs.WithCPURuleFilter(nm.cpuUsage)(nm.RuleItf); !ok {
		alertMsgChan <- am.NewNodeAlertMessage(
			nm.GetJob(), nm.instance, nm.nodeName, CPU_LIMIT,
			ut.FormatF2S(alertM.(float32)), ut.FormatF2S(nm.cpuUsage),
			ut.FormatF2S(nm.cpuUsage), ut.FormatF2S(nm.before1DayCPUUsage),
			ut.FormatF2S(nm.before1WeekCPUUsage))

		global.Logger.Info(CPU_LIMIT, zap.String("job", nm.GetJob()), zap.String("instance", nm.instance), zap.String("cpu_usage", ut.FormatF2S(nm.cpuUsage)))
		return CPU_LIMIT, false
	}
	if alertM, ok := rs.WithMemRuleFilter(nm.memUsage)(nm.RuleItf); !ok {
		alertMsgChan <- am.NewNodeAlertMessage(
			nm.GetJob(), nm.instance, nm.nodeName, MEM_LIMIT,
			ut.FormatF2S(alertM.(float32)), ut.FormatF2S(nm.memUsage),
			ut.FormatF2S(nm.memUsage), ut.FormatF2S(nm.before1DayMemUsage),
			ut.FormatF2S(nm.before1WeekMemUsage))

		global.Logger.Info("mem exceeds the threshold", zap.String("job", nm.GetJob()), zap.String("instance", nm.instance), zap.String("mem_usage", ut.FormatF2S(nm.memUsage)))
		return MEM_LIMIT, false
	}
	if alertM, ok := rs.WithTCPConnRuleFilter(nm.tcpConnUsage)(nm.RuleItf); !ok {
		alertMsgChan <- am.NewNodeAlertMessage(
			nm.GetJob(), nm.instance, nm.nodeName, TCP_CONN_LIMIT,
			ut.FormatF(alertM.(float32)), ut.FormatF(nm.tcpConnUsage),
			ut.FormatF(nm.tcpConnUsage), ut.FormatF(nm.before1DayTCPConnUsage),
			ut.FormatF(nm.before1WeekTCPConnUsage))

		global.Logger.Info("tcp conn counts exceeds the threshold", zap.String("job", nm.GetJob()), zap.String("instance", nm.instance), zap.Float32("tcp_conn_counts", nm.tcpConnUsage))
		return TCP_CONN_LIMIT, false
	}
	if alertM, ok := rs.WithDiskRuleFilter(nm.diskUsage)(nm.RuleItf); !ok {
		alertMsgChan <- am.NewNodeAlertMessage(
			nm.GetJob(), nm.instance, nm.nodeName, DISK_LIMIT,
			ut.FormatF2S(alertM.(float32)), ut.FormatF2S(nm.diskUsage),
			ut.FormatF2S(nm.diskUsage), ut.FormatF2S(nm.before1DayDiskUsage),
			ut.FormatF2S(nm.before1WeekDiskUsage))

		global.Logger.Info("disk exceeds the threshold", zap.String("job", nm.GetJob()), zap.String("instance", nm.instance), zap.String("disk_usage", ut.FormatF2S(nm.diskUsage)))
		return DISK_LIMIT, false
	}

	return "", true
}

func (nm *NodeMetrics) Print() string {
	return fmt.Sprintf("## job: %s,nodeName: %s,instance: %s,cpuUsage: %.2f%%,cpuUsageBefore1Day: %.2f%%,cpuUsageBefore1Week: %.2f%%,memUsage: %.2f%%,memUsageBefore1Day: %.2f%%,memUsageBefore1Week: %.2f%%,diskUsage: %.2f%%,diskUsageBefore1Day: %.2f%%,diskUsageBefore1Week: %.2f%%,tcpConns: %.2f,tcpConnsBefore1Day: %.2f,tcpConnsBefore1Week: %.2f", nm.job, nm.nodeName, nm.instance, nm.cpuUsage, nm.before1DayCPUUsage, nm.before1WeekCPUUsage, nm.memUsage, nm.before1DayMemUsage, nm.before1WeekMemUsage, nm.diskUsage, nm.before1DayDiskUsage, nm.before1WeekDiskUsage, nm.tcpConnUsage, nm.before1DayTCPConnUsage, nm.before1WeekTCPConnUsage)
}
func WithNodeJob(job string) MetricsOption {
	return func(nm MetricsItf) {
		nm.(*NodeMetrics).job = job
	}
}
func WithNodeName(nodeName string) MetricsOption {
	return func(nm MetricsItf) {
		nm.(*NodeMetrics).nodeName = nodeName
	}
}
func WithCPUUsage(cpuUsage float32) MetricsOption {
	return func(nm MetricsItf) {
		nm.(*NodeMetrics).cpuUsage = cpuUsage
	}
}
func WithBefore1DayCPUUsage(beforeUsage float32) MetricsOption {
	return func(nm MetricsItf) {
		nm.(*NodeMetrics).before1DayCPUUsage = beforeUsage
	}
}
func WithBefore1WeekCPUUsage(beforeUsage float32) MetricsOption {
	return func(nm MetricsItf) {
		nm.(*NodeMetrics).before1WeekCPUUsage = beforeUsage
	}
}
func WithMemUsage(memUsage float32) MetricsOption {
	return func(nm MetricsItf) {
		nm.(*NodeMetrics).memUsage = memUsage
	}
}
func WithBefore1DayMemUsage(beforeMemUsage float32) MetricsOption {
	return func(nm MetricsItf) {
		nm.(*NodeMetrics).before1DayMemUsage = beforeMemUsage
	}
}
func WithBefore1WeekMemUsage(beforeMemUsage float32) MetricsOption {
	return func(nm MetricsItf) {
		nm.(*NodeMetrics).before1WeekMemUsage = beforeMemUsage
	}
}
func WithDiskUsage(diskUsage float32) MetricsOption {
	return func(nm MetricsItf) {
		nm.(*NodeMetrics).diskUsage = diskUsage
	}
}
func WithBefore1DayDiskUsage(beforeDiskUsage float32) MetricsOption {
	return func(nm MetricsItf) {
		nm.(*NodeMetrics).before1DayDiskUsage = beforeDiskUsage
	}
}
func WithBefore1WeekDiskUsage(beforeDiskUsage float32) MetricsOption {
	return func(nm MetricsItf) {
		nm.(*NodeMetrics).before1WeekDiskUsage = beforeDiskUsage
	}
}
func WithTCPConnUsage(tcpConnUsage float32) MetricsOption {
	return func(nm MetricsItf) {
		nm.(*NodeMetrics).tcpConnUsage = tcpConnUsage
	}
}
func WithBefore1DayTCPConnUsage(beforeUsage float32) MetricsOption {
	return func(nm MetricsItf) {
		nm.(*NodeMetrics).before1DayTCPConnUsage = beforeUsage
	}
}
func WithBefore1WeekTCPConnUsage(beforeUsage float32) MetricsOption {
	return func(nm MetricsItf) {
		nm.(*NodeMetrics).before1WeekTCPConnUsage = beforeUsage
	}
}
func NewNodeMetrics(instance string, options ...MetricsOption) *NodeMetrics {
	mi := &BaseMetrics{instance: instance}
	nm := &NodeMetrics{
		BaseMetrics: mi,
	}
	for _, option := range options {
		option(nm)
	}
	return nm
}
