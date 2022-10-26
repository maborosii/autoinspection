package metrics

import (
	"fmt"
	"node_metrics_go/global"
	rs "node_metrics_go/internal/rules"

	"github.com/jedib0t/go-pretty/v6/table"
	"go.uber.org/zap"
)

type RedisMetrics struct {
	*BaseMetrics
	connsUsage            float32
	before1DayConnsUsage  float32
	before1WeekConnsUsage float32
}

type RedisOutputMessage struct {
	job, instance, alertMessage          string
	alertMetricsLimit, alertMetricsUsage float32
}

func NewRedisOutputMessage(job, instance, alertMessage string, alertMetricsLimit, alertMetricsUsage float32) *RedisOutputMessage {
	return &RedisOutputMessage{
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
func (n *RedisOutputMessage) PrintAlertFormatTable() table.Row {
	return table.Row{"Redis 指标异常", n.job, n.instance, "", n.alertMessage, n.alertMetricsUsage, n.alertMetricsLimit}
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
	connsInc1Day := increaseRate(sr.before1DayConnsUsage, sr.connsUsage)
	connsInc1Week := increaseRate(sr.before1WeekConnsUsage, sr.connsUsage)

	if alertM, ok := rs.WithRedisConnsRuleFilter(sr.connsUsage)(sr.RuleItf); !ok {
		alertMsgChan <- NewRedisOutputMessage(sr.GetJob(), sr.instance, REDIS_CONN_LIMIT, alertM.(float32), sr.connsUsage)
		global.Logger.Info(REDIS_CONN_LIMIT, zap.String("job", sr.GetJob()), zap.String("instance", sr.instance), zap.Float32("redis_conns_usage", sr.connsUsage))
		return REDIS_CONN_LIMIT, false
	}

	if alertM, ok := rs.WithRedisConnsIncrease1DayRuleFilter(connsInc1Day)(sr.RuleItf); !ok {
		alertMsgChan <- NewRedisOutputMessage(sr.GetJob(), sr.instance, REDIS_CONN_RATE_LIMIT_1DAY, alertM.(float32), connsInc1Day)
		global.Logger.Info(REDIS_CONN_RATE_LIMIT_1DAY, zap.String("job", sr.GetJob()), zap.String("instance", sr.instance), zap.Float32("redis_conns_increase_usage_1day", connsInc1Day))
		return REDIS_CONN_RATE_LIMIT_1DAY, false
	}

	if alertM, ok := rs.WithRedisConnsIncrease1WeekRuleFilter(connsInc1Week)(sr.RuleItf); !ok {
		alertMsgChan <- NewRedisOutputMessage(sr.GetJob(), sr.instance, REDIS_CONN_RATE_LIMIT_1WEEK, alertM.(float32), connsInc1Week)
		global.Logger.Info(REDIS_CONN_RATE_LIMIT_1WEEK, zap.String("job", sr.GetJob()), zap.String("instance", sr.instance), zap.Float32("redis_conns_increase_usage_1week", connsInc1Week))
		return REDIS_CONN_RATE_LIMIT_1WEEK, false
	}

	return "", true
}

func (sr *RedisMetrics) Print() string {
	return fmt.Sprintf("## job: %s,instance: %s,connsUsage: %.2f,connsUsageBefore1Day: %.2f,connsUsageBefore1Week: %.2f", sr.job, sr.instance, sr.connsUsage, sr.before1DayConnsUsage, sr.before1WeekConnsUsage)
}
func WithRedisJob(job string) MetricsOption {
	return func(sr MetricsItf) {
		sr.(*RedisMetrics).job = job
	}
}

func WithRedisConnsUsage(connsUsage float32) MetricsOption {
	return func(sr MetricsItf) {
		sr.(*RedisMetrics).connsUsage = connsUsage
	}
}
func WithBefore1DayRedisConnsUsage(beforeConnsUsage float32) MetricsOption {
	return func(sr MetricsItf) {
		sr.(*RedisMetrics).before1DayConnsUsage = beforeConnsUsage
	}
}
func WithBefore1WeekRedisConnsUsage(beforeConnsUsage float32) MetricsOption {
	return func(sr MetricsItf) {
		sr.(*RedisMetrics).before1WeekConnsUsage = beforeConnsUsage
	}
}

func NewRedisMetrics(instance string, options ...MetricsOption) *RedisMetrics {
	mi := &BaseMetrics{instance: instance}
	sr := &RedisMetrics{
		BaseMetrics: mi,
	}
	for _, option := range options {
		option(sr)
	}
	return sr
}
