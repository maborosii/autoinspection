package alert

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
)

type nodeAlertMessage struct {
	// job, instance, nodeName, alertMessage, alertMetricsLimit, alertMetricsUsage string
	*baseAlertMessage
	nodeName string
}

func NewNodeAlertMessage(job, instance, nodeName, alertMessage, alertMetricsLimit, alertMetricsUsage, now, before1Day, before1Week string) *nodeAlertMessage {
	// return &nodeAlertMessage{
	// 	job:               job,
	// 	instance:          instance,
	// 	nodeName:          nodeName,
	// 	alertMessage:      alertMessage,
	// 	alertMetricsLimit: alertMetricsLimit,
	// 	alertMetricsUsage: alertMetricsUsage,
	// }
	b := NewBaseAlertMessage(job, instance, alertMessage, alertMetricsLimit, alertMetricsUsage, now, before1Day, before1Week)
	return &nodeAlertMessage{
		baseAlertMessage: b,
		nodeName:         nodeName,
	}
}

func (n *nodeAlertMessage) PrintAlert(mType string) string {
	return fmt.Sprintf("%s 指标异常 >>> job: %s, instance: %s, 主机名:%s, 告警信息:%s, 当前值: %s, 预警值： %s, 指标值（当前）: %s, 指标值（一天前）: %s, 指标值（一周前）: %s\n", mType, n.job, n.instance, n.nodeName, n.alertMessage, n.alertMetricsUsage, n.alertMetricsLimit, n.metricsNow, n.metricsBefore1Day, n.metricsBefore1Week)
}
func (n *nodeAlertMessage) PrintAlertFormatTable(mType string) table.Row {
	return table.Row{fmt.Sprintf("%s 指标异常", mType), n.job, n.instance, n.nodeName, n.alertMessage, n.alertMetricsUsage, n.alertMetricsLimit, n.metricsNow, n.metricsBefore1Day, n.metricsBefore1Week}
}
