package metrics

import (
	"fmt"
	"node_metrics_go/global"
	rs "node_metrics_go/internal/rules"

	"go.uber.org/zap"
)

type RedisMetrics struct {
	*BaseMetrics
	conns                 float32
	before1DayConnsUsage  float32
	before1WeekConnsUsage float32
}

type RedisOutputMessage struct {
	job, instance, alertMessage          string
	alertMetricsLimit, alertMetricsUsage float32
}

func NewRedisOutputMessage(job, instance, alertMessage string, alertMetricsLimit, alertMetricsUsage float32) *NodeOutputMessage {
	return &NodeOutputMessage{
		job:               job,
		instance:          instance,
		alertMessage:      alertMessage,
		alertMetricsLimit: alertMetricsLimit,
		alertMetricsUsage: alertMetricsUsage,
	}
}

func (n *RedisOutputMessage) PrintAlert() string {
	return fmt.Sprintf("Redis 指标异常 >>> job: %s, instance: %s,  告警信息:%s, 当前值:%.2f, 预警值：%.2f\n", n.job, n.instance, n.alertMessage, n.alertMetricsUsage, n.alertMetricsLimit)
}

func (b *RedisMetrics) GetJob() string {
	return b.job
}
func (b *RedisMetrics) GetInstance() string {
	return b.instance
}

// pre asset
// var a MetricsItf
// var _ = a.(*RedisMetrics)
// 适配过滤规则
func (sr *RedisMetrics) AdaptRules(r rs.RuleItf) {
	sr.RuleItf = r
}

