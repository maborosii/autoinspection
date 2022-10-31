package alert

type kafkaAlertMessage struct {
	// job, instance, alertMessage, alertMetricsLimit, alertMetricsUsage string
	*baseAlertMessage
}

func NewKafkaAlertMessage(job, instance, alertMessage, alertMetricsLimit, alertMetricsUsage string) *kafkaAlertMessage {
	// return &kafkaAlertMessage{
	// 	job:               job,
	// 	instance:          instance,
	// 	alertMessage:      alertMessage,
	// 	alertMetricsLimit: alertMetricsLimit,
	// 	alertMetricsUsage: alertMetricsUsage,
	// }
	b := NewBaseAlertMessage(job, instance, alertMessage, alertMetricsLimit, alertMetricsUsage)
	return &kafkaAlertMessage{
		baseAlertMessage: b,
	}
}

// func (n *kafkaAlertMessage) PrintAlert() string {
// 	return fmt.Sprintf("Kafka 指标异常 >>> job: %s, instance: %s,  告警信息:%s, 当前值:%s, 预警值：%.s\n", n.job, n.instance, n.alertMessage, n.alertMetricsUsage, n.alertMetricsLimit)
// }
// func (n *kafkaAlertMessage) PrintAlertFormatTable() table.Row {
// 	return table.Row{"Kafka 指标异常", n.job, n.instance, "", n.alertMessage, n.alertMetricsUsage, n.alertMetricsLimit}
// }
