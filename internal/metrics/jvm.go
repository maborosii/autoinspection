package metrics

import (
	"fmt"
	"node_metrics_go/global"
	am "node_metrics_go/internal/alert"
	rs "node_metrics_go/internal/rules"
	ut "node_metrics_go/internal/utils"

	"go.uber.org/zap"
)

type JVMMetrics struct {
	*BaseMetrics
	appName             string
	blockedThreadCount  int8
	garbageCollectCount float32
	garbageCollectTime  float32
}

func (jm *JVMMetrics) GetJob() string {
	return jm.job
}
func (jm *JVMMetrics) GetInstance() string {
	return jm.instance
}

// pre asset
// var a MetricsItf
// var _ = a.(*JVMMetrics)
// 适配过滤规则
func (jm *JVMMetrics) AdaptRules(r rs.RuleItf) {
	jm.RuleItf = r
}

// 指标过滤
func (jm *JVMMetrics) Filter(alertMsgChan chan<- am.AlertInfo) (string, bool) {
	logMsg, alertNotFlag := "", true
	// 若该指标项未匹配到规则
	if jm.RuleItf == nil {
		return logMsg, alertNotFlag
	}

	/* 阻塞线程数瞬时值过滤
	 */
	if alertM, ok := rs.WithJVMBlockedThreadCountRuleFilter(jm.blockedThreadCount)(jm.RuleItf); !ok {
		alertMsgChan <- am.NewJVMAlertMessage(
			jm.GetJob(), jm.instance, jm.appName, JVM_BLOCKED_THREAD_COUNT,
			ut.FormatD(alertM.(int8)),
			ut.FormatD(jm.blockedThreadCount),
			ut.FormatD(jm.blockedThreadCount), "", "")

		global.Logger.Info(
			JVM_BLOCKED_THREAD_COUNT,
			zap.String("job", jm.GetJob()),
			zap.String("appName", jm.appName),
			zap.String("instance", jm.instance),
			zap.Int8("jvm_blocked_thread_count", jm.blockedThreadCount))
		// return JVM_BLOCKED_THREAD_COUNT, false
		logMsg = logMsg + ": " + JVM_BLOCKED_THREAD_COUNT + "\n"
		alertNotFlag = false
	}

	/* gc 时间瞬时值过滤
	 */
	if alertM, ok := rs.WithJVMGarbageCollectTimeRuleFilter(jm.garbageCollectTime)(jm.RuleItf); !ok {
		alertMsgChan <- am.NewJVMAlertMessage(
			jm.GetJob(), jm.instance, jm.appName, JVM_GARBAGE_COLLECT_TIME,
			ut.FormatF(alertM.(float32)),
			ut.FormatF(jm.garbageCollectTime),
			ut.FormatF(jm.garbageCollectTime), "", "")

		global.Logger.Info(
			JVM_GARBAGE_COLLECT_TIME,
			zap.String("job", jm.GetJob()),
			zap.String("appName", jm.appName),
			zap.String("instance", jm.instance),
			zap.String("jvm_gc_time", ut.FormatF(jm.garbageCollectTime)))
		// return JVM_GARBAGE_COLLECT_TIME, false
		logMsg = logMsg + ": " + JVM_GARBAGE_COLLECT_TIME + "\n"
		alertNotFlag = false
	}
	/* gc 次数瞬时值过滤
	 */
	if alertM, ok := rs.WithJVMGarbageCollectCountRuleFilter(jm.garbageCollectCount)(jm.RuleItf); !ok {
		alertMsgChan <- am.NewJVMAlertMessage(
			jm.GetJob(), jm.instance, jm.appName, JVM_GARBAGE_COLLECT_COUNT,
			ut.FormatF(alertM.(float32)),
			ut.FormatF(jm.garbageCollectCount),
			ut.FormatF(jm.garbageCollectCount), "", "")

		global.Logger.Info(
			JVM_GARBAGE_COLLECT_COUNT,
			zap.String("job", jm.GetJob()),
			zap.String("appName", jm.appName),
			zap.String("instance", jm.instance),
			zap.String("jvm_gc_count", ut.FormatF(jm.garbageCollectCount)))
		// return JVM_GARBAGE_COLLECT_COUNT, false
		logMsg = logMsg + ": " + JVM_GARBAGE_COLLECT_COUNT + "\n"
		alertNotFlag = false
	}
	return logMsg, alertNotFlag
}

func (jm *JVMMetrics) Print() string {
	return fmt.Sprintf("## job: %s,appName: %s,instance: %s,blockedThreadCount: %d,gcTime: %f, gcCount: %f", jm.job, jm.appName, jm.instance, jm.blockedThreadCount, jm.garbageCollectTime, jm.garbageCollectCount)
}
func WithJVMJob(job string) MetricsOption {
	return func(jm MetricsItf) {
		jm.(*JVMMetrics).job = job
	}
}

func WithJVMAppName(appName string) MetricsOption {
	return func(jm MetricsItf) {
		jm.(*JVMMetrics).appName = appName
	}
}

func WithJVMBlockedThreadCount(blockedThreadCount int8) MetricsOption {
	return func(jm MetricsItf) {
		jm.(*JVMMetrics).blockedThreadCount = blockedThreadCount
	}
}
func WithJVMGarbageCollectTime(gcTime float32) MetricsOption {
	return func(jm MetricsItf) {
		jm.(*JVMMetrics).garbageCollectTime = gcTime
	}
}
func WithJVMGarbageCollectCount(gcCount float32) MetricsOption {
	return func(jm MetricsItf) {
		jm.(*JVMMetrics).garbageCollectCount = gcCount
	}
}
func NewJVMMetrics(instance string, options ...MetricsOption) *JVMMetrics {
	mi := &BaseMetrics{instance: instance}
	jm := &JVMMetrics{
		BaseMetrics: mi,
	}
	for _, option := range options {
		option(jm)
	}
	return jm
}
