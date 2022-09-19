package etl

type MetricsInfo struct {
	Job      string
	Instance string
}
type MetricsItf interface {
	GetJob() string
	GetInstance() string
}
