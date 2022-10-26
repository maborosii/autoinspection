package metrics

import (
	"fmt"
	"node_metrics_go/global"
	"sync"

	ph "node_metrics_go/internal/pusher"
	"node_metrics_go/internal/pusher/mail"
	rs "node_metrics_go/internal/rules"

	"github.com/jedib0t/go-pretty/v6/table"

	"go.uber.org/zap"
)

var wgForStopChan sync.WaitGroup

// 公有指标
type BaseMetrics struct {
	job      string
	instance string
	rs.RuleItf
}
type AlertInfo interface {
	PrintAlert() string
	PrintAlertFormatTable() table.Row
}

// 指标接口
type MetricsItf interface {
	GetJob() string
	GetInstance() string
	Print() string
	AdaptRules(rs.RuleItf)
	Filter(chan<- AlertInfo) (string, bool)
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

// 合并告警信息
// 文本信息
func mergeAlertInfo(a <-chan AlertInfo) string {
	var mm string
	alertInfoByKind := make(map[string]string, 5)
	for v := range a {
		switch v.(type) {
		case *NodeOutputMessage:
			alertInfoByKind["node"] += v.PrintAlert()
		case *RedisOutputMessage:
			alertInfoByKind["redis"] += v.PrintAlert()
		default:
			global.Logger.Warn("alert info not found suitable type", zap.String("info", v.PrintAlert()))
		}
	}
	for _, infos := range alertInfoByKind {
		mm += infos
	}
	return mm
}

// 表格信息
func mergeAlertInfoFormatTable(a <-chan AlertInfo) []table.Row {
	var mm []table.Row
	alertInfoByKind := make(map[string][]table.Row, 5)
	for v := range a {
		switch v.(type) {
		case *NodeOutputMessage:
			alertInfoByKind["node"] = append(alertInfoByKind["node"], v.PrintAlertFormatTable())
		case *RedisOutputMessage:
			alertInfoByKind["redis"] = append(alertInfoByKind["redis"], v.PrintAlertFormatTable())
		default:
			global.Logger.Warn("alert info not found suitable type", zap.String("info", v.PrintAlert()))
		}
	}
	for _, infos := range alertInfoByKind {
		mm = append(mm, infos...)
	}
	return mm
}

// render table to html
func renderTable(rows []table.Row) string {
	t := table.NewWriter()
	t.AppendHeader(tableHeader)
	t.AppendRows(rows)
	t.Style().HTML = table.HTMLOptions{
		CSSClass:    "",
		EmptyColumn: "&nbsp;",
		EscapeText:  true,
		Newline:     "<br/>",
	}

	prefixMailHtml := fmt.Sprintf("<style>\n%s\n</style>\n", styleCss)
	htmlContext := prefixMailHtml + t.RenderHTML()
	return htmlContext
}

func (m MetricsMap) Notify() {
	var alertMessageChan = make(chan AlertInfo, 10)

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

	tableRows := mergeAlertInfoFormatTable(alertMessageChan)
	mailMessage := renderTable(tableRows)
	mm := mail.NewMailMessage(mailMessage)
	if mailMessage != "" {
		ph.PusherList.Exec(mm)
	}
}
