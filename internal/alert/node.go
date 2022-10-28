package alert

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
)

type nodeAlertMessage struct {
	job, instance, nodeName, alertMessage, alertMetricsLimit, alertMetricsUsage string
}

func NewNodeAlertMessage(job, instance, nodeName, alertMessage, alertMetricsLimit, alertMetricsUsage string) *nodeAlertMessage {
	return &nodeAlertMessage{
		job:               job,
		instance:          instance,
		nodeName:          nodeName,
		alertMessage:      alertMessage,
		alertMetricsLimit: alertMetricsLimit,
		alertMetricsUsage: alertMetricsUsage,
	}
}

func (n *nodeAlertMessage) PrintAlert() string {
	return fmt.Sprintf("主机指标异常 >>> job: %s, instance: %s, 主机名:%s, 告警信息:%s, 当前值:%s, 预警值：%s\n", n.job, n.instance, n.nodeName, n.alertMessage, n.alertMetricsUsage, n.alertMetricsLimit)
}
func (n *nodeAlertMessage) PrintAlertFormatTable() table.Row {
	return table.Row{"主机指标异常", n.job, n.instance, n.nodeName, n.alertMessage, n.alertMetricsUsage, n.alertMetricsLimit}
}
