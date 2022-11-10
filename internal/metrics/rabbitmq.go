package metrics

import (
	"fmt"
	"node_metrics_go/global"
	am "node_metrics_go/internal/alert"
	rs "node_metrics_go/internal/rules"
	ut "node_metrics_go/internal/utils"

	"go.uber.org/zap"
)

type RabbitMQMetrics struct {
	*BaseMetrics
	runningNodesUsage            int8
	before1DayRunningNodesUsage  int8
	before1WeekRunningNodesUsage int8
	lagSumUsage                  float32
	before1DayLagSumUsage        float32
	before1WeekLagSumUsage       float32
}

func (rm *RabbitMQMetrics) GetJob() string {
	return rm.job
}
func (rm *RabbitMQMetrics) GetInstance() string {
	return rm.instance
}

// pre asset
// var a MetricsItf
// var _ = a.(*RabbitMQMetrics)
// 适配过滤规则
func (rm *RabbitMQMetrics) AdaptRules(r rs.RuleItf) {
	rm.RuleItf = r
}

// 指标过滤
func (rm *RabbitMQMetrics) Filter(alertMsgChan chan<- am.AlertInfo) (string, bool) {
	// 若该指标项未匹配到规则
	if rm.RuleItf == nil {
		return "", true
	}

	// nodeChange1day := rm.runningNodesUsage - rm.before1DayRunningNodesUsage
	// nodeChange1Week := rm.runningNodesUsage - rm.before1WeekRunningNodesUsage
	lagSumInc1Day := ut.IncreaseRate(rm.before1DayLagSumUsage, rm.lagSumUsage)
	lagSumInc1Week := ut.IncreaseRate(rm.before1WeekLagSumUsage, rm.lagSumUsage)

	/* 节点数瞬时值过滤
	 */
	if alertM, ok := rs.WithRabbitMQRunningNodesRuleFilter(rm.runningNodesUsage)(rm.RuleItf); !ok {
		alertMsgChan <- am.NewRabbitMQAlertMessage(
			rm.GetJob(), rm.instance, RABBITMQ_RUNNING_NODES,
			ut.FormatD(alertM.(int8)),
			ut.FormatD(rm.runningNodesUsage),
			ut.FormatD(rm.runningNodesUsage),
			ut.FormatD(rm.before1DayRunningNodesUsage),
			ut.FormatD(rm.before1WeekRunningNodesUsage))

		global.Logger.Info(
			RABBITMQ_RUNNING_NODES,
			zap.String("job", rm.GetJob()),
			zap.String("instance", rm.instance),
			zap.Int8("rabbitmq_running_nodes", rm.runningNodesUsage))
		return RABBITMQ_RUNNING_NODES, false
	}

	/* 节点数一天变化过滤
	 */
	// if alertM, ok := rs.WithRabbitMQRunningNodesChange1DayRuleFilter(nodeChange1day)(rm.RuleItf); !ok {
	// 	alertMsgChan <- am.NewRabbitMQAlertMessage(
	// 		rm.GetJob(), rm.instance, RABBITMQ_RUNNING_NODES_CHANGED_1DAY,
	// 		ut.FormatD(alertM.(int8)),
	// 		ut.FormatD(nodeChange1day),
	// 		ut.FormatD(rm.runningNodesUsage),
	// 		ut.FormatD(rm.before1DayRunningNodesUsage),
	// 		ut.FormatD(rm.before1WeekRunningNodesUsage))

	// 	global.Logger.Info(
	// 		RABBITMQ_RUNNING_NODES_CHANGED_1DAY,
	// 		zap.String("job", rm.GetJob()),
	// 		zap.String("instance", rm.instance),
	// 		zap.Int8("rabbitmq_running_nodes_changed_1day", nodeChange1day))
	// 	return RABBITMQ_RUNNING_NODES_CHANGED_1DAY, false
	// }

	/* 节点数一周变化过滤
	 */
	// if alertM, ok := rs.WithRabbitMQRunningNodesChange1WeekRuleFilter(nodeChange1Week)(rm.RuleItf); !ok {
	// 	alertMsgChan <- am.NewRabbitMQAlertMessage(
	// 		rm.GetJob(), rm.instance, RABBITMQ_RUNNING_NODES_CHANGED_1WEEK,
	// 		ut.FormatD(alertM.(int8)),
	// 		ut.FormatD(nodeChange1Week),
	// 		ut.FormatD(rm.runningNodesUsage),
	// 		ut.FormatD(rm.before1DayRunningNodesUsage),
	// 		ut.FormatD(rm.before1WeekRunningNodesUsage))

	// 	global.Logger.Info(
	// 		RABBITMQ_RUNNING_NODES_CHANGED_1WEEK,
	// 		zap.String("job", rm.GetJob()),
	// 		zap.String("instance", rm.instance),
	// 		zap.Int8("rabbitmq_running_nodes_changed_1week", nodeChange1Week))
	// 	return RABBITMQ_RUNNING_NODES_CHANGED_1WEEK, false
	// }

	/* 堆积量一周增长率过滤
	 */
	if alertM, ok := rs.WithRabbitMQLagSumIncrease1WeekRuleFilter(lagSumInc1Week)(rm.RuleItf); !ok {
		alertMsgChan <- am.NewRabbitMQAlertMessage(
			rm.GetJob(), rm.instance, RABBITMQ_LAG_SUM_RATE_LIMIT_1WEEK,
			ut.FormatF2S(alertM.(float32)),
			ut.FormatF2S(lagSumInc1Week),
			ut.FormatF(rm.lagSumUsage),
			ut.FormatF(rm.before1DayLagSumUsage),
			ut.FormatF(rm.before1WeekLagSumUsage))

		global.Logger.Info(
			RABBITMQ_LAG_SUM_RATE_LIMIT_1WEEK,
			zap.String("job", rm.GetJob()),
			zap.String("instance", rm.instance),
			zap.String("rabbitmq_lag_sum_increase_usage_1week", ut.FormatF2S(lagSumInc1Week)))
		return RABBITMQ_LAG_SUM_RATE_LIMIT_1WEEK, false
	}

	/* 一天增长率过滤
	 */
	if alertM, ok := rs.WithRabbitMQLagSumIncrease1DayRuleFilter(lagSumInc1Day)(rm.RuleItf); !ok {
		alertMsgChan <- am.NewRabbitMQAlertMessage(
			rm.GetJob(), rm.instance, RABBITMQ_LAG_SUM_RATE_LIMIT_1DAY,
			ut.FormatF2S(alertM.(float32)),
			ut.FormatF2S(lagSumInc1Day),
			ut.FormatF(rm.lagSumUsage),
			ut.FormatF(rm.before1DayLagSumUsage),
			ut.FormatF(rm.before1WeekLagSumUsage))

		global.Logger.Info(
			RABBITMQ_LAG_SUM_RATE_LIMIT_1DAY,
			zap.String("job", rm.GetJob()),
			zap.String("instance", rm.instance),
			zap.String("rabbitmq_lag_sum_increase_usage_1day", ut.FormatF2S(lagSumInc1Day)))
		return RABBITMQ_LAG_SUM_RATE_LIMIT_1DAY, false
	}

	/* 瞬时值过滤
	 */
	if alertM, ok := rs.WithRabbitMQLagSumRuleFilter(rm.lagSumUsage)(rm.RuleItf); !ok {
		alertMsgChan <- am.NewRabbitMQAlertMessage(
			rm.GetJob(), rm.instance, RABBITMQ_LAG_SUM_LIMIT,
			ut.FormatF(alertM.(float32)),
			ut.FormatF(rm.lagSumUsage),
			ut.FormatF(rm.lagSumUsage),
			ut.FormatF(rm.before1DayLagSumUsage),
			ut.FormatF(rm.before1WeekLagSumUsage))

		global.Logger.Info(
			RABBITMQ_LAG_SUM_LIMIT,
			zap.String("job", rm.GetJob()),
			zap.String("instance", rm.instance),
			zap.Float32("rabbitmq_lag_sum_usage", rm.lagSumUsage))
		return RABBITMQ_LAG_SUM_LIMIT, false
	}
	return "", true
}

