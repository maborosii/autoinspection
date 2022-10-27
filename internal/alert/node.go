package alert

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
)

type NodeAlertMessage struct {
	job, instance, nodeName, alertMessage string
	alertMetricsLimit, alertMetricsUsage  float32
}

func NewNodeAlertMessage(job, instance, nodeName, alertMessage string, alertMetricsLimit, alertMetricsUsage float32) *NodeAlertMessage {
	return &NodeAlertMessage{
		job:               job,
		instance:          instance,
		nodeName:          nodeName,
		alertMessage:      alertMessage,
		alertMetricsLimit: alertMetricsLimit,
		alertMetricsUsage: alertMetricsUsage,
	}
}

func (n *NodeAlertMessage) PrintAlert() string {
	return fmt.Sprintf("主机指标异常 >>> job: %s, instance: %s, 主机名:%s, 告警信息:%s, 当前值:%.2f, 预警值：%.2f\n", n.job, n.instance, n.nodeName, n.alertMessage, n.alertMetricsUsage, n.alertMetricsLimit)
}
func (n *NodeAlertMessage) PrintAlertFormatTable() table.Row {
	return table.Row{"主机指标异常", n.job, n.instance, n.nodeName, n.alertMessage, fmt.Sprintf("%.2f", n.alertMetricsUsage), fmt.Sprintf("%.2f", n.alertMetricsLimit)}
}
