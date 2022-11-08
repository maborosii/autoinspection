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

func (rm *RedisMetrics) GetJob() string {
	return rm.job
}
func (rm *RedisMetrics) GetInstance() string {
	return rm.instance
}

// pre asset
// var a MetricsItf
// var _ = a.(*RedisMetrics)
// 适配过滤规则
func (rm *RedisMetrics) AdaptRules(r rs.RuleItf) {
	rm.RuleItf = r
}

// 指标过滤
func (rm *RedisMetrics) Filter(alertMsgChan chan<- am.AlertInfo) (string, bool) {
	// 若该指标项未匹配到规则
	if rm.RuleItf == nil {
		return "", true
	}

	connsInc1Day := ut.IncreaseRate(rm.before1DayConnsUsage, rm.connsUsage)
	connsInc1Week := ut.IncreaseRate(rm.before1WeekConnsUsage, rm.connsUsage)
	memInc1Day := ut.IncreaseRate(rm.before1DayMemUsage, rm.memUsage)
	memInc1Week := ut.IncreaseRate(rm.before1WeekMemUsage, rm.memUsage)

	/* 一周增长率过滤
	 */
	if alertM, ok := rs.WithRedisConnsIncrease1WeekRuleFilter(connsInc1Week)(rm.RuleItf); !ok {
		alertMsgChan <- am.NewRedisAlertMessage(
			rm.GetJob(), rm.instance, REDIS_CONN_RATE_LIMIT_1WEEK,
			ut.FormatF2S(alertM.(float32)),
			ut.FormatF2S(connsInc1Week),
			ut.FormatF(rm.connsUsage),
			ut.FormatF(rm.before1DayConnsUsage),
			ut.FormatF(rm.before1WeekConnsUsage))

		global.Logger.Info(
			REDIS_CONN_RATE_LIMIT_1WEEK,
			zap.String("job", rm.GetJob()),
			zap.String("instance", rm.instance),
			zap.String("redis_conns_increase_usage_1week", ut.FormatF2S(connsInc1Week)))
		return REDIS_CONN_RATE_LIMIT_1WEEK, false
	}
	if alertM, ok := rs.WithRedisMemIncrease1WeekRuleFilter(memInc1Week)(rm.RuleItf); !ok {
		alertMsgChan <- am.NewRedisAlertMessage(
			rm.GetJob(), rm.instance, REDIS_MEM_RATE_LIMIT_1WEEK,
			ut.FormatF2S(alertM.(float32)),
			ut.FormatF2S(memInc1Week),
			ut.FormatF2S(rm.memUsage),
			ut.FormatF2S(rm.before1DayMemUsage),
			ut.FormatF2S(rm.before1WeekMemUsage))

		global.Logger.Info(
			REDIS_MEM_RATE_LIMIT_1WEEK,
			zap.String("job", rm.GetJob()),
			zap.String("instance", rm.instance),
			zap.String("redis_memory_increase_usage_1week", ut.FormatF2S(memInc1Week)))
		return REDIS_MEM_RATE_LIMIT_1WEEK, false
	}

	/* 一天增长率过滤
	 */
	if alertM, ok := rs.WithRedisConnsIncrease1DayRuleFilter(connsInc1Day)(rm.RuleItf); !ok {
		alertMsgChan <- am.NewRedisAlertMessage(
			rm.GetJob(), rm.instance, REDIS_CONN_RATE_LIMIT_1DAY,
			ut.FormatF2S(alertM.(float32)),
			ut.FormatF2S(connsInc1Day),
			ut.FormatF(rm.connsUsage),
			ut.FormatF(rm.before1DayConnsUsage),
			ut.FormatF(rm.before1WeekConnsUsage))

		global.Logger.Info(
			REDIS_CONN_RATE_LIMIT_1DAY,
			zap.String("job", rm.GetJob()),
			zap.String("instance", rm.instance),
			zap.String("redis_conns_increase_usage_1day", ut.FormatF2S(connsInc1Day)))
		return REDIS_CONN_RATE_LIMIT_1DAY, false
	}
	if alertM, ok := rs.WithRedisMemIncrease1DayRuleFilter(memInc1Day)(rm.RuleItf); !ok {
		alertMsgChan <- am.NewRedisAlertMessage(
			rm.GetJob(), rm.instance, REDIS_MEM_RATE_LIMIT_1DAY,
			ut.FormatF2S(alertM.(float32)),
			ut.FormatF2S(memInc1Day),
			ut.FormatF2S(rm.memUsage),
			ut.FormatF2S(rm.before1DayMemUsage),
			ut.FormatF2S(rm.before1WeekMemUsage))

		global.Logger.Info(
			REDIS_MEM_RATE_LIMIT_1DAY,
			zap.String("job", rm.GetJob()),
			zap.String("instance", rm.instance),
			zap.String("redis_memory_increase_usage_1day", ut.FormatF2S(memInc1Day)))
		return REDIS_MEM_RATE_LIMIT_1DAY, false
	}

	/* 瞬时值过滤
	 */
	if alertM, ok := rs.WithRedisConnsRuleFilter(rm.connsUsage)(rm.RuleItf); !ok {
		alertMsgChan <- am.NewRedisAlertMessage(
			rm.GetJob(), rm.instance, REDIS_CONN_LIMIT,
			ut.FormatF(alertM.(float32)),
			ut.FormatF(rm.connsUsage),
			ut.FormatF(rm.connsUsage),
			ut.FormatF(rm.before1DayConnsUsage),
			ut.FormatF(rm.before1WeekConnsUsage))

		global.Logger.Info(
			REDIS_CONN_LIMIT,
			zap.String("job", rm.GetJob()),
			zap.String("instance", rm.instance),
			zap.Float32("redis_conns_usage", rm.connsUsage))
		return REDIS_CONN_LIMIT, false
	}
	if alertM, ok := rs.WithRedisMemRuleFilter(rm.memUsage)(rm.RuleItf); !ok {
		alertMsgChan <- am.NewRedisAlertMessage(
			rm.GetJob(), rm.instance, REDIS_MEM_LIMIT,
			ut.FormatF2S(alertM.(float32)),
			ut.FormatF2S(rm.memUsage),
			ut.FormatF2S(rm.memUsage),
			ut.FormatF2S(rm.before1DayMemUsage),
			ut.FormatF2S(rm.before1WeekMemUsage))

		global.Logger.Info(
			REDIS_MEM_LIMIT,
			zap.String("job", rm.GetJob()),
			zap.String("instance", rm.instance),
			zap.String("redis_memory_usage", ut.FormatF2S(rm.memUsage)))
		return REDIS_MEM_LIMIT, false
	}

	return "", true
}

