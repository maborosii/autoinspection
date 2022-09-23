package etl

import (
	"node_metrics_go/global"
	rs "node_metrics_go/pkg/rules"

	"go.uber.org/zap"
)

// 公有指标
type BaseMetrics struct {
	job      string
	instance string
	rs.RuleItf
}

// 指标接口
type MetricsItf interface {
	GetJob() string
	GetInstance() string
	Print() string
	AdaptRules(rs.RuleItf)
	Filter() (string, bool)
}

func (b *BaseMetrics) GetJob() string {
	return "basic metrics job"
}
func (b *BaseMetrics) GetInstance() string {
	return "basic metrics instance"
}

type MetricsOption func(MetricsItf)

type MetricsMap map[string]MetricsItf

func (m MetricsMap) CreateOrModify(key string, t MetricsItf, opts ...MetricsOption) {
	if _, ok := m[key]; !ok {
		m[key] = t
	}
	for _, opt := range opts {
		opt(m[key])
	}
}
func (m MetricsMap) MapToJobAndNodeName() {
	for k, v := range m {
		if _, ok := instanceToJob[k]; !ok {
			global.Logger.Warn("in instance to job mapping, ", zap.String("key", k))
		}
		if _, ok := instanceToNodeName[k]; !ok {
			global.Logger.Warn("in instance to nodeName mapping, ", zap.String("key", k))
		}
		switch v.(type) {
		case *NodeMetrics:
			WithNodeJob(instanceToJob[k])(v)
			WithNodeName(instanceToNodeName[k])(v)
			global.Logger.Info("mapping instance to nodeName and job mapping, ", zap.String("key", k), zap.String("job", v.GetJob()))
		default:
			global.Logger.Info("unknown type for MetricsItf")
		}
	}
}

func (m MetricsMap) MapToRules() {
	for _, v := range m {
		metricsJob := v.GetJob()
		if _, ok := global.NotifyRules[metricsJob]; !ok {
			global.Logger.Error("can not find matched notify rule", zap.String("job", metricsJob))
			v.AdaptRules(nil)
		}
		v.AdaptRules(global.NotifyRules[metricsJob])
	}
}
func (m MetricsMap) Notify() {

	for _, v := range m {
		if str, ok := v.Filter(); !ok {
			global.Logger.Info(str, zap.String("metrics", v.Print()))
		}
	}
}
