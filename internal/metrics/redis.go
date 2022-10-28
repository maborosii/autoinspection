package metrics

import (
	"fmt"
	"node_metrics_go/global"
	am "node_metrics_go/internal/alert"
	rs "node_metrics_go/internal/rules"

	ut "node_metrics_go/internal/utils"

	"go.uber.org/zap"
)

type RedisMetrics struct {
	*BaseMetrics
	connsUsage            float32
	before1DayConnsUsage  float32
	before1WeekConnsUsage float32
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
func (sr *RedisMetrics) Filter(alertMsgChan chan<- am.AlertInfo) (string, bool) {
	// 若该指标项未匹配到规则
	if sr.RuleItf == nil {
		return "", true
	}

	connsInc1Day := ut.IncreaseRate(sr.before1DayConnsUsage, sr.connsUsage)
	connsInc1Week := ut.IncreaseRate(sr.before1WeekConnsUsage, sr.connsUsage)

	/* 一周增长率过滤
	 */
	if alertM, ok := rs.WithRedisConnsIncrease1WeekRuleFilter(connsInc1Week)(sr.RuleItf); !ok {
		alertMsgChan <- am.NewRedisAlertMessage(sr.GetJob(), sr.instance, REDIS_CONN_RATE_LIMIT_1WEEK, fmt.Sprintf("%.2f%%", alertM.(float32)), fmt.Sprintf("%.2f%%", connsInc1Week))
		global.Logger.Info(REDIS_CONN_RATE_LIMIT_1WEEK, zap.String("job", sr.GetJob()), zap.String("instance", sr.instance), zap.String("redis_conns_increase_usage_1week", fmt.Sprintf("%.2f%%", connsInc1Week)))
		return REDIS_CONN_RATE_LIMIT_1WEEK, false
	}

	/* 一天增长率过滤
	 */
	if alertM, ok := rs.WithRedisConnsIncrease1DayRuleFilter(connsInc1Day)(sr.RuleItf); !ok {
		alertMsgChan <- am.NewRedisAlertMessage(sr.GetJob(), sr.instance, REDIS_CONN_RATE_LIMIT_1DAY, fmt.Sprintf("%.2f%%", alertM.(float32)), fmt.Sprintf("%.2f%%", connsInc1Day))
		global.Logger.Info(REDIS_CONN_RATE_LIMIT_1DAY, zap.String("job", sr.GetJob()), zap.String("instance", sr.instance), zap.String("redis_conns_increase_usage_1day", fmt.Sprintf("%.2f%%", connsInc1Day)))
		return REDIS_CONN_RATE_LIMIT_1DAY, false
	}

	/* 瞬时值过滤
	 */
	if alertM, ok := rs.WithRedisConnsRuleFilter(sr.connsUsage)(sr.RuleItf); !ok {
		alertMsgChan <- am.NewRedisAlertMessage(sr.GetJob(), sr.instance, REDIS_CONN_LIMIT, fmt.Sprintf("%.2f", alertM.(float32)), fmt.Sprintf("%.2f", sr.connsUsage))
		global.Logger.Info(REDIS_CONN_LIMIT, zap.String("job", sr.GetJob()), zap.String("instance", sr.instance), zap.Float32("redis_conns_usage", sr.connsUsage))
		return REDIS_CONN_LIMIT, false
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
