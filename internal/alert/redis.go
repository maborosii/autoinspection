package alert

type redisAlertMessage struct {
	// job, instance, alertMessage, alertMetricsLimit, alertMetricsUsage string
	*baseAlertMessage
}

func NewRedisAlertMessage(job, instance, alertMessage, alertMetricsLimit, alertMetricsUsage string) *redisAlertMessage {
	// return &redisAlertMessage{
	// 	job:               job,
	// 	instance:          instance,
	// 	alertMessage:      alertMessage,
	// 	alertMetricsLimit: alertMetricsLimit,
	// 	alertMetricsUsage: alertMetricsUsage,
	// }
	b := NewBaseAlertMessage(job, instance, alertMessage, alertMetricsLimit, alertMetricsUsage)
	return &redisAlertMessage{
		baseAlertMessage: b,
	}
}

// func (n *redisAlertMessage) PrintAlert() string {
// 	return fmt.Sprintf("Redis 指标异常 >>> job: %s, instance: %s,  告警信息:%s, 当前值:%s, 预警值：%s\n", n.job, n.instance, n.alertMessage, n.alertMetricsUsage, n.alertMetricsLimit)
// }
// func (n *redisAlertMessage) PrintAlertFormatTable() table.Row {
// 	return table.Row{"Redis 指标异常", n.job, n.instance, "", n.alertMessage, n.alertMetricsUsage, n.alertMetricsLimit}
// }
