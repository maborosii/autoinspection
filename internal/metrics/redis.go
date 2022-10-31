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
	memUsage              float32
	before1DayMemUsage    float32
	before1WeekMemUsage   float32
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
	memInc1Day := ut.IncreaseRate(sr.before1DayMemUsage, sr.memUsage)
	memInc1Week := ut.IncreaseRate(sr.before1WeekMemUsage, sr.memUsage)

	/* 一周增长率过滤
	 */
	if alertM, ok := rs.WithRedisConnsIncrease1WeekRuleFilter(connsInc1Week)(sr.RuleItf); !ok {
		alertMsgChan <- am.NewRedisAlertMessage(
			sr.GetJob(), sr.instance, REDIS_CONN_RATE_LIMIT_1WEEK,
			ut.FormatF2S(alertM.(float32)), ut.FormatF2S(connsInc1Week),
			ut.FormatF(sr.connsUsage), ut.FormatF(sr.before1DayConnsUsage),
			ut.FormatF(sr.before1WeekConnsUsage))

		global.Logger.Info(REDIS_CONN_RATE_LIMIT_1WEEK, zap.String("job", sr.GetJob()), zap.String("instance", sr.instance), zap.String("redis_conns_increase_usage_1week", ut.FormatF2S(connsInc1Week)))
		return REDIS_CONN_RATE_LIMIT_1WEEK, false
	}
	if alertM, ok := rs.WithRedisMemIncrease1WeekRuleFilter(memInc1Week)(sr.RuleItf); !ok {
		alertMsgChan <- am.NewRedisAlertMessage(
			sr.GetJob(), sr.instance, REDIS_MEM_RATE_LIMIT_1WEEK,
			ut.FormatF2S(alertM.(float32)), ut.FormatF2S(memInc1Week),
			ut.FormatF2S(sr.memUsage), ut.FormatF2S(sr.before1DayMemUsage),
			ut.FormatF2S(sr.before1WeekMemUsage))

		global.Logger.Info(REDIS_MEM_RATE_LIMIT_1WEEK, zap.String("job", sr.GetJob()), zap.String("instance", sr.instance), zap.String("redis_memory_increase_usage_1week", ut.FormatF2S(memInc1Week)))
		return REDIS_MEM_RATE_LIMIT_1WEEK, false
	}

	/* 一天增长率过滤
	 */
	if alertM, ok := rs.WithRedisConnsIncrease1DayRuleFilter(connsInc1Day)(sr.RuleItf); !ok {
		alertMsgChan <- am.NewRedisAlertMessage(
			sr.GetJob(), sr.instance, REDIS_CONN_RATE_LIMIT_1DAY,
			ut.FormatF2S(alertM.(float32)), ut.FormatF2S(connsInc1Day),
			ut.FormatF(sr.connsUsage), ut.FormatF(sr.before1DayConnsUsage),
			ut.FormatF(sr.before1WeekConnsUsage))

		global.Logger.Info(REDIS_CONN_RATE_LIMIT_1DAY, zap.String("job", sr.GetJob()), zap.String("instance", sr.instance), zap.String("redis_conns_increase_usage_1day", ut.FormatF2S(connsInc1Day)))
		return REDIS_CONN_RATE_LIMIT_1DAY, false
	}
	if alertM, ok := rs.WithRedisMemIncrease1DayRuleFilter(memInc1Day)(sr.RuleItf); !ok {
		alertMsgChan <- am.NewRedisAlertMessage(
			sr.GetJob(), sr.instance, REDIS_MEM_RATE_LIMIT_1DAY,
			ut.FormatF2S(alertM.(float32)), ut.FormatF2S(memInc1Day),
			ut.FormatF2S(sr.memUsage), ut.FormatF2S(sr.before1DayMemUsage),
			ut.FormatF2S(sr.before1WeekMemUsage))

		global.Logger.Info(REDIS_MEM_RATE_LIMIT_1DAY, zap.String("job", sr.GetJob()), zap.String("instance", sr.instance), zap.String("redis_memory_increase_usage_1day", ut.FormatF2S(memInc1Day)))
		return REDIS_MEM_RATE_LIMIT_1DAY, false
	}

	/* 瞬时值过滤
	 */
	if alertM, ok := rs.WithRedisConnsRuleFilter(sr.connsUsage)(sr.RuleItf); !ok {
		alertMsgChan <- am.NewRedisAlertMessage(
			sr.GetJob(), sr.instance, REDIS_CONN_LIMIT,
			ut.FormatF(alertM.(float32)), ut.FormatF(sr.connsUsage),
			ut.FormatF(sr.connsUsage), ut.FormatF(sr.before1DayConnsUsage),
			ut.FormatF(sr.before1WeekConnsUsage))

		global.Logger.Info(REDIS_CONN_LIMIT, zap.String("job", sr.GetJob()), zap.String("instance", sr.instance), zap.Float32("redis_conns_usage", sr.connsUsage))
		return REDIS_CONN_LIMIT, false
	}
	if alertM, ok := rs.WithRedisMemRuleFilter(sr.memUsage)(sr.RuleItf); !ok {
		alertMsgChan <- am.NewRedisAlertMessage(
			sr.GetJob(), sr.instance, REDIS_MEM_LIMIT,
			ut.FormatF2S(alertM.(float32)), ut.FormatF2S(sr.memUsage),
			ut.FormatF2S(sr.memUsage), ut.FormatF2S(sr.before1DayMemUsage),
			ut.FormatF2S(sr.before1WeekMemUsage))

		global.Logger.Info(REDIS_MEM_LIMIT, zap.String("job", sr.GetJob()), zap.String("instance", sr.instance), zap.String("redis_memory_usage", ut.FormatF2S(sr.memUsage)))
		return REDIS_MEM_LIMIT, false
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

func WithRedisMemUsage(memUsage float32) MetricsOption {
	return func(sr MetricsItf) {
		sr.(*RedisMetrics).memUsage = memUsage
	}
}
func WithBefore1DayRedisMemUsage(beforeMemUsage float32) MetricsOption {
	return func(sr MetricsItf) {
		sr.(*RedisMetrics).before1DayMemUsage = beforeMemUsage
	}
}
func WithBefore1WeekRedisMemUsage(beforeMemUsage float32) MetricsOption {
	return func(sr MetricsItf) {
		sr.(*RedisMetrics).before1WeekMemUsage = beforeMemUsage
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
