package metrics

import (
	"fmt"
	"node_metrics_go/global"
	am "node_metrics_go/internal/alert"
	rs "node_metrics_go/internal/rules"
	ut "node_metrics_go/internal/utils"

	"go.uber.org/zap"
)

type KafkaMetrics struct {
	*BaseMetrics
	lagSumUsage            float32
	before1DayLagSumUsage  float32
	before1WeekLagSumUsage float32
}

func (b *KafkaMetrics) GetJob() string {
	return b.job
}
func (b *KafkaMetrics) GetInstance() string {
	return b.instance
}

// pre asset
// var a MetricsItf
// var _ = a.(*KafkaMetrics)
// 适配过滤规则
func (sr *KafkaMetrics) AdaptRules(r rs.RuleItf) {
	sr.RuleItf = r
}

// 指标过滤
func (sr *KafkaMetrics) Filter(alertMsgChan chan<- am.AlertInfo) (string, bool) {
	// 若该指标项未匹配到规则
	if sr.RuleItf == nil {
		return "", true
	}

	lagSumInc1Day := ut.IncreaseRate(sr.before1DayLagSumUsage, sr.lagSumUsage)
	lagSumInc1Week := ut.IncreaseRate(sr.before1WeekLagSumUsage, sr.lagSumUsage)

	/* 一周增长率过滤
	 */
	if alertM, ok := rs.WithKafkaLagSumIncrease1WeekRuleFilter(lagSumInc1Week)(sr.RuleItf); !ok {
		alertMsgChan <- am.NewKafkaAlertMessage(
			sr.GetJob(), sr.instance, KAFKA_LAG_SUM_RATE_LIMIT_1WEEK,
			ut.FormatF2S(alertM.(float32)), ut.FormatF2S(lagSumInc1Week),
			ut.FormatF(sr.lagSumUsage), ut.FormatF(sr.before1DayLagSumUsage),
			ut.FormatF(sr.before1WeekLagSumUsage))

		global.Logger.Info(KAFKA_LAG_SUM_RATE_LIMIT_1WEEK, zap.String("job", sr.GetJob()), zap.String("instance", sr.instance), zap.String("kafka_lag_sum_increase_usage_1week", ut.FormatF2S(lagSumInc1Week)))
		return KAFKA_LAG_SUM_RATE_LIMIT_1WEEK, false
	}

	/* 一天增长率过滤
	 */
	if alertM, ok := rs.WithKafkaLagSumIncrease1DayRuleFilter(lagSumInc1Day)(sr.RuleItf); !ok {
		alertMsgChan <- am.NewKafkaAlertMessage(
			sr.GetJob(), sr.instance, KAFKA_LAG_SUM_RATE_LIMIT_1DAY,
			ut.FormatF2S(alertM.(float32)), ut.FormatF2S(lagSumInc1Day),
			ut.FormatF(sr.lagSumUsage), ut.FormatF(sr.before1DayLagSumUsage),
			ut.FormatF(sr.before1WeekLagSumUsage))

		global.Logger.Info(KAFKA_LAG_SUM_RATE_LIMIT_1DAY, zap.String("job", sr.GetJob()), zap.String("instance", sr.instance), zap.String("kafka_lag_sum_increase_usage_1day", ut.FormatF2S(lagSumInc1Day)))
		return KAFKA_LAG_SUM_RATE_LIMIT_1DAY, false
	}

	/* 瞬时值过滤
	 */
	if alertM, ok := rs.WithKafkaLagSumRuleFilter(sr.lagSumUsage)(sr.RuleItf); !ok {
		alertMsgChan <- am.NewKafkaAlertMessage(
			sr.GetJob(), sr.instance, KAFKA_LAG_SUM_LIMIT,
			ut.FormatF(alertM.(float32)), ut.FormatF(sr.lagSumUsage),
			ut.FormatF(sr.lagSumUsage), ut.FormatF(sr.before1DayLagSumUsage),
			ut.FormatF(sr.before1WeekLagSumUsage))

		global.Logger.Info(KAFKA_LAG_SUM_LIMIT, zap.String("job", sr.GetJob()), zap.String("instance", sr.instance), zap.Float32("kafka_lag_sum_usage", sr.lagSumUsage))
		return KAFKA_LAG_SUM_LIMIT, false
	}
	return "", true
}

func (sr *KafkaMetrics) Print() string {
	return fmt.Sprintf("## job: %s,instance: %s,lagSumUsage: %.2f,lagSumUsageBefore1Day: %.2f,lagSumUsageBefore1Week: %.2f", sr.job, sr.instance, sr.lagSumUsage, sr.before1DayLagSumUsage, sr.before1WeekLagSumUsage)
}
func WithKafkaJob(job string) MetricsOption {
	return func(sr MetricsItf) {
		sr.(*KafkaMetrics).job = job
	}
}

func WithKafkaLagSumUsage(lagSumUsage float32) MetricsOption {
	return func(sr MetricsItf) {
		sr.(*KafkaMetrics).lagSumUsage = lagSumUsage
	}
}
func WithBefore1DayKafkaLagSumUsage(beforeLagSumUsage float32) MetricsOption {
	return func(sr MetricsItf) {
		sr.(*KafkaMetrics).before1DayLagSumUsage = beforeLagSumUsage
	}
}
func WithBefore1WeekKafkaLagSumUsage(beforeLagSumUsage float32) MetricsOption {
	return func(sr MetricsItf) {
		sr.(*KafkaMetrics).before1WeekLagSumUsage = beforeLagSumUsage
	}
}

func NewKafkaMetrics(instance string, options ...MetricsOption) *KafkaMetrics {
	mi := &BaseMetrics{instance: instance}
	sr := &KafkaMetrics{
		BaseMetrics: mi,
	}
	for _, option := range options {
		option(sr)
	}
	return sr
}