// 指标过滤
func (sr *RedisMetrics) Filter(alertMsgChan chan<- AlertInfo) (string, bool) {
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
	connsInc1Day := increaseRate(sr.before1DayConnsUsage, sr.conns)
	connsInc1Week := increaseRate(sr.before1WeekConnsUsage, sr.conns)

	if alertM, ok := rs.WithCpuRuleFilter(sr.cpuUsage)(sr.RuleItf); !ok {
		// nodeOutputMessageList = append(nodeOutputMessageList, NewNodeOutputMessage(sr.GetJob(), sr.instance, sr.nodeName, CPU_LIMIT, alertM.(float32), sr.cpuUsage))
		alertMsgChan <- NewNodeOutputMessage(sr.GetJob(), sr.instance, sr.nodeName, CPU_LIMIT, alertM.(float32), sr.cpuUsage)

		global.Logger.Info(CPU_LIMIT, zap.String("job", sr.GetJob()), zap.String("instance", sr.instance), zap.Float32("cpu_usage", sr.cpuUsage))
		return CPU_LIMIT, false
	}

	if alertM, ok := rs.WithCpuIncrease1DayRuleFilter(cpuInc1Day)(sr.RuleItf); !ok {
		// nodeOutputMessageList = append(nodeOutputMessageList, NewNodeOutputMessage(sr.GetJob(), sr.instance, sr.nodeName, CPU_RATE_LIMIT_1DAY, alertM.(float32), cpuInc1Day))
		alertMsgChan <- NewNodeOutputMessage(sr.GetJob(), sr.instance, sr.nodeName, CPU_RATE_LIMIT_1DAY, alertM.(float32), cpuInc1Day)

		global.Logger.Info(CPU_RATE_LIMIT_1DAY, zap.String("job", sr.GetJob()), zap.String("instance", sr.instance), zap.Float32("cpu_increase_usage_1day", cpuInc1Day))
		return CPU_RATE_LIMIT_1DAY, false
	}

	if alertM, ok := rs.WithCpuIncrease1WeekRuleFilter(cpuInc1Week)(sr.RuleItf); !ok {
		// nodeOutputMessageList = append(nodeOutputMessageList, NewNodeOutputMessage(sr.GetJob(), sr.instance, sr.nodeName, CPU_RATE_LIMIT_1WEEK, alertM.(float32), cpuInc1Week))
		alertMsgChan <- NewNodeOutputMessage(sr.GetJob(), sr.instance, sr.nodeName, CPU_RATE_LIMIT_1WEEK, alertM.(float32), cpuInc1Week)
		global.Logger.Info(CPU_RATE_LIMIT_1WEEK, zap.String("job", sr.GetJob()), zap.String("instance", sr.instance), zap.Float32("cpu_increase_usage_1week", cpuInc1Week))
		return CPU_RATE_LIMIT_1WEEK, false
	}

	if alertM, ok := rs.WithMemRuleFilter(sr.memUsage)(sr.RuleItf); !ok {
		// nodeOutputMessageList = append(nodeOutputMessageList, NewNodeOutputMessage(sr.GetJob(), sr.instance, sr.nodeName, MEM_LIMIT, alertM.(float32), sr.memUsage))
		alertMsgChan <- NewNodeOutputMessage(sr.GetJob(), sr.instance, sr.nodeName, MEM_LIMIT, alertM.(float32), sr.memUsage)

		global.Logger.Info("mem exceeds the threshold", zap.String("job", sr.GetJob()), zap.String("instance", sr.instance), zap.Float32("mem_usage", sr.memUsage))
		return MEM_LIMIT, false
	}

	if alertM, ok := rs.WithMemIncrease1DayRuleFilter(memInc1Day)(sr.RuleItf); !ok {
		// nodeOutputMessageList = append(nodeOutputMessageList, NewNodeOutputMessage(sr.GetJob(), sr.instance, sr.nodeName, MEM_RATE_LIMIT_1DAY, alertM.(float32), memInc1Day))
		alertMsgChan <- NewNodeOutputMessage(sr.GetJob(), sr.instance, sr.nodeName, MEM_RATE_LIMIT_1DAY, alertM.(float32), memInc1Day)

		global.Logger.Info(MEM_RATE_LIMIT_1DAY, zap.String("job", sr.GetJob()), zap.String("instance", sr.instance), zap.Float32("mem_increase_usage_1day", memInc1Day))
		return MEM_RATE_LIMIT_1DAY, false
	}

	if alertM, ok := rs.WithMemIncrease1WeekRuleFilter(memInc1Week)(sr.RuleItf); !ok {
		// nodeOutputMessageList = append(nodeOutputMessageList, NewNodeOutputMessage(sr.GetJob(), sr.instance, sr.nodeName, MEM_RATE_LIMIT_1WEEK, alertM.(float32), memInc1Week))
		alertMsgChan <- NewNodeOutputMessage(sr.GetJob(), sr.instance, sr.nodeName, MEM_RATE_LIMIT_1WEEK, alertM.(float32), memInc1Week)

		global.Logger.Info(MEM_RATE_LIMIT_1WEEK, zap.String("job", sr.GetJob()), zap.String("instance", sr.instance), zap.Float32("mem_increase_usage_1week", memInc1Week))
		return MEM_RATE_LIMIT_1WEEK, false
	}

	if alertM, ok := rs.WithDiskRuleFilter(sr.diskUsage)(sr.RuleItf); !ok {
		// nodeOutputMessageList = append(nodeOutputMessageList, NewNodeOutputMessage(sr.GetJob(), sr.instance, sr.nodeName, DISK_LIMIT, alertM.(float32), sr.diskUsage))
		alertMsgChan <- NewNodeOutputMessage(sr.GetJob(), sr.instance, sr.nodeName, DISK_LIMIT, alertM.(float32), sr.diskUsage)

		global.Logger.Info("disk exceeds the threshold", zap.String("job", sr.GetJob()), zap.String("instance", sr.instance), zap.Float32("disk_usage", sr.diskUsage))
		return DISK_LIMIT, false
	}

	if alertM, ok := rs.WithDiskIncrease1DayRuleFilter(diskInc1Day)(sr.RuleItf); !ok {
		// nodeOutputMessageList = append(nodeOutputMessageList, NewNodeOutputMessage(sr.GetJob(), sr.instance, sr.nodeName, DISK_RATE_LIMIT_1DAY, alertM.(float32), diskInc1Day))
		alertMsgChan <- NewNodeOutputMessage(sr.GetJob(), sr.instance, sr.nodeName, DISK_RATE_LIMIT_1DAY, alertM.(float32), diskInc1Day)

		global.Logger.Info(DISK_RATE_LIMIT_1DAY, zap.String("job", sr.GetJob()), zap.String("instance", sr.instance), zap.Float32("disk_increase_usage_1day", diskInc1Day))
		return DISK_RATE_LIMIT_1DAY, false
	}

	if alertM, ok := rs.WithDiskIncrease1WeekRuleFilter(diskInc1Week)(sr.RuleItf); !ok {
		// nodeOutputMessageList = append(nodeOutputMessageList, NewNodeOutputMessage(sr.GetJob(), sr.instance, sr.nodeName, DISK_RATE_LIMIT_1WEEK, alertM.(float32), diskInc1Week))
		alertMsgChan <- NewNodeOutputMessage(sr.GetJob(), sr.instance, sr.nodeName, DISK_RATE_LIMIT_1WEEK, alertM.(float32), diskInc1Week)

		global.Logger.Info(DISK_RATE_LIMIT_1WEEK, zap.String("job", sr.GetJob()), zap.String("instance", sr.instance), zap.Float32("disk_increase_usage_1week", diskInc1Week))
		return DISK_RATE_LIMIT_1WEEK, false
	}

	if alertM, ok := rs.WithTcpConnRuleFilter(sr.tcpConnUsage)(sr.RuleItf); !ok {
		// nodeOutputMessageList = append(nodeOutputMessageList, NewNodeOutputMessage(sr.GetJob(), sr.instance, sr.nodeName, TCP_CONN_LIMIT, alertM.(float32), sr.tcpConnUsage))
		alertMsgChan <- NewNodeOutputMessage(sr.GetJob(), sr.instance, sr.nodeName, TCP_CONN_LIMIT, alertM.(float32), sr.tcpConnUsage)

		global.Logger.Info("tcp conn counts exceeds the threshold", zap.String("job", sr.GetJob()), zap.String("instance", sr.instance), zap.Float32("tcp_conn_counts", sr.tcpConnUsage))
		return TCP_CONN_LIMIT, false
	}

	if alertM, ok := rs.WithTcpConnIncrease1DayRuleFilter(tcpConnInc1Day)(sr.RuleItf); !ok {
		// nodeOutputMessageList = append(nodeOutputMessageList, NewNodeOutputMessage(sr.GetJob(), sr.instance, sr.nodeName, TCP_CONN_RATE_LIMIT_1DAY, alertM.(float32), tcpConnInc1Day))
		alertMsgChan <- NewNodeOutputMessage(sr.GetJob(), sr.instance, sr.nodeName, TCP_CONN_RATE_LIMIT_1DAY, alertM.(float32), tcpConnInc1Day)

		global.Logger.Info(TCP_CONN_RATE_LIMIT_1DAY, zap.String("job", sr.GetJob()), zap.String("instance", sr.instance), zap.Float32("tcp_conn_increase_counts_1day", tcpConnInc1Day))
		return TCP_CONN_RATE_LIMIT_1DAY, false
	}

	if alertM, ok := rs.WithTcpConnIncrease1WeekRuleFilter(tcpConnInc1Week)(sr.RuleItf); !ok {
		// nodeOutputMessageList = append(nodeOutputMessageList, NewNodeOutputMessage(sr.GetJob(), sr.instance, sr.nodeName, TCP_CONN_RATE_LIMIT_1WEEK, alertM.(float32), tcpConnInc1Week))
		alertMsgChan <- NewNodeOutputMessage(sr.GetJob(), sr.instance, sr.nodeName, TCP_CONN_RATE_LIMIT_1WEEK, alertM.(float32), tcpConnInc1Week)

		global.Logger.Info(TCP_CONN_RATE_LIMIT_1WEEK, zap.String("job", sr.GetJob()), zap.String("instance", sr.instance), zap.Float32("tcp_conn_increase_counts_1day", tcpConnInc1Week))
		return TCP_CONN_RATE_LIMIT_1WEEK, false
	}
	return "", true
}

func (sr *NodeMetrics) Print() string {
	return fmt.Sprintf("## job: %s,nodeName: %s,instance: %s,cpuUsage: %.2f,cpuUsageBefore1Day: %.2f,cpuUsageBefore1Week: %.2f,memUsage: %.2f,memUsageBefore1Day: %.2f,memUsageBefore1Week: %.2f,diskUsage: %.2f,diskUsageBefore1Day: %.2f,diskUsageBefore1Week: %.2f,tcpConns: %.2f,tcpConnsBefore1Day: %.2f,tcpConnsBefore1Week: %.2f", sr.job, sr.nodeName, sr.instance, sr.cpuUsage, sr.before1DayCpuUsage, sr.before1WeekCpuUsage, sr.memUsage, sr.before1DayMemUsage, sr.before1WeekMemUsage, sr.diskUsage, sr.before1DayDiskUsage, sr.before1WeekDiskUsage, sr.tcpConnUsage, sr.before1DayTcpConnUsage, sr.before1WeekTcpConnUsage)
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
