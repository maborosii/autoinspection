package metrics

import (
	"node_metrics_go/global"
	"sync"

	am "node_metrics_go/internal/alert"
	ph "node_metrics_go/internal/pusher"
	"node_metrics_go/internal/pusher/mail"
	rs "node_metrics_go/internal/rules"

	"go.uber.org/zap"
)

var wgForStopChan sync.WaitGroup

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
	Filter(chan<- am.AlertInfo) (string, bool)
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
func (m MetricsMap) MapToJobAndNodeName(instanceToJob, instanceToNodeName map[string]string) {
	for k, v := range m {
		if _, ok := instanceToJob[k]; !ok {
			global.Logger.Warn("this instance not found in job mapping, ", zap.String("key", k))
		}
		switch v.(type) {
		case *NodeMetrics:
			if _, ok := instanceToNodeName[k]; !ok {
				global.Logger.Warn("this instance not found in nodeName mapping, ", zap.String("key", k))
			}
			WithNodeJob(instanceToJob[k])(v)
			WithNodeName(instanceToNodeName[k])(v)
			global.Logger.Debug("[nodeMetrics] mapping instance to nodeName and job mapping, ", zap.String("key", k), zap.String("job", v.GetJob()))

		case *RedisMetrics:
			WithRedisJob(instanceToJob[k])(v)
			global.Logger.Debug("[redisMetrics] mapping instance and job mapping, ", zap.String("key", k), zap.String("job", v.GetJob()))

		default:
			global.Logger.Warn("unknown type for MetricsItf")
		}
	}
}

// 映射告警规则
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
	var alertMessageChan = make(chan am.AlertInfo, 10)

	for _, v := range m {
		// 并发处理规则匹配
		wgForStopChan.Add(1)
		go func(v MetricsItf) {
			defer wgForStopChan.Done()
			if str, ok := v.Filter(alertMessageChan); !ok {
				global.Logger.Debug(str, zap.String("metrics", v.Print()))
			}
		}(v)
	}
	// 用于关闭通道
	go func() {
		wgForStopChan.Wait()
		close(alertMessageChan)
	}()

	// mailMessage := MergeAlertInfo(alertMessageChan)
	// mm := mail.NewMailMessage(mailMessage)

	tableRows := am.MergeAlertInfoFormatTable(alertMessageChan)
	mailMessage := am.RenderTable(tableRows)
	mm := mail.NewMailMessage(mailMessage)
	if mailMessage != "" {
		ph.PusherList.Exec(mm)
	}
}
