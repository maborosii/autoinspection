package alert

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
)

type jvmAlertMessage struct {
	*baseAlertMessage
	appName string
}

func NewJVMAlertMessage(job, instance, appName, alertMessage, alertMetricsLimit, alertMetricsUsage, now, before1Day, before1Week string) *jvmAlertMessage {
	b := NewBaseAlertMessage(job, instance, alertMessage, alertMetricsLimit, alertMetricsUsage, now, before1Day, before1Week)
	return &jvmAlertMessage{
		baseAlertMessage: b,
		appName:          appName,
	}
}

func (n *jvmAlertMessage) PrintAlert(mType string) string {
	return fmt.Sprintf("%s 指标异常 >>> job: %s, instance: %s,应用名:%s, 告警信息:%s, 当前值: %s, 预警值： %s, 指标值（当前）: %s, 指标值（一天前）: %s, 指标值（一周前）: %s\n", mType, n.job, n.instance, n.appName, n.alertMessage, n.alertMetricsUsage, n.alertMetricsLimit, n.metricsNow, n.metricsBefore1Day, n.metricsBefore1Week)
}
func (n *jvmAlertMessage) PrintAlertFormatTable(mType string) table.Row {
	return table.Row{fmt.Sprintf("%s 指标异常", mType), n.job, n.instance, n.appName, n.alertMessage, n.alertMetricsUsage, n.alertMetricsLimit, n.metricsNow, n.metricsBefore1Day, n.metricsBefore1Week}
}
