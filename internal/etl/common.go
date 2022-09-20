package etl

// 公有指标
type BaseMetrics struct {
	Job      string
	Instance string
}

// 指标接口
type MetricsItf interface {
	GetJob() string
	GetInstance() string
}
