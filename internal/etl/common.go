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

func (b *BaseMetrics) GetJob() string {
	return "basic metrics job"
}
func (b *BaseMetrics) GetInstance() string {
	return "basic metrics instance"
}

type MetricsOption func(MetricsItf)

type MetricsMap map[string]MetricsItf

func (m MetricsMap) CreateOrModify(key string, mopts MetricsOption) {

}
