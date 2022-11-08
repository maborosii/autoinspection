package alert

type rabbitMQAlertMessage struct {
	*baseAlertMessage
}

func NewRabbitMQAlertMessage(job, instance, alertMessage, alertMetricsLimit, alertMetricsUsage, now, before1Day, before1Week string) *rabbitMQAlertMessage {
	b := NewBaseAlertMessage(job, instance, alertMessage, alertMetricsLimit, alertMetricsUsage, now, before1Day, before1Week)
	return &rabbitMQAlertMessage{
		baseAlertMessage: b,
	}
}
