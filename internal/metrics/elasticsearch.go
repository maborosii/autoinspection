package metrics

import (
	"fmt"
	"node_metrics_go/global"
	am "node_metrics_go/internal/alert"
	rs "node_metrics_go/internal/rules"
	ut "node_metrics_go/internal/utils"

	"go.uber.org/zap"
)

type ElasticSearchMetrics struct {
	*BaseMetrics
	healthStatus            int8
	before1DayHealthStatus  int8
	before1WeekHealthStatus int8
}

func (em *ElasticSearchMetrics) GetJob() string {
	return em.job
}
func (em *ElasticSearchMetrics) GetInstance() string {
	return em.instance
}

// pre asset
// var a MetricsItf
// var _ = a.(*ElasticSearchMetrics)
// 适配过滤规则
func (em *ElasticSearchMetrics) AdaptRules(r rs.RuleItf) {
	em.RuleItf = r
}

// 指标过滤
func (em *ElasticSearchMetrics) Filter(alertMsgChan chan<- am.AlertInfo) (string, bool) {
	// 若该指标项未匹配到规则
	if em.RuleItf == nil {
		return "", true
	}

	// statusChange1day := em.healthStatus - em.before1DayHealthStatus
	// statusChange1Week := em.healthStatus - em.before1WeekHealthStatus

	/* 健康状态瞬时值过滤
	 */
	if alertM, ok := rs.WithElasticSearchHealthStatusRuleFilter(em.healthStatus)(em.RuleItf); !ok {
		alertMsgChan <- am.NewElasticSearchAlertMessage(
			em.GetJob(), em.instance, ELASTICSEARCH_HEALTH_STATUS,
			ut.FormatD(alertM.(int8)),
			ut.FormatD(em.healthStatus),
			ut.FormatD(em.healthStatus),
			ut.FormatD(em.before1DayHealthStatus),
			ut.FormatD(em.before1WeekHealthStatus))

		global.Logger.Info(
			ELASTICSEARCH_HEALTH_STATUS,
			zap.String("job", em.GetJob()),
			zap.String("instance", em.instance),
			zap.Int8("elasticsearch_health_status", em.healthStatus))
		return ELASTICSEARCH_HEALTH_STATUS, false
	}

	/* 健康状态一天变化过滤
	 */
	// if alertM, ok := rs.WithElasticSearchHealthStatusChange1DayRuleFilter(statusChange1day)(em.RuleItf); !ok {
	// 	alertMsgChan <- am.NewElasticSearchAlertMessage(
	// 		em.GetJob(), em.instance, ELASTICSEARCH_HEALTH_STATUS_CHANGED_1DAY,
	// 		ut.FormatD(alertM.(int8)),
	// 		ut.FormatD(statusChange1day),
	// 		ut.FormatD(em.healthStatus),
	// 		ut.FormatD(em.before1DayHealthStatus),
	// 		ut.FormatD(em.before1WeekHealthStatus))

	// 	global.Logger.Info(
	// 		ELASTICSEARCH_HEALTH_STATUS_CHANGED_1DAY,
	// 		zap.String("job", em.GetJob()),
	// 		zap.String("instance", em.instance),
	// 		zap.Int8("elasticsearch_health_status_changed_1day", statusChange1day))
	// 	return ELASTICSEARCH_HEALTH_STATUS_CHANGED_1DAY, false
	// }

	/* 节点数一周变化过滤
	 */
	// if alertM, ok := rs.WithElasticSearchHealthStatusChange1WeekRuleFilter(statusChange1Week)(em.RuleItf); !ok {
	// 	alertMsgChan <- am.NewElasticSearchAlertMessage(
	// 		em.GetJob(), em.instance, ELASTICSEARCH_HEALTH_STATUS_CHANGED_1WEEK,
	// 		ut.FormatD(alertM.(int8)),
	// 		ut.FormatD(statusChange1Week),
	// 		ut.FormatD(em.healthStatus),
	// 		ut.FormatD(em.before1DayHealthStatus),
	// 		ut.FormatD(em.before1WeekHealthStatus))

	// 	global.Logger.Info(
	// 		ELASTICSEARCH_HEALTH_STATUS_CHANGED_1WEEK,
	// 		zap.String("job", em.GetJob()),
	// 		zap.String("instance", em.instance),
	// 		zap.Int8("elasticsearch_health_status_changed_1week", statusChange1Week))
	// 	return ELASTICSEARCH_HEALTH_STATUS_CHANGED_1WEEK, false
	// }

	return "", true
}

func (em *ElasticSearchMetrics) Print() string {
	return fmt.Sprintf("## job: %s,instance: %s,runnningNodes: %d,runningNodesBefore1Day: %d,runningNodesBefore1Week: %d", em.job, em.instance, em.healthStatus, em.before1DayHealthStatus, em.before1WeekHealthStatus)
}
func WithElasticSearchJob(job string) MetricsOption {
	return func(em MetricsItf) {
		em.(*ElasticSearchMetrics).job = job
	}
}

func WithElasticSearchHeathStatus(healthStatus int8) MetricsOption {
	return func(em MetricsItf) {
		em.(*ElasticSearchMetrics).healthStatus = healthStatus
	}
}
func WithBefore1DayElasticSearchHealthStatus(beforeHealthStatus int8) MetricsOption {
	return func(em MetricsItf) {
		em.(*ElasticSearchMetrics).before1DayHealthStatus = beforeHealthStatus
	}
}
func WithBefore1WeekElasticSearchHealthStatus(beforeHealthStatus int8) MetricsOption {
	return func(em MetricsItf) {
		em.(*ElasticSearchMetrics).before1WeekHealthStatus = beforeHealthStatus
	}
}

func NewElasticSearchMetrics(instance string, options ...MetricsOption) *ElasticSearchMetrics {
	mi := &BaseMetrics{instance: instance}
	em := &ElasticSearchMetrics{
		BaseMetrics: mi,
	}
	for _, option := range options {
		option(em)
	}
	return em
}