func (rm *RabbitMQMetrics) Print() string {
	return fmt.Sprintf("## job: %s,instance: %s,runnningNodes: %d,runningNodesBefore1Day: %d,runningNodesBefore1Week: %d,lagSum: %.2f,lagSumBefore1Day: %.2f,lagSumBefore1Week: %.2f", rm.job, rm.instance, rm.runningNodesUsage, rm.before1DayRunningNodesUsage, rm.before1WeekRunningNodesUsage, rm.lagSumUsage, rm.before1DayLagSumUsage, rm.before1WeekLagSumUsage)
}

func WithRabbitMQJob(job string) MetricsOption {
	return func(rm MetricsItf) {
		rm.(*RabbitMQMetrics).job = job
	}
}

func WithRabbitMQRunningNodesUsage(runningNodes int8) MetricsOption {
	return func(rm MetricsItf) {
		rm.(*RabbitMQMetrics).runningNodesUsage = runningNodes
	}
}
func WithBefore1DayRabbitMQRunningNodesUsage(beforeRunningNodes int8) MetricsOption {
	return func(rm MetricsItf) {
		rm.(*RabbitMQMetrics).before1DayRunningNodesUsage = beforeRunningNodes
	}
}
func WithBefore1WeekRabbitMQRunningNodesUsage(beforeRunningNodes int8) MetricsOption {
	return func(rm MetricsItf) {
		rm.(*RabbitMQMetrics).before1WeekRunningNodesUsage = beforeRunningNodes
	}
}
func WithRabbitMQLagSumUsage(lagSumUsage float32) MetricsOption {
	return func(km MetricsItf) {
		km.(*RabbitMQMetrics).lagSumUsage = lagSumUsage
	}
}
func WithBefore1DayRabbitMQLagSumUsage(beforeLagSumUsage float32) MetricsOption {
	return func(km MetricsItf) {
		km.(*RabbitMQMetrics).before1DayLagSumUsage = beforeLagSumUsage
	}
}
func WithBefore1WeekRabbitMQLagSumUsage(beforeLagSumUsage float32) MetricsOption {
	return func(km MetricsItf) {
		km.(*RabbitMQMetrics).before1WeekLagSumUsage = beforeLagSumUsage
	}
}

func NewRabbitMQMetrics(instance string, options ...MetricsOption) *RabbitMQMetrics {
	mi := &BaseMetrics{instance: instance}
	rm := &RabbitMQMetrics{
		BaseMetrics: mi,
	}
	for _, option := range options {
		option(rm)
	}
	return rm
}