func (rm *RedisMetrics) Print() string {
	return fmt.Sprintf("## job: %s,instance: %s,connsUsage: %.2f,connsUsageBefore1Day: %.2f,connsUsageBefore1Week: %.2f", rm.job, rm.instance, rm.connsUsage, rm.before1DayConnsUsage, rm.before1WeekConnsUsage)
}
func WithRedisJob(job string) MetricsOption {
	return func(rm MetricsItf) {
		rm.(*RedisMetrics).job = job
	}
}

func WithRedisConnsUsage(connsUsage float32) MetricsOption {
	return func(rm MetricsItf) {
		rm.(*RedisMetrics).connsUsage = connsUsage
	}
}
func WithBefore1DayRedisConnsUsage(beforeConnsUsage float32) MetricsOption {
	return func(rm MetricsItf) {
		rm.(*RedisMetrics).before1DayConnsUsage = beforeConnsUsage
	}
}
func WithBefore1WeekRedisConnsUsage(beforeConnsUsage float32) MetricsOption {
	return func(rm MetricsItf) {
		rm.(*RedisMetrics).before1WeekConnsUsage = beforeConnsUsage
	}
}

func WithRedisMemUsage(memUsage float32) MetricsOption {
	return func(rm MetricsItf) {
		rm.(*RedisMetrics).memUsage = memUsage
	}
}
func WithBefore1DayRedisMemUsage(beforeMemUsage float32) MetricsOption {
	return func(rm MetricsItf) {
		rm.(*RedisMetrics).before1DayMemUsage = beforeMemUsage
	}
}
func WithBefore1WeekRedisMemUsage(beforeMemUsage float32) MetricsOption {
	return func(rm MetricsItf) {
		rm.(*RedisMetrics).before1WeekMemUsage = beforeMemUsage
	}
}

func NewRedisMetrics(instance string, options ...MetricsOption) *RedisMetrics {
	mi := &BaseMetrics{instance: instance}
	rm := &RedisMetrics{
		BaseMetrics: mi,
	}
	for _, option := range options {
		option(rm)
	}
	return rm
}
