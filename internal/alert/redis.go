package alert

import (
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
)

type RedisAlertMessage struct {
	job, instance, alertMessage          string
	alertMetricsLimit, alertMetricsUsage float32
}

func NewRedisAlertMessage(job, instance, alertMessage string, alertMetricsLimit, alertMetricsUsage float32) *RedisAlertMessage {
	return &RedisAlertMessage{
		job:               job,
		instance:          instance,
		alertMessage:      alertMessage,
		alertMetricsLimit: alertMetricsLimit,
		alertMetricsUsage: alertMetricsUsage,
	}
}

func (n *RedisAlertMessage) PrintAlert() string {
	return fmt.Sprintf("Redis 指标异常 >>> job: %s, instance: %s,  告警信息:%s, 当前值:%.2f, 预警值：%.2f\n", n.job, n.instance, n.alertMessage, n.alertMetricsUsage, n.alertMetricsLimit)
}
func (n *RedisAlertMessage) PrintAlertFormatTable() table.Row {
	return table.Row{"Redis 指标异常", n.job, n.instance, "", n.alertMessage, fmt.Sprintf("%.2f", n.alertMetricsUsage), fmt.Sprintf("%.2f", n.alertMetricsLimit)}
}
