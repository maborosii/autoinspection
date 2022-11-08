package alert

type elasticSearchAlertMessage struct {
	*baseAlertMessage
}

func NewElasticSearchAlertMessage(job, instance, alertMessage, alertMetricsLimit, alertMetricsUsage, now, before1Day, before1Week string) *elasticSearchAlertMessage {
	b := NewBaseAlertMessage(job, instance, alertMessage, alertMetricsLimit, alertMetricsUsage, now, before1Day, before1Week)
	return &elasticSearchAlertMessage{
		baseAlertMessage: b,
	}
}
