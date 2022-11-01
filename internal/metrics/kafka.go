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

func (km *KafkaMetrics) GetJob() string {
	return km.job
}
func (km *KafkaMetrics) GetInstance() string {
	return km.instance
}

// pre asset
// var a MetricsItf
// var _ = a.(*KafkaMetrics)
// 适配过滤规则
func (km *KafkaMetrics) AdaptRules(r rs.RuleItf) {
	km.RuleItf = r
}

// 指标过滤
func (km *KafkaMetrics) Filter(alertMsgChan chan<- am.AlertInfo) (string, bool) {
	// 若该指标项未匹配到规则
	if km.RuleItf == nil {
		return "", true
	}

	lagSumInc1Day := ut.IncreaseRate(km.before1DayLagSumUsage, km.lagSumUsage)
	lagSumInc1Week := ut.IncreaseRate(km.before1WeekLagSumUsage, km.lagSumUsage)

	/* 一周增长率过滤
	 */
	if alertM, ok := rs.WithKafkaLagSumIncrease1WeekRuleFilter(lagSumInc1Week)(km.RuleItf); !ok {
		alertMsgChan <- am.NewKafkaAlertMessage(
			km.GetJob(), km.instance, KAFKA_LAG_SUM_RATE_LIMIT_1WEEK,
			ut.FormatF2S(alertM.(float32)), ut.FormatF2S(lagSumInc1Week),
			ut.FormatF(km.lagSumUsage), ut.FormatF(km.before1DayLagSumUsage),
			ut.FormatF(km.before1WeekLagSumUsage))

		global.Logger.Info(KAFKA_LAG_SUM_RATE_LIMIT_1WEEK, zap.String("job", km.GetJob()), zap.String("instance", km.instance), zap.String("kafka_lag_sum_increase_usage_1week", ut.FormatF2S(lagSumInc1Week)))
		return KAFKA_LAG_SUM_RATE_LIMIT_1WEEK, false
	}

	/* 一天增长率过滤
	 */
	if alertM, ok := rs.WithKafkaLagSumIncrease1DayRuleFilter(lagSumInc1Day)(km.RuleItf); !ok {
		alertMsgChan <- am.NewKafkaAlertMessage(
			km.GetJob(), km.instance, KAFKA_LAG_SUM_RATE_LIMIT_1DAY,
			ut.FormatF2S(alertM.(float32)), ut.FormatF2S(lagSumInc1Day),
			ut.FormatF(km.lagSumUsage), ut.FormatF(km.before1DayLagSumUsage),
			ut.FormatF(km.before1WeekLagSumUsage))

		global.Logger.Info(KAFKA_LAG_SUM_RATE_LIMIT_1DAY, zap.String("job", km.GetJob()), zap.String("instance", km.instance), zap.String("kafka_lag_sum_increase_usage_1day", ut.FormatF2S(lagSumInc1Day)))
		return KAFKA_LAG_SUM_RATE_LIMIT_1DAY, false
	}

	/* 瞬时值过滤
	 */
	if alertM, ok := rs.WithKafkaLagSumRuleFilter(km.lagSumUsage)(km.RuleItf); !ok {
		alertMsgChan <- am.NewKafkaAlertMessage(
			km.GetJob(), km.instance, KAFKA_LAG_SUM_LIMIT,
			ut.FormatF(alertM.(float32)), ut.FormatF(km.lagSumUsage),
			ut.FormatF(km.lagSumUsage), ut.FormatF(km.before1DayLagSumUsage),
			ut.FormatF(km.before1WeekLagSumUsage))

		global.Logger.Info(KAFKA_LAG_SUM_LIMIT, zap.String("job", km.GetJob()), zap.String("instance", km.instance), zap.Float32("kafka_lag_sum_usage", km.lagSumUsage))
		return KAFKA_LAG_SUM_LIMIT, false
	}
	return "", true
}

func (km *KafkaMetrics) Print() string {
	return fmt.Sprintf("## job: %s,instance: %s,lagSumUsage: %.2f,lagSumUsageBefore1Day: %.2f,lagSumUsageBefore1Week: %.2f", km.job, km.instance, km.lagSumUsage, km.before1DayLagSumUsage, km.before1WeekLagSumUsage)
}
func WithKafkaJob(job string) MetricsOption {
	return func(km MetricsItf) {
		km.(*KafkaMetrics).job = job
	}
}

func WithKafkaLagSumUsage(lagSumUsage float32) MetricsOption {
	return func(km MetricsItf) {
		km.(*KafkaMetrics).lagSumUsage = lagSumUsage
	}
}
func WithBefore1DayKafkaLagSumUsage(beforeLagSumUsage float32) MetricsOption {
	return func(km MetricsItf) {
		km.(*KafkaMetrics).before1DayLagSumUsage = beforeLagSumUsage
	}
}
func WithBefore1WeekKafkaLagSumUsage(beforeLagSumUsage float32) MetricsOption {
	return func(km MetricsItf) {
		km.(*KafkaMetrics).before1WeekLagSumUsage = beforeLagSumUsage
	}
}

func NewKafkaMetrics(instance string, options ...MetricsOption) *KafkaMetrics {
	mi := &BaseMetrics{instance: instance}
	km := &KafkaMetrics{
		BaseMetrics: mi,
	}
	for _, option := range options {
		option(km)
	}
	return km
}
